package main

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

// CREATE TYPE car_type AS ENUM (
// 'SEDAN',
// 'HATCHBACK',
// 'MINIVAN');

type carType string

const (
	SEDAN     carType = "SEDAN"
	HATCHBACK carType = "HATCHBACK"
	MINIVAN   carType = "MINIVAN"
)

func (ct *carType) Scan(value interface{}) error {
	*ct = carType(value.([]byte))
	return nil
}

func (ct carType) Value() (driver.Value, error) {
	return string(ct), nil
}

type MyTable struct {
	gorm.Model
	CarType carType `sql:"car_type"`
}

func (MyTable) TableName() string {
	return "my_table"
}
