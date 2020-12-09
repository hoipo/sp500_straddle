package models

import "dao"

// HkETFData 香港ETF数据结构
type HkETFData struct {
	ID    int     `json:"id"`
	Price float32 `json:"price"`
	Date  string  `json:"date"`
	Name  string  `json:name`
}

// CreateHkETF : Create a new record in db
func CreateHkETF(HkETF *HkETFData) (err error) {
	err = dao.DB.Create(&HkETF).Error
	return
}

// GetHkETF : Get the HkETF data from db
func GetHkETF(limit int64) (HkETF []*HkETFData, err error) {
	if err = dao.DB.Order("id desc").Limit(int(limit)).Find(&HkETF).Error; err != nil {
		return nil, err
	}
	return
}
