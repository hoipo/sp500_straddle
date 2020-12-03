package models

import (
	"dao"
)

// FutureData 期货的数据结构
type FutureData struct {
	ID         int     `json:"id"`
	Last       float32 `json:"last"`
	Bid        float32 `json:"bid"`
	Ask        float32 `json:"ask"`
	Highest    float32 `json:"highest"`
	Lowest     float32 `json:"lowest"`
	Settle     float32 `json:"settle"`
	Open       float32 `json:"open"`
	HoldAmount int32   `json:"hold_amount"`
	Time       string  `json:"time"`
	Date       string  `json:"date"`
}

func CreateFuture(future *FutureData) (err error) {
	err = dao.DB.Create(&future).Error
	return
}

func GetFuture(limit int64) (futures []*FutureData, err error) {
	if err = dao.DB.Order("id desc").Limit(int(limit)).Find(&futures).Error; err != nil {
		return nil, err
	}
	return
}
