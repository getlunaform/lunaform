package workers

import (
	"github.com/gammazero/workerpool"
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/go-openapi/swag"
	"github.com/getlunaform/go-terraform"
	"fmt"
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
	maxWorkers int
	db         database.Database
	pool       *workerpool.WorkerPool
}

func NewAgentPool(maxWorkers int) *TfAgentPool {
	return &TfAgentPool{
		maxWorkers: maxWorkers,
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

		logs := goterraform.NewOutputLogs()

		action := goterraform.NewTerraformClient().Plan().Init()
		if err := action.InitLogger(logs); err != nil {

			fmt.Printf("An error occured initialising task logger: " + err.Error())
			return
		}

		if err := action.Cmd.Start(); err != nil {
			logs.Stderr(err.Error())
			return
		}

		if err := action.Cmd.Wait(); err != nil {
			logs.Stderr(err.Error())
			return
		}

	})

}
