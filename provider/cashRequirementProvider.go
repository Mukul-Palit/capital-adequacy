package provider

import (
	database "airflow-report/capital-adequacy/driver"
	models "airflow-report/capital-adequacy/models"
)

//ThresholdValue : Threshold value for the job run
// const ThresholdValue int = 30

//GetLatest :
func GetLatest() models.CashRequirement {
	var cashRequirement models.CashRequirement
	db := database.DbConn()
	c := models.CreateCashRequirement(db)
	cashRequirement, err := c.GetCashRequirement()
	if err != nil {
		database.WriteLogFile(err)
	}
	return cashRequirement

}
