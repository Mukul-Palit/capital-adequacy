package provider

import (
	database "airflow-report/capital-adequacy/driver"
	metrics "airflow-report/capital-adequacy/metrics"
	models "airflow-report/capital-adequacy/models"
	"time"
)

//ThresholdValue : Threshold value for the job run
const ThresholdValue int = 30

//CURRENCY : constant currency value as USD
const CURRENCY = "USD"

//GetCurrentRealisedPnL :
func GetCurrentRealisedPnL() models.Money {
	var symbolExposure models.SymExposure
	db := database.DbConn()
	s := models.CreateSymExposure(db)
	symbolExposure, err := s.GetSymExposure()
	symbolExposureAge(symbolExposure.Timestamp)
	if err != nil {
		database.WriteLogFile(err)
	}
	realisedPnlFloat, err := s.GetSumRealisedPnl()
	if err != nil {
		database.WriteLogFile(err)
	}

	realisedPnlMoney := convertRealisedPnlFloatToMoney(realisedPnlFloat)
	return realisedPnlMoney
}

func convertRealisedPnlFloatToMoney(realisedPnlFloat float64) models.Money {
	realisedPnlIntegerCents := (int)(realisedPnlFloat * 100)
	realisedPnlMoney := models.Money{Amount: realisedPnlIntegerCents, Currency: CURRENCY}
	return realisedPnlMoney
}

func symbolExposureAge(ts string) int {
	symbolExposureTimeStamp, err := time.Parse("2006-01-02 15:04:05", ts)

	if err != nil {
		database.WriteLogFile(err)
	}
	exposureAge := time.Since(symbolExposureTimeStamp)
	age := (int)(exposureAge.Hours() * 0.041666667)
	metrics.CalculateSymbolAge(age)
	return age
}
