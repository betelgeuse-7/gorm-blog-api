package models

import "gorm.io/gorm"

var db *gorm.DB

// assign global db variable in models package to dbase passed from main()
func RegisterDB(dbase *gorm.DB) {
	db = dbase
}

// return global db var
func GetDB() *gorm.DB {
	return db
}
