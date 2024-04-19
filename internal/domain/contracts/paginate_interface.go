package contracts

type PaginateInterface interface {
	Items() []any
	Total() int64
	Page() int64
	FirstPage() int64
	LastPage() int64
	CurrentPage() int64
	To() int64
	From() int64
}
