package workers

import (
	"path/filepath"
	"github.com/getlunaform/go-terraform"
	"fmt"
	"os"
)

func (a *TfActionInit) BuildJob(scratchFolder string) func() {
	return func() {
		stackDirectory := "stack-" + a.Stack.ID
		fullStackDirectory := filepath.Join(scratchFolder, stackDirectory)

		params := goterraform.NewTerraformInitParams()
		params.FromModule = *a.Module.Source

		action := goterraform.NewTerraformClient().
			WithWorkingDirectory(fullStackDirectory).
			Init(params).
			Initialise()

		logs := goterraform.NewOutputLogs()
		if err := action.InitLogger(logs); err != nil {
			fmt.Printf("An error occured initialising task logger: " + err.Error())
			return
		}

		if err := os.Mkdir(fullStackDirectory, 0700); err != nil {
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

}
