package doer

import "github.com/aws/aws-sdk-go/service/cloudwatch"

//go:generate mockgen -destination=../mocks/mock_doer.go -package=mocks testing-with-gomock/doer Doer

type Doer interface {
	PutMetricData(input *cloudwatch.PutMetricDataInput) error
}
