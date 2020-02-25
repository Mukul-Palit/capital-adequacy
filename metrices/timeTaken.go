package metrices

import (
	"time"

	"github.com/gookit/color"
)

type Time struct {
	Start       time.Time
	End         time.Time
	ElapsedTime time.Duration
}

func ElapsedTime(start time.Time) {
	var calculateTime Time
	calculateTime.Start = start
	calculateTime.End = time.Now()
	color.New(color.FgWhite, color.BgGreen).Println("End Time: ", calculateTime.End)
	calculateTime.ElapsedTime = time.Since(calculateTime.Start)
	color.New(color.FgGreen, color.BgWhite).Println("Time Taken: ", calculateTime.ElapsedTime)
}
