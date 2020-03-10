package metrics

import (
	handle "airflow-report/capital-adequacy/handler"

	"github.com/gookit/color"
)

//SymbolAge : structure for the age of SymbolExposure
type SymbolAge struct {
	SymbolExposure int
}

//CalculateSymbolAge : to calculate the age of PositionSnapshot as a metric parameter
func CalculateSymbolAge(fetchAge int) {
	var age SymbolAge
	age.SymbolExposure = fetchAge
	color.New(color.FgWhite, color.BgGreen).Printf("Age : %d days\n", age.SymbolExposure)
	handle.PutmetricData("Symbol Exposure Age", "Count", float64(age.SymbolExposure), "Symbol Exposure Age", "Symbol Exposure data")
}
