package repository

type PaginateInterface interface {
	Total() int64
	CurrentPage() int64
	FirstPage() int64
	LastPage() int64
}
