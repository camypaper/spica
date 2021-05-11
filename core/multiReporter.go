package core

import (
	"github.com/camypaper/libra"
)

/*
MultiReporter merges some repoters.
*/
type MultiReporter struct {
	reporters []libra.Reporter
}

/*
NewMultiReporter initialize MultiReporter.
*/
func NewMultiReporter(items ...libra.Reporter) *MultiReporter {
	return &MultiReporter{reporters: items}
}

/*
ReportStart calls ReportStart one by one.
*/
func (reporter *MultiReporter) ReportStart(job libra.Job) {
	for _, v := range reporter.reporters {
		v.ReportStart(job)
	}
}

/*
Report calls Report one by one.
*/
func (reporter *MultiReporter) Report(task libra.Task, status libra.Status) {
	for _, v := range reporter.reporters {
		v.Report(task, status)
	}
}

/*
ReportEnd calls ReportEnd one by one.
*/
func (reporter *MultiReporter) ReportEnd() {
	for _, v := range reporter.reporters {
		v.ReportEnd()
	}
}
