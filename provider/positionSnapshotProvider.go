package provider

import (
	models "capital-adequacy/models"
	"fmt"
	"time"
)

//ThresholdValue : Threshold value for the job run
// const ThresholdValue int = 30

//GetCurrentPositionSnapshot :
func GetCurrentPositionSnapshot() models.PositionSnapshot {
	var positionSnapshot models.PositionSnapshot
	positionSnapshot = models.GetPositionSnapshot()
	if isPositionSnapshotFresh(positionSnapshot.Timestamp) {
		//fmt.Println("Hello")
	}
	//fmt.Println(latestPositionSnapshot)
	return positionSnapshot

}
func isPositionSnapshotFresh(ts string) bool {
	positionSnapshotTimeStamp, err := time.Parse("2006-01-02 15:04:05", ts)
	if err != nil {
		fmt.Println(err)
	}
	thresholdDateTime := time.Now().Add(time.Duration(-ThresholdValue) * time.Minute)
	//fmt.Println(thresholdDateTime)
	return positionSnapshotTimeStamp.After(thresholdDateTime)
}
