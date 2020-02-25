package models

import (
	database "capital-adequacy/driver"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //blank
)

type Sym interface {
	GetSymExposure() (SymExposure, error)
	//GetSumRealisedPnl() (float64, error)
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
	Exposure    float64 `gorm:"type:decimal(12,2);NOT NULL" json:"exposure"`
	DayPnl      float64 `gorm:"type:decimal(12,2);NOT NULL" json:"daypnl"`
	RealisedPnl float64 `gorm:"type:decimal(12,2);NOT NULL" json:"realisedpnl"`
	CreatedAt   string  `gorm:"type:datetime;NOT NULL" json:"created_at"`
}

//GetSymExposure :
func (s *Symbol) GetSymExposure() (SymExposure, error) {
	//var symExposure []SymExposure
	var symExposure SymExposure
	//db = database.DbConn()
	var err error
	// defer func() {
	// 	err = s.DB.Close()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }()
	//	var id, timestamp, sym, cptype string
	//	var exposure, daypnl, realisedpnl float64
	//fmt.Println(s.DB.Table("sym_exposure").Limit(1).Where("ts = (?)", s.DB.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Select("id,ts,sym,cptype,exposure,daypnl,realisedpnl").Take(&symExposure).QueryExpr())
	//db := database.DbConn()
	//	row, err := s.DB.Table("sym_exposure").Limit(1).Where("ts = (?)", s.DB.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Select("id,ts,sym,cptype,exposure,daypnl,realisedpnl").Scan(&symExposure).Rows()
	// fmt.Println(row)
	// if err != nil {
	// 	fmt.Println("Hello")
	// 	fmt.Println(err)
	// }
	//defer row.Close()
	err = s.DB.Table("sym_exposure").Limit(1).Select("id,ts,sym,cptype,exposure,daypnl,realisedpnl").Where("ts = (?)", s.DB.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Scan(&symExposure).Error
	if err != nil {
		fmt.Println("Hello")
		fmt.Println(err)
	}
	//fmt.Println(s.DB.Table("sym_exposure").Where("ts = (?)", s.DB.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Select("id,ts,sym,cptype,exposure,daypnl,realisedpnl").First(&symExposure).QueryExpr())
	// for row.Next() {
	// 	row.Scan(&id, &timestamp, &sym, &cptype, &exposure, &daypnl, &realisedpnl)
	// }
	// symExposure = SymExposure{ID: id, Timestamp: timestamp, Sym: sym, CpType: cptype, Exposure: exposure, DayPnl: daypnl, RealisedPnl: realisedpnl}
	fmt.Println(symExposure)
	return symExposure, err
}

//GetSumRealisedPnl : returns the sum of realisedpnl
func GetSumRealisedPnl() (float64, error) {
	db := database.DbConn()
	var err error
	defer func() {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	var sum float64
	//fmt.Println(s.DB.Table("sym_exposure").Where("ts = (?)", s.DB.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Select("Sum(realisedpnl)").QueryExpr())
	row := db.Table("sym_exposure").Where("ts = (?)", db.Table("sym_exposure").Select("MAX(ts)").QueryExpr()).Select("Sum(realisedpnl)").Row()
	row.Scan(&sum)
	return sum, err
}

func CreateSymExposure(db *gorm.DB) Sym {
	return &Symbol{
		DB: db,
	}
}
