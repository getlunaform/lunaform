package workers

import (
	"github.com/gammazero/workerpool"
	"github.com/getlunaform/lunaform/backend/database"
	"os"
)

type TfAgentPool struct {
	maxWorkers    int
	db            database.Database
	pool          *workerpool.WorkerPool
	scratchFolder string
	debug         bool
}

func NewAgentPool(maxWorkers int) *TfAgentPool {
	tempdir := os.TempDir()
	return &TfAgentPool{
		maxWorkers:    maxWorkers,
		scratchFolder: tempdir,
	}
}

func (p *TfAgentPool) SetDebug() {
	p.debug = true
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
	a.BuildJob(p.scratchFolder)
}

func (p *TfAgentPool) DoPlan(a *TfActionPlan) {
	a.BuildJob(p.scratchFolder)
}
