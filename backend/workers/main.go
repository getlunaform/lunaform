package workers

import (
	"github.com/gammazero/workerpool"
	"github.com/getlunaform/lunaform/models"
	"github.com/getlunaform/lunaform/backend/database"
	"github.com/go-openapi/swag"
	"github.com/getlunaform/go-terraform"
	"fmt"
	"os"
)

const (
	TF_ACTION_PLAN_TYPE = "plan"
	TF_ACTION_INIT_TYPE = "init"
)

type TfAction interface {
	Type() *string
}

type TfActionPlan struct {
	Module     *models.ResourceTfModule
	Stack      *models.ResourceTfStack
	Deployment *models.ResourceTfDeployment
	DoInit     bool
}

type TfActionInit struct {
	Module     *models.ResourceTfModule
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

func (p *TfAgentPool) DoInit(a *TfActionInit) {
	p.pool.Submit(
		a.BuildJob(p.scratchFolder))
}

func (p *TfAgentPool) DoPlan(a *TfActionPlan) {
	p.pool.Submit(
		a.BuildJob(p.scratchFolder))
}

type gitProgressBuffer struct {
	logs *goterraform.OutputLog
}

func (gpb gitProgressBuffer) Write(b []byte) (delta int, err error) {
	l := len(b)
	fmt.Print(
		gpb.logs.StdoutWithTags(
			string(b),
			[]string{"git"},
		).String(),
	)
	return l, nil

}
