package vo

type PageParam struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

type SheepFilter struct {
	Code     string `json:"code"`
	Category string `json:"category"`
	Age      int    `json:"age"`
}

type SheepQueryRequest struct {
	Sheep     SheepFilter `json:"sheep"`
	PageParam PageParam   `json:"page"`
}
