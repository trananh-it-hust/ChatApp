package util

type Pagination struct {
	CurrPage int `json:"page"`
	PageSize int `json:"pageSize"`
	Total    int `json:"total"`
	Pages    int `json:"pages"`
}

func NewPagination(currPage, pageSize, total int) *Pagination {
	return &Pagination{
		CurrPage: currPage,
		PageSize: pageSize,
		Total:    total,
		Pages:    (total + pageSize - 1) / pageSize,
	}
}

func (p *Pagination) GetPage() int {
	if p.CurrPage <= 0 {
		return 1
	}
	return p.CurrPage
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize <= 0 {
		return 10
	}
	return p.PageSize
}
