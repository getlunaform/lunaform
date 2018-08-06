package workers

import (
	"github.com/getlunaform/go-terraform"
	"github.com/go-openapi/swag"
	"fmt"
	"os"
	"golang.org/x/sys/unix"
)

func (a *TfActionPlan) BuildJob(scratchFolder string) {

	logs := goterraform.NewOutputLogs()

	if a.DoInit {
		init := &TfActionInit{
			Module:     a.Module,
			Stack:      a.Stack,
			Deployment: a.Deployment,
		}
		init.BuildJob(scratchFolder)
	}

	bs := NewBuildSpace(scratchFolder).
		WithModule(a.Module).
		WithStack(a.Stack).
		WithDeployment(a.Deployment)

	varFilePath, err := bs.VarFilePath(true)
	if err != nil {
		fmt.Print(logs.Error(err))
		return
	}

	vars := newVariableFile(varFilePath)

	vars.Parse(a.Stack.Variables)
	if err := vars.WriteToFile(); err != nil {
		fmt.Print(logs.Error(err))
		return
	}

	providerFilePath, err := bs.ProviderFilePath(true)
	if err != nil {
		fmt.Print(logs.Error(err))
		return
	}

	providers := newProviderFile(providerFilePath)
	for _, conf := range a.Stack.Embedded.ProviderConfigurations {
		providers.Providers[*conf.Embedded.Provider.Name] = conf.Configuration
	}

	if err := providers.WriteToFile(); err != nil {
		fmt.Print(logs.Error(err))
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

	if err := action.InitLogger(logs); err != nil {
		fmt.Printf(
			"An error occured initialising task logger: %s", err.Error())
		return
	}

	if err := os.MkdirAll(bs.MustDeploymentDirectory(true), 0700); err != nil {
		if pathErr, isPathErr := err.(*os.PathError); isPathErr {
			if pathErr.Err != unix.EEXIST {
				fmt.Print(logs.Error(err))
				return
			}
		}
	}

	if err := action.Cmd.Start(); err != nil {
		fmt.Print(logs.Error(err))
		return
	}

	if err := action.Cmd.Wait(); err != nil {
		fmt.Print(logs.Error(err))
		return
	}

}
