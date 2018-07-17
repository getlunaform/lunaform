package workers

import (
	"github.com/vardius/worker-pool"
)

type TfAgentPool struct {
	QueueLength int
	MaxWorkers  int
}

func (p *TfAgentPool) Start() {

}
