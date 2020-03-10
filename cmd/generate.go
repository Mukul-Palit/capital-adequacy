package cmd

import (
	database "airflow-report/capital-adequacy/driver"
	handle "airflow-report/capital-adequacy/handler"
	metrics "airflow-report/capital-adequacy/metrics"
	models "airflow-report/capital-adequacy/models"
	provider "airflow-report/capital-adequacy/provider"
	"math"
	"time"

	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "It generates the capital adequacy report",
	Long: `This command takes input from three tables i.e. SymbolExposure,
	Var and CashRequirement and then calculates the values for capital adequacy table`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		start := time.Now()
		metrics.NewMonitor()
		color.New(color.FgWhite, color.BgGreen).Println("Start Time: ", start)
		currentRealisedPnl := provider.GetCurrentRealisedPnL()
		positionSnapshot := provider.GetCurrentPositionSnapshot()
		latestCashRequirement := provider.GetLatest()
		generateCapitalAdequacy(currentRealisedPnl, positionSnapshot, latestCashRequirement)
		metrics.NewMonitor()
		metrics.ElapsedTime(start)
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
}

func generateCapitalAdequacy(currentRealisedPnl models.Money, positionSnapshot models.PositionSnapshot, latestCashRequirement models.CashRequirement) {
	var capitalAdequacy models.CapitalAdequacy
	var netOpenPositonUsd models.Money
	var ValueAtRiskUsd models.Money

	db := database.DbConn()
	m := models.CreateMt4Prices(db)
	bid, _, err := m.GetExchangeRatio()
	if err != nil {
		database.WriteLogFile(err)
	}
	cashFloor := latestCashRequirement.CashRequirement + math.Round(latestCashRequirement.CashRequirement*1.5)
	exposureCentsUsd := (int)(positionSnapshot.Exposure * 100)
	netOpenPositonUsd.Amount = exposureCentsUsd
	netOpenPositonUsd.Currency = provider.CURRENCY
	netOpenPositonAud := (float64)(netOpenPositonUsd.Amount) * bid
	varCents := (int)(positionSnapshot.Var * 100)
	ValueAtRiskUsd.Amount = varCents
	ValueAtRiskUsd.Currency = provider.CURRENCY
	valueAtRiskAud := (float64)(ValueAtRiskUsd.Amount) * bid
	realisedPnlAud := (float64)(currentRealisedPnl.Amount) * bid
	totalRequiredCapital := valueAtRiskAud + cashFloor
	forecastedRequiredCapital := latestCashRequirement.MaxValueAtRisk + cashFloor

	realtimeTotalCash := realisedPnlAud + latestCashRequirement.TotalCash
	cashExcessCurrent := realtimeTotalCash - totalRequiredCapital
	cashForecastedExcess := realtimeTotalCash - forecastedRequiredCapital
	capitalAdequacy.CashRequirementID = latestCashRequirement.ID
	capitalAdequacy.CashBuffer = math.Round(latestCashRequirement.CashRequirement * 1.5)
	capitalAdequacy.CashFloor = cashFloor
	capitalAdequacy.RequiredCapitalTotal = totalRequiredCapital
	capitalAdequacy.RequiredCapitalForecasted = forecastedRequiredCapital
	capitalAdequacy.TotalCashRealtime = realtimeTotalCash
	capitalAdequacy.CashExcessCurrent = cashExcessCurrent
	capitalAdequacy.CashExcessForecasted = cashForecastedExcess
	capitalAdequacy.NetPositionAUD = netOpenPositonAud
	capitalAdequacy.ValueAtRiskAUD = valueAtRiskAud
	capitalAdequacy.RealisedPnlAUD = realisedPnlAud
	capitalAdequacy.TapaasReferenceID = positionSnapshot.ID
	capitalAdequacy.CreatedAt = time.Now().Format("2006-01-02T15:04:05")
	db = database.DbConn()
	c := models.CreateCapitalAdequacy(db)
	err = c.SetCapitalAdequacy(capitalAdequacy)
	if err != nil {
		database.WriteLogFile(err)
	}
	handle.PutmetricData("ValueAtRisk", "Count", capitalAdequacy.ValueAtRiskAUD, "Value at risk", "Value at risk data")

}
