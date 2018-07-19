package workers

import (
	"fmt"
	"github.com/gammazero/workerpool"
	models "github.com/getlunaform/lunaform-models-go"
	"github.com/getlunaform/lunaform/server/backend/database"
	"github.com/go-openapi/swag"
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
		action, output := NewTerraformClient().Plan(a)
		action.Run()
		fmt.Print(action)
		fmt.Print(output)
	})
}
