package judge

import (
	"github.com/Camypaper/libra"
	"github.com/Camypaper/spica/core"
)

/*
SpicaWorker is worker
*/
type SpicaWorker struct {
	libra.Worker
	Context libra.WorkerContext
	Config  core.Config
	Problem core.Problem
}
