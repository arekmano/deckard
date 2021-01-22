package stats

import (
	"time"

	"github.com/arekmano/deckard/executor"
	"github.com/arekmano/deckard/transaction"
	"github.com/montanaflynn/stats"
)

type Statistics struct {
	StartTime                       time.Time
	EndTime                         time.Time
	MeanExecutionTime               float64
	P90ExecutionTime                float64
	P99ExecutionTime                float64
	MaxExecutionTime                float64
	MinExecutionTime                float64
	SuccessfulTransactionCount      int64
	FailedTransactionCount          int64
	StartedTransactionCount         int64
	TransactionCount                int64
	TotalTransactionCount           int64
	TotalEndedTransactionCount      int64
	TotalFailedTransactionCount     int64
	TotalSuccessfulTransactionCount int64
	TotalStartedTransactionCount    int64
}

func Stats(reports []executor.Report, starts []time.Time) (*Statistics, error) {
	var fail int64
	var success int64
	var executionTimes []float64
	startTime := time.Unix(1<<62-1, 0)
	if len(starts) > 0 {
		startTime = starts[0]
	}
	endTime := time.Unix(0, 0)
	for _, report := range reports {
		if endTime.Before(report.EndTime) {
			endTime = report.EndTime
		}
		executionTimes = append(executionTimes, float64(report.EndTime.Sub(report.StartTime).Milliseconds()))
		if report.Status == transaction.Success {
			success++
		} else {
			fail++
		}
	}
	mean, err := stats.Mean(executionTimes)
	if err != nil {
		return nil, err
	}
	p90, err := stats.Percentile(executionTimes, 90)
	if err != nil {
		return nil, err
	}
	p99, err := stats.Percentile(executionTimes, 99)
	if err != nil {
		return nil, err
	}
	max, err := stats.Max(executionTimes)
	if err != nil {
		return nil, err
	}
	min, err := stats.Min(executionTimes)
	if err != nil {
		return nil, err
	}
	return &Statistics{
		MeanExecutionTime:               mean,
		P90ExecutionTime:                p90,
		P99ExecutionTime:                p99,
		MaxExecutionTime:                max,
		MinExecutionTime:                min,
		SuccessfulTransactionCount:      success,
		FailedTransactionCount:          fail,
		StartedTransactionCount:         int64(len(starts)),
		TransactionCount:                success + fail,
		TotalTransactionCount:           int64(len(reports)),
		TotalEndedTransactionCount:      int64(len(reports)),
		StartTime:                       startTime,
		EndTime:                         endTime,
		TotalFailedTransactionCount:     fail,
		TotalSuccessfulTransactionCount: success,
		TotalStartedTransactionCount:    int64(len(starts)),
	}, nil
}

func StatsBetween(reports []executor.Report, starts []time.Time, after time.Time, before time.Time) (*Statistics, error) {
	var executionTimesBetween stats.Float64Data
	var successBetween int64
	var failBetween int64
	var startsBetween int64
	var fail int64
	var success int64
	startTime := time.Unix(1<<62-1, 0)
	if len(starts) > 0 {
		startTime = starts[0]
	}
	for _, start := range starts {
		if start.After(after) && start.Before(before) {
			startsBetween++
		}
	}
	endTime := time.Unix(0, 0)
	for _, report := range reports {
		if endTime.Before(report.EndTime) {
			endTime = report.EndTime
		}
		if report.EndTime.After(after) && report.EndTime.Before(before) {
			executionTimesBetween = append(executionTimesBetween, float64(report.EndTime.Sub(report.StartTime).Milliseconds()))
			if report.Status == transaction.Success {
				successBetween++
			} else {
				failBetween++
			}
		} else {
			if report.Status == transaction.Success {
				success++
			} else {
				fail++
			}
		}
	}
	meanBetween, err := stats.Mean(executionTimesBetween)
	if err != nil {
		return nil, err
	}
	p90Between, err := stats.Percentile(executionTimesBetween, 90)
	if err != nil {
		return nil, err
	}
	p99Between, err := stats.Percentile(executionTimesBetween, 99)
	if err != nil {
		return nil, err
	}
	maxBetween, err := stats.Max(executionTimesBetween)
	if err != nil {
		return nil, err
	}
	minBetween, err := stats.Min(executionTimesBetween)
	if err != nil {
		return nil, err
	}
	return &Statistics{
		MeanExecutionTime:               meanBetween,
		P90ExecutionTime:                p90Between,
		P99ExecutionTime:                p99Between,
		MaxExecutionTime:                maxBetween,
		MinExecutionTime:                minBetween,
		SuccessfulTransactionCount:      successBetween,
		FailedTransactionCount:          failBetween,
		StartedTransactionCount:         startsBetween,
		TransactionCount:                successBetween + failBetween,
		TotalTransactionCount:           int64(len(reports)),
		TotalEndedTransactionCount:      int64(len(reports)),
		StartTime:                       startTime,
		EndTime:                         endTime,
		TotalFailedTransactionCount:     fail,
		TotalSuccessfulTransactionCount: success,
		TotalStartedTransactionCount:    int64(len(starts)),
	}, nil
}
