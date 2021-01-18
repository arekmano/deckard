package reporter

import "github.com/sirupsen/logrus"

type PrintReporter struct {
	Logger *logrus.Logger
}

func (c *PrintReporter) Report(r *Report) error {
	c.Logger.
		WithField("StartTime", r.StartTime).
		WithField("EndTime", r.EndTime).
		WithField("Status", r.Status).
		WithField("Message", r.Message).
		Info()
	return nil
}
