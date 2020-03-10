package test

import (
	"airflow-report/capital-adequacy/testing-with-gomock/doer"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

//structure containing function to be mocked
type User struct {
	Doer doer.Doer
}

// Function which is mocked
func (u *User) Putmetricdata() error {
	return u.Doer.PutMetricData(&cloudwatch.PutMetricDataInput{
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
	})

}
