package models

import (
	database "airflow-report/capital-adequacy/driver"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

type Mt4 interface {
	GetExchangeRatio() (float64, float64, error)
}

type Mt4Price struct {
	DB *gorm.DB
}

//Mt4Prices : Structure for the mt4_price table
type Mt4Prices struct {
	Symbol     string  `gorm:"type:varchar(255);NOT NULL"`
	Time       string  `gorm:"type:datetime;NOT NULL"`
	Bid        float64 `gorm:"type:decimal(12,2);NOT NULL"`
	Ask        float64 `gorm:"type:decimal(12,2);NOT NULL"`
	Low        float64 `gorm:"type:decimal(12,2);NOT NULL"`
	High       float64 `gorm:"type:decimal(12,2);NOT NULL"`
	Direction  int     `gorm:"type:int;NOT NULL"`
	Digits     int     `gorm:"type:int;NOT NULL"`
	Spread     int     `gorm:"type:int;NOT NULL"`
	ModifyTime string  `gorm:"type:datetime;NOT NULL"`
}

func (m *Mt4Price) GetExchangeRatio() (float64, float64, error) {
	var bid, ask float64
	row, err := m.DB.Table("mt4_prices").Limit("1").Order("time desc").Where("symbol='AUDUSD'").Select("bid,ask").Rows()
	defer row.Close()
	if err != nil {
		database.WriteLogFile(err)
	}
	for row.Next() {
		row.Scan(&bid, &ask)
	}
	return bid, ask, err
}

func CreateMt4Prices(db *gorm.DB) Mt4 {
	return &Mt4Price{
		DB: db,
	}
}
