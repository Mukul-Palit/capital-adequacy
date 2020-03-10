package provider

import (
	"testing"
	"time"
)

func TestPositionSnapshotAge(t *testing.T) {
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
		testTime := positionSnapshotAge(testDatum.x)
		if testTime != testDatum.y {
			t.Errorf("Output for %s was incorrect, got: %d , want: %d", testDatum.x, testTime, testDatum.y)
		}

	}
}
