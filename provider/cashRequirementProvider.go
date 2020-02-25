package provider

import (
	models "capital-adequacy/models"
)

//ThresholdValue : Threshold value for the job run
// const ThresholdValue int = 30

//GetLatest :
func GetLatest() models.CashRequirement {
	var cashRequirement models.CashRequirement
	cashRequirement = models.GetCashRequirement()
	//fmt.Println(latestPositionSnapshot)
	return cashRequirement

}


