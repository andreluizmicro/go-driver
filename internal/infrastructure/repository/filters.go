package repository

type Filters struct {
	Order string
	Limit int64
}

func NewFilters(order string, limit int64) *Filters {
	if order == "" && order != "ASC" && order != "DESC" {
		order = "ASC"
	}

	if limit <= 0 {
		limit = 1
	}

	return &Filters{
		Order: order,
		Limit: limit,
	}
}
