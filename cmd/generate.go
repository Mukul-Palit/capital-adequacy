// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	metrices "capital-adequacy/metrices"
	models "capital-adequacy/models"
	provider "capital-adequacy/provider"
	"fmt"
	"math"
	"time"

	"github.com/shopspring/decimal"

	"github.com/disiqueira/gocurrency"
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
		metrices.NewMonitor()
		color.New(color.FgWhite, color.BgGreen).Println("Start Time: ", start)
		currentRealisedPnl := provider.GetCurrentRealisedPnL()
		//pos = models.GetPositionSnapshot()
		fmt.Println(currentRealisedPnl.Amount)
		fmt.Println(currentRealisedPnl.Currency)
		positionSnapshot := provider.GetCurrentPositionSnapshot()
		fmt.Println(positionSnapshot)
		latestCashRequirement := provider.GetLatest()
		fmt.Println(latestCashRequirement)
		generateCapitalAdequacy(currentRealisedPnl, positionSnapshot, latestCashRequirement)
		metrices.NewMonitor()
		metrices.ElapsedTime(start)
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
}

func generateCapitalAdequacy(currentRealisedPnl models.Money, positionSnapshot models.PositionSnapshot, latestCashRequirement models.CashRequirement) {
	var desiredCURRENCY = gocurrency.NewCurrency("AUD")
	var dollar = gocurrency.NewCurrency("USD")
	var capitalAdequacy models.CapitalAdequacy
	var netOpenPositonUsd models.Money
	var ValueAtRiskUsd models.Money
	cashFloor := latestCashRequirement.CashRequirement + math.Round(latestCashRequirement.CashRequirement*1.5)
	exposureCentsUsd := (int)(positionSnapshot.Exposure * 100)
	netOpenPositonUsd.Amount = exposureCentsUsd
	netOpenPositonUsd.Currency = provider.CURRENCY
	netOpenPositonAud, err := gocurrency.ConvertCurrency(dollar, desiredCURRENCY, decimal.NewFromFloat((float64)(netOpenPositonUsd.Amount)))
	if err != nil {
		fmt.Println(err)
	}
	varCents := (int)(positionSnapshot.Var * 100)
	ValueAtRiskUsd.Amount = varCents
	ValueAtRiskUsd.Currency = provider.CURRENCY
	valueAtRiskAud, err := gocurrency.ConvertCurrency(dollar, desiredCURRENCY, decimal.NewFromFloat((float64)(ValueAtRiskUsd.Amount)))
	if err != nil {
		fmt.Println(err)
	}
	realisedPnlAud, err := gocurrency.ConvertCurrency(dollar, desiredCURRENCY, decimal.NewFromFloat((float64)(currentRealisedPnl.Amount)))
	if err != nil {
		fmt.Println(err)
	}
	totalRequiredCapital := valueAtRiskAud.Add(decimal.NewFromFloat(cashFloor))
	forecastedRequiredCapital := decimal.NewFromFloat(latestCashRequirement.MaxValueAtRisk + cashFloor)

	realtimeTotalCash := realisedPnlAud.Add(decimal.NewFromFloat(latestCashRequirement.TotalCash))
	cashExcessCurrent := realtimeTotalCash.Sub(totalRequiredCapital)
	cashForecastedExcess := realtimeTotalCash.Sub(forecastedRequiredCapital)
	capitalAdequacy.CashRequirementID = latestCashRequirement.ID
	capitalAdequacy.CashBuffer = decimal.NewFromFloat(math.Round(latestCashRequirement.CashRequirement * 1.5))
	capitalAdequacy.CashFloor = decimal.NewFromFloat(cashFloor)
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
	models.SetCapitalAdequacy(capitalAdequacy)
}
