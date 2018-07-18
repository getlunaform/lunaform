package workers

import (
	"github.com/gammazero/workerpool"
	"github.com/drewsonne/lunaform/backend/database"
	"github.com/drewsonne/lunaform/server/models"
	"fmt"
)

func NewAgentPool(maxWorkers int) *TfAgentPool {
	return &TfAgentPool{
		maxWorkers: maxWorkers,
	}
}

type TfAgentPool struct {
	maxWorkers int

	DB database.Database

	pool *workerpool.WorkerPool
}

func (p *TfAgentPool) Start() *TfAgentPool {
	p.pool = workerpool.New(p.maxWorkers)
	return p
}

func (p *TfAgentPool) WithDB(db database.Database) *TfAgentPool {
	p.DB = db
	return p
}

func (p *TfAgentPool) DoPlan(s *models.ResourceTfStack) {
	p.pool.Submit(func() {
		action, output := NewTerraformClient().Plan(s)
		action.Run()
		fmt.Print(action)
		fmt.Print(output)
	})
}
