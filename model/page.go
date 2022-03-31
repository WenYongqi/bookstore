package model

type Page struct {
	Books *[]Book
	PageNo int64
	PageSize int64
	TotalPageNo int64
	TotalRecord int64
	MinPrice string
	MaxPrice string //根据min和max是否均为空字符串来判断是否为价格区间查询
	IsLogin bool
	Username string
}

func (p *Page) HasPrev() bool {
	return p.PageNo != 1
}

func (p *Page) HasNext() bool {
	return p.PageNo != p.TotalPageNo
}

func (p *Page) GetPrev() int64 {
	return p.PageNo - 1
}

func (p *Page) GetNext() int64 {
	return p.PageNo + 1
}