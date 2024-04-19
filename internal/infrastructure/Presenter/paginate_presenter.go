package Presenter

type PaginatePresenter struct {
	Data        []any
	Total       int64
	CurrentPage int64
	TotalPage   int64
	FirstPage   int64
	LastPage    int64
}

func (p *PaginatePresenter) GetItems() []any {
	return p.Data
}

func (p *PaginatePresenter) GetTotal() int64 {
	return p.Total
}

func (p *PaginatePresenter) GetCurrentPage() int64 {
	return p.CurrentPage
}

func (p *PaginatePresenter) GetTotalPage() int64 {
	return p.TotalPage
}

func (p *PaginatePresenter) GetFirstPage() int64 {
	return p.FirstPage
}

func (p *PaginatePresenter) GetLastPage() int64 {
	return p.LastPage
}
