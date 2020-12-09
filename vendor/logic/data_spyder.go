package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"models"
	"myutils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//GetFuture : get the data of ES future
func GetFuture(symbol string) (models.FutureData, error) {
	res, err := http.DefaultClient.Get("https://hq.sinajs.cn/?_=" + strconv.FormatInt(time.Now().Unix(), 10) + "&list=hf_" + symbol)
	if err != nil {
		return models.FutureData{}, errors.New("Connection error")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.FutureData{}, errors.New("Read response body failed")
	}

	// 数据转化为UTF-8
	s, err := myutils.GbkToUtf8(body)
	if err != nil {
		return models.FutureData{}, errors.New("Transform to UTF-8 failed")
	}

	if strings.HasSuffix(string(s), "=\"\";\n") {
		return models.FutureData{}, errors.New("No data with the symbol " + symbol)
	}

	strArr := strings.Split(string(s), ",")

	last, _ := strconv.ParseFloat(strings.Split(strArr[0], "\"")[1], 32)
	bid, _ := strconv.ParseFloat(strArr[2], 32)
	ask, _ := strconv.ParseFloat(strArr[3], 32)
	highest, _ := strconv.ParseFloat(strArr[4], 32)
	lowest, _ := strconv.ParseFloat(strArr[5], 32)
	settle, _ := strconv.ParseFloat(strArr[7], 32)
	open, _ := strconv.ParseFloat(strArr[8], 32)
	holdAmount, _ := strconv.ParseFloat(strArr[9], 32)

	future := models.FutureData{
		Last:       float32(last),
		Bid:        float32(bid),
		Ask:        float32(ask),
		Highest:    float32(highest),
		Lowest:     float32(lowest),
		Settle:     float32(settle),
		Open:       float32(open),
		HoldAmount: int32(holdAmount),
		Time:       strArr[6],
		Date:       strArr[12],
		Name:       strArr[13],
	}

	return future, nil
}

//GetLOF : 获取国内LOF的数据
func GetLOF(symbol string) (models.LofData, error) {
	res, err := http.DefaultClient.Get("https://hq.sinajs.cn/?_=" + strconv.FormatInt(time.Now().Unix(), 10) + "&list=sz" + symbol + ",f_" + symbol)
	if err != nil {
		return models.LofData{}, errors.New("Connection error")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.LofData{}, errors.New("Read response body failed")
	}

	// 数据转化为UTF-8
	s, err := myutils.GbkToUtf8(body)
	if err != nil {
		return models.LofData{}, errors.New("Transform to UTF-8 failed")
	}

	if strings.HasSuffix(string(s), "=\"\";\n") {
		return models.LofData{}, errors.New("No data with the symbol " + symbol)
	}
	// 分割数据
	slideData := strings.Split(string(s), "\n")
	sz, f := slideData[0], slideData[1]
	// 先处理场内的数据
	szArr := strings.Split(sz, ",")

	open, _ := strconv.ParseFloat(szArr[1], 32)
	close, _ := strconv.ParseFloat(szArr[2], 32)
	last, _ := strconv.ParseFloat(szArr[3], 32)
	highest, _ := strconv.ParseFloat(szArr[4], 32)
	lowest, _ := strconv.ParseFloat(szArr[5], 32)
	bid, _ := strconv.ParseFloat(szArr[6], 32)
	ask, _ := strconv.ParseFloat(szArr[7], 32)
	date := szArr[30]
	time := szArr[31]
	name := strings.Split(szArr[0], "=\"")[1]

	// 处理基金的净值数据
	fArr := strings.Split(f, ",")
	value, _ := strconv.ParseFloat(fArr[1], 32)
	valueDate := fArr[4]

	return models.LofData{
		Open:      float32(open),
		Close:     float32(close),
		Last:      float32(last),
		Highest:   float32(highest),
		Lowest:    float32(lowest),
		Bid:       float32(bid),
		Ask:       float32(ask),
		Date:      date,
		Time:      time,
		Value:     float32(value),
		ValueDate: valueDate,
		Name:      name,
	}, nil
}

//GetHKStock : 获取国内LOF的数据
func GetHKStock(symbol string) (models.HkStockData, error) {
	res, err := http.DefaultClient.Get("https://hq.sinajs.cn/?_=" + strconv.FormatInt(time.Now().Unix(), 10) + "&list=rt_hk" + symbol)
	if err != nil {
		return models.HkStockData{}, errors.New("Connection error")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.HkStockData{}, errors.New("Read response body failed")
	}
	// 数据转化为UTF-8
	s, err := myutils.GbkToUtf8(body)
	if err != nil {
		return models.HkStockData{}, errors.New("Transform to UTF-8 failed")
	}

	if strings.HasSuffix(string(s), "=\"\";\n") {
		return models.HkStockData{}, errors.New("No data with the symbol " + symbol)
	}

	// 分割数据
	sArr := strings.Split(string(s), ",")

	name := sArr[1]
	open, _ := strconv.ParseFloat(sArr[2], 32)
	close, _ := strconv.ParseFloat(sArr[3], 32)
	highest, _ := strconv.ParseFloat(sArr[4], 32)
	lowest, _ := strconv.ParseFloat(sArr[5], 32)
	last, _ := strconv.ParseFloat(sArr[6], 32)
	date := sArr[17]
	time := sArr[18]

	return models.HkStockData{
		Last:    float32(last),
		Close:   float32(close),
		Highest: float32(highest),
		Lowest:  float32(lowest),
		Open:    float32(open),
		Name:    name,
		Date:    date,
		Time:    time,
	}, nil
}

//GetHkETF : 获取HK的ETF数据
func GetHkETF() (models.HkETFData, error) {
	now := time.Now()
	d, _ := time.ParseDuration("-72h")
	fmt.Println("https://www.vanguard.com.hk/portal/mvc/getETFNAVHistory.json?portId=9583&startDate=" + now.Add(d).Format("02-01-2006") + "&endDate=" + now.Format("02-01-2006"))
	res, err := http.DefaultClient.Get("https://www.vanguard.com.hk/portal/mvc/getETFNAVHistory.json?portId=9583&startDate=" + now.Add(d).Format("02-01-2006") + "&endDate=" + now.Format("02-01-2006"))
	if err != nil {
		return models.HkETFData{}, errors.New("Connection error")
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.HkETFData{}, errors.New("Read response body failed")
	}

	result := make(map[string]interface{})
	_ = json.Unmarshal(body, &result)
	data := result["histories"].([]interface{})[3].(map[string]interface{})["fundPrices"].([]interface{})[0].(map[string]interface{})
	return models.HkETFData{
		Name:  "领航标普五百",
		Price: float32(data["price"].(float64)),
		Date:  data["asOfDate"].(string),
	}, nil
}

// GetForex : 获取Forex的数据
func GetForex(symbols []string) ([]models.ForexData, error) {
	// 把输入转为小写
	for i, v := range symbols {
		symbols[i] = strings.ToLower(v)
	}
	fmt.Println(strings.Join(symbols, ",fx_s"))
	res, err := http.DefaultClient.Get("https://hq.sinajs.cn/?_=" + strconv.FormatInt(time.Now().Unix(), 10) + "&list=fx_s" + strings.Join(symbols, ",fx_s"))
	if err != nil {
		return []models.ForexData{}, errors.New("Connection error")
	}
	defer res.Body.Close()
	// 读取body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []models.ForexData{}, errors.New("Read response body failed")
	}
	// 数据转化为UTF-8
	s, err := myutils.GbkToUtf8(body)
	if err != nil {
		return []models.ForexData{}, errors.New("Transform to UTF-8 failed")
	}

	resStr := string(s)

	arr := strings.Split(resStr, "\n")
	var result []models.ForexData
	for i := 0; i < len(arr)-1; i++ {
		dataSlide := strings.Split(arr[i], ",")
		result = append(result, models.ForexData{})
		result[i].Time = strings.Split(dataSlide[0], "=\"")[1]
		result[i].Date = strings.Split(dataSlide[17], "\"")[0]
		last, _ := strconv.ParseFloat(dataSlide[1], 32)
		result[i].Last = float32(last)
		result[i].Name = dataSlide[9]
	}

	return result, nil
}
