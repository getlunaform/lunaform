package workers

import (
	"fmt"
	"github.com/getlunaform/go-terraform"
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/swag"
	"golang.org/x/sys/unix"
	"os"
)

type TfActionPlan struct {
	Module     *models.ResourceTfModule
	Stack      *models.ResourceTfStack
	Deployment *models.ResourceTfDeployment
	log        *goterraform.OutputLog
	bs         *BuildSpace
	DoInit     bool
}

func (tap *TfActionPlan) Type() TFActionType {
	return TfActionPlanType
}

func NewTfActionPlan(doInit bool) *TfActionPlan {
	return &TfActionPlan{
		DoInit: doInit,
		log:    goterraform.NewOutputLogs(),
	}
}

func (tap *TfActionPlan) WithStack(stack *models.ResourceTfStack) *TfActionPlan {
	tap.Stack = stack
	return tap
}

func (tap *TfActionPlan) WithDeployment(deployment *models.ResourceTfDeployment) *TfActionPlan {
	tap.Deployment = deployment
	return tap
}

func (tap *TfActionPlan) BuildJob(scratchFolder string) (err error) {

	if tap.log == nil {
		tap.log = goterraform.NewOutputLogs()
	}

	if tap.DoInit {
		action := newTfActionInit().
			WithStack(tap.Stack).
			WithDeployment(tap.Deployment).
			WithLogs(tap.log)
		if err = action.BuildJob(scratchFolder); err != nil {
			return
		}
	}

	bs := NewBuildSpace(scratchFolder).
		WithStack(tap.Stack).
		WithDeployment(tap.Deployment).
		WithModule(tap.Module)

	varFilePath, err := bs.VarFilePath(true)
	if err != nil {
		fmt.Print(tap.log.Error(err))
		return
	}

	vars := newVariableFile(varFilePath)

	vars.Parse(tap.Stack.Variables)
	if err = vars.WriteToFile(); err != nil {
		fmt.Print(tap.log.Error(err))
		return
	}

	providerFilePath, err := bs.ProviderFilePath(true)
	if err != nil {
		fmt.Print(tap.log.Error(err))
		return
	}

	providers := newProviderFile(providerFilePath)
	for _, conf := range tap.Stack.Embedded.ProviderConfigurations {
		providers.Providers[*conf.Embedded.Provider.Name] = conf.Configuration
	}

	if err = providers.WriteToFile(); err != nil {
		fmt.Print(tap.log.Error(err))
		return
	}

	tf := goterraform.NewTerraformClient().WithWorkingDirectory(
		bs.MustDeploymentDirectory(true),
	)

	action := tf.Plan(&goterraform.TerraformPlanParams{
		Out:     swag.String(bs.MustPlanPath(false)),
		Input:   swag.Bool(false),
		VarFile: swag.StringSlice([]string{varFilePath}),
	}).Initialise()

	if err = action.InitLogger(tap.log); err != nil {
		fmt.Printf(
			"An error occured initialising task logger: %s", err.Error())
		return
	}

	if err = os.MkdirAll(bs.MustDeploymentDirectory(true), 0700); err != nil {
		if pathErr, isPathErr := err.(*os.PathError); isPathErr {
			if pathErr.Err != unix.EEXIST {
				fmt.Print(tap.log.Error(err))
				return
			}
		}
	}

	if err = action.Cmd.Start(); err != nil {
		fmt.Print(tap.log.Error(err))
		return
	}

	if err = action.Cmd.Wait(); err != nil {
		fmt.Print(tap.log.Error(err))
		return
	}

	return
}
