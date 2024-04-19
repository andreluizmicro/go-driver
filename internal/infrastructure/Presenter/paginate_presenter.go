package Presenter

type PaginatePresenter struct {
	Data []any `json:"data"`
}

func NewPaginatePresenter(data []any) *PaginatePresenter {
	return &PaginatePresenter{
		Data: data,
	}
}

func (p *PaginatePresenter) Items() []any {
	return p.Data
}

func (p *PaginatePresenter) Total() int64 {
	return int64(len(p.Data))
}

func (p *PaginatePresenter) Page() int64 {
	return 0
}

func (p *PaginatePresenter) FirstPage() int64 {
	return 0
}

func (p *PaginatePresenter) LastPage() int64 {
	return 0
}

func (p *PaginatePresenter) CurrentPage() int64 {
	return 0
}

func (p *PaginatePresenter) To() int64 {
	return 0
}

func (p *PaginatePresenter) From() int64 {
	return 0
}
