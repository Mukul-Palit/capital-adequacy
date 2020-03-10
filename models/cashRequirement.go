package models

import (
	database "airflow-report/capital-adequacy/driver"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

type Cash interface {
	GetCashRequirement() (CashRequirement, error)
}

type CashReq struct {
	DB *gorm.DB
}

//CashRequirement : Structure for the cash_requirement table
type CashRequirement struct {
	ID              string  `gorm:"type:bigint(8);PRIMARY_KEY"`
	CashRequirement float64 `gorm:"type:decimal(12,2);NOT NULL"`
	MaxValueAtRisk  float64 `gorm:"type:decimal(12,2);NOT NULL"`
	PrimeBrokerCash float64 `gorm:"type:decimal(12,2);NOT NULL"`
	TotalCash       float64 `gorm:"type:decimal(12,2);NOT NULL"`
}

//GetCashRequirement : returns the latest entry in the table
func (c *CashReq) GetCashRequirement() (CashRequirement, error) {
	var cashRequirement CashRequirement
	var id string
	var cashrequirement, maxvalueatrisk, primebrokercash, totalcash float64
	row, err := c.DB.Table("cash_requirement").Order("id desc").Limit(1).Select("id,cash_requirement,max_value_at_risk,prime_broker_cash,total_cash").Rows()
	defer row.Close()
	if err != nil {
		database.WriteLogFile(err)
	}
	for row.Next() {
		row.Scan(&id, &cashrequirement, &maxvalueatrisk, &primebrokercash, &totalcash)
	}
	cashRequirement = CashRequirement{ID: id, CashRequirement: cashrequirement, MaxValueAtRisk: maxvalueatrisk, PrimeBrokerCash: primebrokercash, TotalCash: totalcash}
	return cashRequirement, err
}

func CreateCashRequirement(db *gorm.DB) Cash {
	return &CashReq{
		DB: db,
	}
}
