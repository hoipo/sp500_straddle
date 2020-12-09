package models

import "dao"

// LofData lof的数据结构
type LofData struct {
	ID        int     `json:"id"`
	Open      float32 `json:"open"`
	Close     float32 `json:"close"`
	Last      float32 `json:"last"`
	Highest   float32 `json:"highest"`
	Lowest    float32 `json:"lowest"`
	Bid       float32 `json:"bid"`
	Ask       float32 `json:"ask"`
	Date      string  `json:"date"`
	Time      string  `json:"time"`
	Value     float32 `json:"value"`
	ValueDate string  `json:"value_date"`
	Name      string  `json:name`
}

// CreateLof : Create a new record in db
func CreateLof(Lof *LofData) (err error) {
	err = dao.DB.Create(&Lof).Error
	return
}

// GetLof : Get the Lof data from db
func GetLof(limit int64) (Lof []*LofData, err error) {
	if err = dao.DB.Order("id desc").Limit(int(limit)).Find(&Lof).Error; err != nil {
		return nil, err
	}
	return
}
