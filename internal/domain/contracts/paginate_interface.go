package contracts

type PaginateInterface interface {
	GetItems() []any
	GetTotal() int64
	GetCurrentPage() int64
	GetTotalPage() int64
	GetFirstPage() int64
	GetLastPage() int64
}
