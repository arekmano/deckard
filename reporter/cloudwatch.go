package reporter

import (
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatch/cloudwatchiface"
)

type CloudwatchReporter struct {
	client cloudwatchiface.CloudWatchAPI
}

func (c *CloudwatchReporter) Report(r *Report) error {
	input := cloudwatch.PutMetricDataInput{}
	_, err := c.client.PutMetricData(&input)
	return err
}
