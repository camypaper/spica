package judge

import (
	"github.com/camypaper/libra"
	"github.com/camypaper/spica/core"
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
