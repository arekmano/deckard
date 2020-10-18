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
						Value: aws.String(string(r.status)),
					},
				},
				MetricName: aws.String("TransactionStarted"),
				Timestamp:  &r.startTime,
				Unit:       aws.String("Count"),
				Value:      aws.Float64(1),
			},
			{
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("Status"),
						Value: aws.String(string(r.status)),
					},
				},
				MetricName: aws.String("TransactionEnded"),
				Timestamp:  &r.endTime,
				Unit:       aws.String("Count"),
				Value:      aws.Float64(1),
			},
			{
				Dimensions: []*cloudwatch.Dimension{
					{
						Name:  aws.String("Status"),
						Value: aws.String(string(r.status)),
					},
				},
				MetricName: aws.String("Duration"),
				Timestamp:  &r.endTime,
				Unit:       aws.String("Milliseconds"),
				Value:      aws.Float64(float64(r.endTime.Sub(r.startTime).Milliseconds())),
			},
		},
	}
	_, err := c.client.PutMetricData(input)
	if err != nil {
		c.logger.
			WithError(err).
			Error("Error putting metric data")
	}
	return err
}
