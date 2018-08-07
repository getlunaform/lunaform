package workers

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/go-terraform"
)

type TfActionDeploy struct {
	Stack      *models.ResourceTfStack
	Deployment *models.ResourceTfDeployment
	logs       *goterraform.OutputLog
	DoInit     bool
}

func (tap *TfActionDeploy) Type() TFActionType {
	return TfActionDeployType
}

func (tap *TfActionDeploy) WithStack(stack *models.ResourceTfStack) *TfActionDeploy {
	tap.Stack = stack
	return tap
}

func (tap *TfActionDeploy) WithDeployment(deployment *models.ResourceTfDeployment) *TfActionDeploy {
	tap.Deployment = deployment
	return tap
}

func (tap *TfActionDeploy) WithLogs(logs *goterraform.OutputLog) *TfActionDeploy {
	tap.logs = logs
	return tap
}

func (tad *TfActionDeploy) BuildJob(scratchFolder string) (err error) {

	if tad.logs == nil {
		tad.logs = goterraform.NewOutputLogs()
	}

	if tad.DoInit {
		action := newTfActionInit().
			WithStack(tad.Stack).
			WithDeployment(tad.Deployment).
			WithLogs(tad.logs)
		if err = action.BuildJob(scratchFolder); err != nil {
			tad.logs.Error(err)
			return
		}
	}
	return
}
