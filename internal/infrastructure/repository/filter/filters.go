package filter

type Filters struct {
	Page    int64
	PerPage int64
	Filters map[string]interface{}
	Order   string
}

func NewFilters(page, perPage int64, filters map[string]interface{}, order string) *Filters {
	if order == "" {
		order = "ASC"
	}

	return &Filters{
		Page:    page,
		PerPage: perPage,
		Filters: filters,
		Order:   order,
	}
}
