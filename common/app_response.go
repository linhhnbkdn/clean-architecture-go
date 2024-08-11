package common

type SuccessRes struct {
	Data    interface{} `json:"data"`
	Paging  interface{} `json:"paging,omitempty"`
	Filters interface{} `json:"filters,omitempty"`
}

func NewSuccessResponse(data interface{}, paging interface{}, filters interface{}) *SuccessRes {
	return &SuccessRes{
		Data:    data,
		Paging:  paging,
		Filters: filters,
	}
}

func SimpleSuccessResponse(data interface{}) *SuccessRes {
	return &SuccessRes{
		Data: data,
	}
}
