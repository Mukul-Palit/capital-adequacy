package models

import (
	database "capital-adequacy/driver"
	"fmt"

	"github.com/shopspring/decimal"

	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

//CapitalAdequacy : Structure for the capital_adequacy table
type CapitalAdequacy struct {
	ID                        string          `gorm:"type:bigint(8);PRIMARY_KEY"`
	CashRequirementID         string          `gorm:"type:bigint(8);NOT NULL"`
	NetPositionAUD            decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	RealisedPnlAUD            decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	ValueAtRiskAUD            decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	CashBuffer                decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	RequiredCapitalTotal      decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	RequiredCapitalForecasted decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	TotalCashRealtime         decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	CashExcessCurrent         decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	CashExcessForecasted      decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	CashFloor                 decimal.Decimal `gorm:"type:decimal(12,2);NOT NULL"`
	TapaasReferenceID         string          `gorm:"type:decimal(12,2);NOT NULL"`
	CreatedAt                 string          `gorm:"type:datetime;NOT NULL"`
}

//Money :
type Money struct {
	Amount   int
	Currency string
}

func getCapitalAdequacy() CapitalAdequacy {
	var capitalAdequacy CapitalAdequacy
	db := database.DbConn()
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	db.Order("timestamp desc").First(&capitalAdequacy)
	return capitalAdequacy
}

//SetCapitalAdequacy :
func SetCapitalAdequacy(capitalAdequacy CapitalAdequacy) {
	db := database.DbConn()
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	fmt.Println("inside set")
	fmt.Println(capitalAdequacy.TapaasReferenceID)
	db.Table("capital_adequacy").Create(&capitalAdequacy)
}
