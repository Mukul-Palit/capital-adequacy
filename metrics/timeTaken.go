package metrics

import (
	handle "airflow-report/capital-adequacy/handler"
	"time"

	"github.com/gookit/color"
)

//Time : Time for execution
type Time struct {
	Start       time.Time
	End         time.Time
	ElapsedTime time.Duration
}

//ElapsedTime : to calculate the elapsed time for execution
func ElapsedTime(start time.Time) {
	var calculateTime Time
	calculateTime.Start = start
	calculateTime.End = time.Now()
	color.New(color.FgWhite, color.BgGreen).Println("End Time: ", calculateTime.End)
	calculateTime.ElapsedTime = time.Since(calculateTime.Start)
	color.New(color.FgGreen, color.BgWhite).Println("Time Taken: ", calculateTime.ElapsedTime)
	handle.PutmetricData("TimeOfCapital-Adequacy", "Count", float64(calculateTime.ElapsedTime), "Elapsed time for capital-adequacy", "Time-taken-by-Capital-Adequacy")
}
