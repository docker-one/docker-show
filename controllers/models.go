package controllers

type Rsp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ImagesInfo struct {
	SensorId string  `json:"sensorId"`
	CH4      float64 `json:"CH4"`
	CO2      float64 `json:"CO2"`
	CO       float64 `json:"CO"`
	T        float64 `json:"T"`
	H        float64 `json:"H"`
	O2       float64 `json:"O2"`
}

type PageHelper struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type ResultDataList struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
