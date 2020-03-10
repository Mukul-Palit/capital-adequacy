package handler

import (
	database "airflow-report/capital-adequacy/driver"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
)

func PutmetricData(metricName string, unit string, value float64, DimensionName string, DimensionValue string) {
	// Initialize a session that the SDK uses to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and configuration from the shared configuration file ~/.aws/config.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-2")})
	if err != nil {
		database.WriteLogFile(err)
		return
	}
	// Create new cloudwatch client.
	svc := cloudwatch.New(sess)
	_, err = svc.PutMetricData(&cloudwatch.PutMetricDataInput{
		Namespace: aws.String("Capital/Adequacy"),
		MetricData: []*cloudwatch.MetricDatum{
			&cloudwatch.MetricDatum{
				MetricName: aws.String(metricName),
				Unit:       aws.String(unit),
				Value:      aws.Float64(value),
				Dimensions: []*cloudwatch.Dimension{
					&cloudwatch.Dimension{
						Name:  aws.String(DimensionName),
						Value: aws.String(DimensionValue),
					},
				},
			},
		},
	})
	if err != nil {
		// fmt.Println("inside handler")
		// fmt.Println(err)
		// database.WriteLogFile(err)
		return
	}
}
