package workers

import (
	"path/filepath"
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
				Module: a.Module,
				Stack:  a.Stack,
			}
			init.BuildJob(scratchFolder)()
		}

		stackDirectory := "stack-" + a.Stack.ID
		fullStackDirectory := filepath.Join(scratchFolder, stackDirectory)
		planFileName := "deployment-" + a.Deployment.ID + ".plan"

		params := &goterraform.TerraformPlanParams{
			Out:   swag.String(planFileName),
			Input: swag.Bool(false),
		}

		action := goterraform.NewTerraformClient().
			WithWorkingDirectory(fullStackDirectory).
			Plan(params).
			Initialise()

		logs := goterraform.NewOutputLogs()
		if err := action.InitLogger(logs); err != nil {
			fmt.Printf("An error occured initialising task logger: " + err.Error())
			return
		}

		if err := os.Mkdir(fullStackDirectory, 0700); err != nil {
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
