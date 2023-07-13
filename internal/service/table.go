package service

import (
	"echarts_example/internal/repository"
	"echarts_example/internal/schema"
)

func GetOneTable() (schema.Echarts, error) {
	echarts := &schema.Echarts{}
	err := repository.PG.Model(echarts).Where("shirt = ?", 5).Select()
	if err != nil {
		return schema.Echarts{}, err
	}
	return *echarts, err
}

func GetAllTable() ([]schema.Echarts, error) {
	echarts := &[]schema.Echarts{}
	err := repository.PG.Model(echarts).Select()
	if err != nil {
		return nil, err
	}
	return *echarts, err
}

func CreateTable(createData schema.Echarts) (any, error) {
	//func CreateTable() (any, error) {
	//	createData := schema.Echarts{
	//		Shirt:    1,
	//		Cardigan: 3,
	//		Chiffon:  3,
	//		Pants:    20,
	//		Heels:    20,
	//		Socks:    40,
	//	}

	result, err := repository.PG.Model(&createData).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return result, err
}
