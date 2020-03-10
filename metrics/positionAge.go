package metrics

import (
	handle "airflow-report/capital-adequacy/handler"

	"github.com/gookit/color"
)

//PositionAge : structure for the age of PositionSnapshot
type PositionAge struct {
	PositionSnapshot int
}

//CalculatePositionAge : to calculate the age of PositionSnapshot as a metric parameter
func CalculatePositionAge(fetchAge int) {
	var age PositionAge
	age.PositionSnapshot = fetchAge
	color.New(color.FgWhite, color.BgGreen).Printf("Age : %d days\n", age.PositionSnapshot)
	handle.PutmetricData("PositionSanpshot", "Count", float64(age.PositionSnapshot), "Position Sanpshot", "Position snapshot data")
}
