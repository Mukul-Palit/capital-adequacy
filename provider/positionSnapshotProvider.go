package provider

import (
	database "airflow-report/capital-adequacy/driver"
	metrics "airflow-report/capital-adequacy/metrics"
	models "airflow-report/capital-adequacy/models"
	"time"
)

//ThresholdValue : Threshold value for the job run
// const ThresholdValue int = 30

//GetCurrentPositionSnapshot :
func GetCurrentPositionSnapshot() models.PositionSnapshot {
	var positionSnapshot models.PositionSnapshot
	db := database.DbConn()
	p := models.CreatePositionSnapshot(db)
	positionSnapshot, err := p.GetPositionSnapshot()
	if err != nil {
		database.WriteLogFile(err)
	}
	positionSnapshotAge(positionSnapshot.Timestamp)
	return positionSnapshot

}
func positionSnapshotAge(ts string) int {
	positionSnapshotTimeStamp, err := time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		database.WriteLogFile(err)
	}
	//thresholdDateTime := time.Now().Add(time.Duration(-ThresholdValue) * time.Minute)
	positionAge := time.Since(positionSnapshotTimeStamp)
	age := (int)(positionAge.Hours() * 0.041666667)
	metrics.CalculatePositionAge(age)
	return age
}
