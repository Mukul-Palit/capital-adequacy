package test

import (
	"airflow-report/capital-adequacy/testing-with-gomock/mocks"
	"errors"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
)

func TestPutmetricdata(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDoer := mocks.NewMockDoer(mockCtrl)
	mockDoer.EXPECT().PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Site/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("UniqueVisits"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(18.90),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("keyoftest"),
						Value: aws.String("valueoftest"),
					},
				},
			},
		},
	}).Return(nil).Times(1)
	var testuser error
	assert.Equal(t, testuser, mockDoer.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Site/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("UniqueVisits"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(18.90),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("keyoftest"),
						Value: aws.String("valueoftest"),
					},
				},
			},
		},
	}))
}
func TestPutmetricdataForDummyError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	dummyError := errors.New("dummy error")
	mockDoer := mocks.NewMockDoer(mockCtrl)
	mockDoer.EXPECT().PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Site/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("UniqueVisits"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(18.90),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("keyoftest"),
						Value: aws.String("valueoftest"),
					},
				},
			},
		},
	}).Return(dummyError).Times(1)
	var testuser error
	testuser = dummyError
	assert.Equal(t, testuser, mockDoer.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Site/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("UniqueVisits"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(18.90),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("keyoftest"),
						Value: aws.String("valueoftest"),
					},
				},
			},
		},
	}))
}
func TestPutmetricdataForAnyInput(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockDoer := mocks.NewMockDoer(mockCtrl)
	mockDoer.EXPECT().PutMetricData(gomock.Any())
	var testuser error
	assert.Equal(t, testuser, mockDoer.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Site/Traffic"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String("UniqueVisits"),
				Unit:       aws.String("Count"),
				Value:      aws.Float64(18.90),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String("keyoftest"),
						Value: aws.String("valueoftest"),
					},
				},
			},
		},
	}))
}
