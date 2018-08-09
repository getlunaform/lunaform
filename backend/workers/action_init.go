package workers

import (
	"fmt"
	"github.com/getlunaform/go-terraform"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
	"os"
)

type TfActionInit struct {
	Stack      *models.ResourceTfStack
	Deployment *models.ResourceTfDeployment
	logs       *goterraform.OutputLog
}

func newTfActionInit() *TfActionInit {
	return &TfActionInit{
		logs: goterraform.NewOutputLogs(),
	}
}

func (tai *TfActionInit) WithStack(stack *models.ResourceTfStack) *TfActionInit {
	tai.Stack = stack
	return tai
}

func (tai *TfActionInit) WithDeployment(deployment *models.ResourceTfDeployment) *TfActionInit {
	tai.Deployment = deployment
	return tai
}

func (tai *TfActionInit) WithLogs(logs *goterraform.OutputLog) *TfActionInit {
	tai.logs = logs
	return tai
}

func (tap *TfActionInit) Type() TFActionType {
	return TfActionInitType
}

func (tai *TfActionInit) BuildJob(scratchFolder string) (err error) {

	if tai.logs == nil {
		tai.logs = goterraform.NewOutputLogs()
	}

	module := tai.Stack.Embedded.Module

	bs := NewBuildSpace(scratchFolder).
		WithModule(module).
		WithStack(tai.Stack).
		WithDeployment(tai.Deployment)

	params := &goterraform.TerraformInitParams{
		FromModule: *module.Source,
		Input:      swag.Bool(false),
		Upgrade:    swag.Bool(true),
	}

	action := goterraform.NewTerraformClient().
		WithWorkingDirectory(bs.MustDeploymentDirectory(true)).
		Init(params).
		Initialise()

	logs := goterraform.NewOutputLogs()
	if err = action.InitLogger(logs); err != nil {
		fmt.Printf("An error occured initialising task logger: " + err.Error())
		return
	}

	if err = os.MkdirAll(bs.MustDeploymentDirectory(true), 0700); err != nil {
		logs.Error(err)
		return
	}

	if err = action.Cmd.Start(); err != nil {
		logs.Error(err)
		return
	}

	if err = action.Cmd.Wait(); err != nil {
		logs.Error(err)
		return
	}
	return
}
