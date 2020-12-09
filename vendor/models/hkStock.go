package models

import "dao"

// HkStockData 港股结构体
type HkStockData struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Last    float32 `json:"last"`
	Open    float32 `json:"open"`
	Close   float32 `json:"close"`
	Highest float32 `json:"highest"`
	Lowest  float32 `json:"lowest"`
	Date    string  `json:"date"`
	Time    string  `json:"time"`
}

// CreateHkStock : Create a new record in db
func CreateHkStock(HkStock *HkStockData) (err error) {
	err = dao.DB.Create(&HkStock).Error
	return
}

// GetHkStock : Get the HkStock data from db
func GetHkStock(limit int64) (HkStock []*HkStockData, err error) {
	if err = dao.DB.Order("id desc").Limit(int(limit)).Find(&HkStock).Error; err != nil {
		return nil, err
	}
	return
}
