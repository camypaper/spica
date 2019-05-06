package core

import (
	"sync"

	"github.com/Camypaper/libra"
)

/*
OutputReporter shows progress of job.
*/
type OutputReporter struct {
	name string
	dic  sync.Map
	All  bool
}

/*
ReportStart initialize progress bar.
*/
func (reporter *OutputReporter) ReportStart(job libra.Job) {
	reporter.name = job.Name()
}

/*
Report updates progress.
*/
func (reporter *OutputReporter) Report(task libra.Task, status libra.Status) {
	if reporter.All || status.Code == libra.OK {
		reporter.dic.LoadOrStore(task.Name(), status)
	}
}

/*
ReportEnd shows result.
*/
func (reporter *OutputReporter) ReportEnd() {

}

/*
Get returns result of
*/
func (reporter *OutputReporter) Get() map[string]libra.Status {
	ret := map[string]libra.Status{}
	reporter.dic.Range(func(key, value interface{}) bool {
		ret[key.(string)] = value.(libra.Status)
		return true
	})
	return ret
}
