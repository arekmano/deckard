package reporter

import (
	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/stats"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/pkg/errors"
)

type DynamoDBReporter struct {
	client               dynamodbiface.DynamoDBAPI
	transactionTableName string
	reportTableName      string
}

func (c *DynamoDBReporter) ReportTransaction(r *executor.Report) error {
	item, err := dynamodbattribute.MarshalMap(r)
	if err != nil {
		return errors.Wrap(err, "Could not marshal the transaction")
	}
	c.client.PutItemRequest(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &c.transactionTableName,
	})
	return nil
}

func (c *DynamoDBReporter) ReportStatistics(r *stats.Statistics) error {
	item, err := dynamodbattribute.MarshalMap(r)
	if err != nil {
		return errors.Wrap(err, "Could not marshal the statistics")
	}
	c.client.PutItemRequest(&dynamodb.PutItemInput{
		Item:      item,
		TableName: &c.transactionTableName,
	})
	return nil
}
