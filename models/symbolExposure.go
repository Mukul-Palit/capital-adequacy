package models

import (
	database "airflow-report/capital-adequacy/driver"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

type Sym interface {
	GetSymExposure() (SymExposure, error)
	GetSumRealisedPnl() (float64, error)
}

type Symbol struct {
	DB *gorm.DB
}

//SymExposure : Structure for the symbol exposure table
type SymExposure struct {
	ID          string  `gorm:"type:bigint;PRIMARY_KEY" json:"id"`
	Timestamp   string  `gorm:"type:datetime;NOT NULL" json:"ts"`
	Sym         string  `gorm:"type:varchar(255);NOT NULL" json:"sym"`
	CpType      string  `gorm:"type:text;NOT NULL" json:"cptype"`
	Exposure    float64 `gorm:"type:float;NOT NULL" json:"exposure"`
	DayPnl      float64 `gorm:"type:float;NOT NULL" json:"daypnl"`
	RealisedPnl float64 `gorm:"type:float;NOT NULL" json:"realisedpnl"`
	//CreatedAt   string  `gorm:"type:datetime;NOT NULL" json:"created_at"`
}

//GetSymExposure :
func (s *Symbol) GetSymExposure() (SymExposure, error) {
	var symExposure SymExposure
	var err error
	var id, timestamp, sym, cptype string
	var exposure, daypnl, realisedpnl float64
	row, err := s.DB.Table("sym_exposure").Limit(1).Where("ts = (?)", s.DB.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Select("id,ts,sym,cptype,exposure,daypnl,realisedpnl").Rows()
	defer row.Close()
	if err != nil {
		database.WriteLogFile(err)
	}
	for row.Next() {
		row.Scan(&id, &timestamp, &sym, &cptype, &exposure, &daypnl, &realisedpnl)
	}
	symExposure = SymExposure{ID: id, Timestamp: timestamp, Sym: sym, CpType: cptype, Exposure: exposure, DayPnl: daypnl, RealisedPnl: realisedpnl}
	return symExposure, err
}

//GetSumRealisedPnl : returns the sum of realisedpnl
func (s *Symbol) GetSumRealisedPnl() (float64, error) {
	var err error
	var sum float64
	row, err := s.DB.Table("sym_exposure").Where("ts = (?)", s.DB.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Select("Sum(realisedpnl)").Rows()
	if err != nil {
		database.WriteLogFile(err)
	}
	for row.Next() {
		row.Scan(&sum)
	}
	return sum, err
}

func CreateSymExposure(db *gorm.DB) Sym {
	return &Symbol{
		DB: db,
	}
}
