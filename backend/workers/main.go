package workers

import (
	"github.com/gammazero/workerpool"
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/go-openapi/swag"
	"github.com/getlunaform/go-terraform"
	"fmt"
	"os"
	"path/filepath"
)

const (
	TF_ACTION_PLAN_TYPE = "plan"
)

type TfAction interface {
	Type() *string
}

type TfActionPlan struct {
	Stack      *models.ResourceTfStack
	Deployment *models.ResourceTfDeployment
}

func (tap *TfActionPlan) Type() *string {
	return swag.String(TF_ACTION_PLAN_TYPE)
}

type TfAgentPool struct {
	maxWorkers    int
	db            database.Database
	pool          *workerpool.WorkerPool
	scratchFolder string
}

func NewAgentPool(maxWorkers int) *TfAgentPool {
	tempdir := os.TempDir()
	return &TfAgentPool{
		maxWorkers:    maxWorkers,
		scratchFolder: tempdir,
	}
}

func (p *TfAgentPool) Shutdown() bool {
	return true
}

func (p *TfAgentPool) Start() *TfAgentPool {
	p.pool = workerpool.New(p.maxWorkers)
	return p
}

func (p *TfAgentPool) WithDB(db database.Database) *TfAgentPool {
	p.db = db
	return p
}

func (p *TfAgentPool) DoPlan(a *TfActionPlan) {
	p.pool.Submit(func() {

		workingDir := filepath.Join(p.scratchFolder, "stack-"+a.Stack.ID)

		params := &goterraform.TerraformPlanParams{
			Out: swag.String(filepath.Join(workingDir, "deployment-"+a.Deployment.ID+".plan")),
		}

		action := goterraform.NewTerraformClient().
			WithWorkingDirectory(workingDir).
			Plan(params).
			Init()

		logs := goterraform.NewOutputLogs()
		if err := action.InitLogger(logs); err != nil {
			fmt.Printf("An error occured initialising task logger: " + err.Error())
			return
		}

		if err := os.Mkdir(workingDir, os.ModeDir); err != nil {
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

	})

}
