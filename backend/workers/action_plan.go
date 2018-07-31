package workers

import (
	"github.com/getlunaform/go-terraform"
	"github.com/go-openapi/swag"
	"fmt"
	"os"
	"golang.org/x/sys/unix"
)

func (a *TfActionPlan) BuildJob(scratchFolder string) func() {
	return func() {

		if a.DoInit {
			init := &TfActionInit{
				Module:     a.Module,
				Stack:      a.Stack,
				Deployment: a.Deployment,
			}
			init.BuildJob(scratchFolder)()
		}

		bs := NewBuildSpace(scratchFolder).
			WithModule(a.Module).
			WithStack(a.Stack).
			WithDeployment(a.Deployment)

		params := &goterraform.TerraformPlanParams{
			Out:   swag.String(bs.MustPlanPath(false)),
			Input: swag.Bool(false),
			Var:   &a.Stack.Variables,
		}

		action := goterraform.NewTerraformClient().
			WithWorkingDirectory(bs.MustStackDir(true)).
			Plan(params).
			Initialise()

		logs := goterraform.NewOutputLogs()
		if err := action.InitLogger(logs); err != nil {
			fmt.Printf("An error occured initialising task logger: " + err.Error())
			return
		}

		if err := os.MkdirAll(bs.MustDeploymentDirectory(true), 0700); err != nil {
			if pathErr, isPathErr := err.(*os.PathError); isPathErr {
				if pathErr.Err != unix.EEXIST {
					logs.Error(err)
					return
				}
			}
		}

		if err := action.Cmd.Start(); err != nil {
			logs.Error(err)
			return
		}

		if err := action.Cmd.Wait(); err != nil {
			logs.Error(err)
			return
		}

	}
}
