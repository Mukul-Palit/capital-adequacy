package provider

import (
	database "capital-adequacy/driver"
	models "capital-adequacy/models"
	"fmt"
	"math"
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
	if isSymbolExposureFresh(symbolExposure.Timestamp) {
		fmt.Println("Hello")
	}
	if err != nil {
		fmt.Println(err)
	}
	db = database.DbConn()
	s = models.CreateSymExposure(db)
	realisedPnlFloat, err := models.GetSumRealisedPnl()
	realisedPnlMoney := convertRealisedPnlFloatToMoney(realisedPnlFloat)
	return realisedPnlMoney
}

func convertRealisedPnlFloatToMoney(realisedPnlFloat float64) models.Money {
	realisedPnlIntegerCents := (int)(realisedPnlFloat * 100)
	realisedPnlMoney := models.Money{Amount: realisedPnlIntegerCents, Currency: CURRENCY}
	return realisedPnlMoney
}

func isSymbolExposureFresh(ts string) bool {
	//symbolExposureTimeStamp := ts
	symbolExposureTimeStamp, err := time.Parse("2006-01-02 15:04:05", ts)

	if err != nil {
		fmt.Println(err)
	}
	thresholdDateTime := time.Now().Add(time.Duration(-ThresholdValue) * time.Minute)
	//fmt.Println(thresholdDateTime)
	agee := time.Since(symbolExposureTimeStamp)
	age := (int)(math.Round(agee.Hours() * 0.041666667))
	fmt.Printf("Age : %d days\n", age)
	return symbolExposureTimeStamp.After(thresholdDateTime)
}
