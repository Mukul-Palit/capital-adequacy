package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

type Capital interface {
	SetCapitalAdequacy(capitalAdequacy CapitalAdequacy) error
}

type CapitalAdeq struct {
	DB *gorm.DB
}

//CapitalAdequacy : Structure for the capital_adequacy table
type CapitalAdequacy struct {
	ID                        string  `gorm:"type:bigint(8);PRIMARY_KEY"`
	CashRequirementID         string  `gorm:"type:bigint(8);NOT NULL"`
	NetPositionAUD            float64 `gorm:"type:decimal(12,2);NOT NULL"`
	RealisedPnlAUD            float64 `gorm:"type:decimal(12,2);NOT NULL"`
	ValueAtRiskAUD            float64 `gorm:"type:decimal(12,2);NOT NULL"`
	CashBuffer                float64 `gorm:"type:decimal(12,2);NOT NULL"`
	RequiredCapitalTotal      float64 `gorm:"type:decimal(12,2);NOT NULL"`
	RequiredCapitalForecasted float64 `gorm:"type:decimal(12,2);NOT NULL"`
	TotalCashRealtime         float64 `gorm:"type:decimal(12,2);NOT NULL"`
	CashExcessCurrent         float64 `gorm:"type:decimal(12,2);NOT NULL"`
	CashExcessForecasted      float64 `gorm:"type:decimal(12,2);NOT NULL"`
	CashFloor                 float64 `gorm:"type:decimal(12,2);NOT NULL"`
	TapaasReferenceID         string  `gorm:"type:decimal(12,2);NOT NULL"`
	CreatedAt                 string  `gorm:"type:datetime;NOT NULL"`
}

//Money :
type Money struct {
	Amount   int
	Currency string
}

// func getCapitalAdequacy() CapitalAdequacy {
// 	var capitalAdequacy CapitalAdequacy
// 	db := database.DbConn()
// 	defer func() {
// 		err := db.Close()
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}()
// 	db.Order("timestamp desc").First(&capitalAdequacy)
// 	return capitalAdequacy
// }

//SetCapitalAdequacy :
func (c *CapitalAdeq) SetCapitalAdequacy(capitalAdequacy CapitalAdequacy) error {
	return c.DB.Table("capital_adequacy").Create(&capitalAdequacy).Error
}

func CreateCapitalAdequacy(db *gorm.DB) Capital {
	return &CapitalAdeq{
		DB: db,
	}
}
