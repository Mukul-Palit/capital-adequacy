package provider

import (
	models "airflow-report/capital-adequacy/models"
	"testing"
	"time"
)

func TestConvertRealisedPnlFloatToMoney(t *testing.T) {
	money := []models.Money{
		{
			Amount:   1556,
			Currency: "USD",
		},
		{
			Amount:   2348,
			Currency: "USD",
		},
	}
	testData := []struct {
		x float64
		y models.Money
	}{
		{15.563, money[0]},
		{23.481, money[1]},
	}
	for _, testDatum := range testData {
		monetPnl := convertRealisedPnlFloatToMoney(testDatum.x)
		if monetPnl != testDatum.y {
			t.Errorf("Output for %f was incorrect, got: %d and %s, want: %d and %s", testDatum.x, monetPnl.Amount, monetPnl.Currency, testDatum.y.Amount, testDatum.y.Currency)
		}

	}
}

func TestSymbolExposureAge(t *testing.T) {
	firstTimeString := "2020-02-25 10:10:10"
	secondTimeString := "2020-02-20 14:14:14"
	firstTime, _ := time.Parse("2006-01-02 15:04:05", firstTimeString)
	secondTime, _ := time.Parse("2006-01-02 15:04:05", secondTimeString)
	firstTimeDuration := (int)(time.Since(firstTime).Hours() * 0.041666667)
	secondTimeDuration := (int)(time.Since(secondTime).Hours() * 0.041666667)
	testData := []struct {
		x string
		y int
	}{
		{firstTimeString, firstTimeDuration},
		{secondTimeString, secondTimeDuration},
	}
	for _, testDatum := range testData {
		testTime := symbolExposureAge(testDatum.x)
		if testTime != testDatum.y {
			t.Errorf("Output for %s was incorrect, got: %d , want: %d", testDatum.x, testTime, testDatum.y)
		}

	}
}
