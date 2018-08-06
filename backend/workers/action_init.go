package workers

import (
	"github.com/getlunaform/go-terraform"
	"fmt"
	"os"
	"github.com/go-openapi/swag"
)

func (a *TfActionInit) BuildJob(scratchFolder string) {

	bs := NewBuildSpace(scratchFolder).
		WithModule(a.Module).
		WithStack(a.Stack).
		WithDeployment(a.Deployment)

	params := &goterraform.TerraformInitParams{
		FromModule: *a.Module.Source,
		Input:      swag.Bool(false),
		Upgrade:    swag.Bool(true),
	}

	action := goterraform.NewTerraformClient().
		WithWorkingDirectory(bs.MustDeploymentDirectory(true)).
		Init(params).
		Initialise()

	logs := goterraform.NewOutputLogs()
	if err := action.InitLogger(logs); err != nil {
		fmt.Printf("An error occured initialising task logger: " + err.Error())
		return
	}

	if err := os.MkdirAll(bs.MustDeploymentDirectory(true), 0700); err != nil {
		logs.Error(err)
		return
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
