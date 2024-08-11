package common

type Paging struct {
	Page     int `json:"page" form:"page"`
	PageSize int `json:"pageSize" form:"pageSize"`
	Total    int `json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page == 0 {
		p.Page = 1
	}
	if p.PageSize <= 1 {
		p.PageSize = 1
	}
	if p.PageSize >= 100 {
		p.PageSize = 100
	}
}
