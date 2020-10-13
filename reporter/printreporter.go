package reporter

import "github.com/sirupsen/logrus"

type PrintReporter struct {
	logger logrus.Logger
}

func (c *PrintReporter) Report(r *Report) error {
	c.logger.
		WithField("startTime", r.startTime).
		WithField("endTime", r.endTime).
		WithField("status", r.status).
		WithField("message", r.message).
		Info()
	return nil
}
