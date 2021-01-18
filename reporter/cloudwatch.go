package reporter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
	"github.com/sirupsen/logrus"
)

type CloudwatchReporter struct {
	client cloudwatchiface.CloudWatchAPI
	logger *logrus.Entry
}

func (c *CloudwatchReporter) Report(r *Report) error {
	// More than 150 TPS and we're borked. TODO: Refactor.
	input := &cloudwatch.PutMetricDataInput{
		MetricData: []*cloudwatch.MetricDatum{
			{
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("Status"),
						Value: aws.String(string(r.Status)),
					},
				},
				MetricName: aws.String("TransactionStarted"),
				Timestamp:  &r.StartTime,
				Unit:       aws.String("Count"),
				Value:      aws.Float64(1),
			},
			{
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("Status"),
						Value: aws.String(string(r.Status)),
					},
				},
				MetricName: aws.String("TransactionEnded"),
				Timestamp:  &r.EndTime,
				Unit:       aws.String("Count"),
				Value:      aws.Float64(1),
			},
			{
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("Status"),
						Value: aws.String(string(r.Status)),
					},
				},
				MetricName: aws.String("Duration"),
				Timestamp:  &r.EndTime,
				Unit:       aws.String("Milliseconds"),
				Value:      aws.Float64(float64(r.EndTime.Sub(r.StartTime).Milliseconds())),
			},
		},
	}
	c.logger.Debug("Sending Report to Cloudwatch")
	_, err := c.client.PutMetricData(input)
	if err != nil {
		c.logger.
			WithError(err).
			Error("Error putting metric data")
	}
	return err
}
