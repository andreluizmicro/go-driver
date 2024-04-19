package filter

var allowedFilters = []string{
	"name",
	"email",
}

type Filters struct {
	Page    int64
	PerPage int64
	Fields  []string
	Order   string
	Email   string
}

func NewFilters(page, perPage int64, fields []string, order, email string) *Filters {
	if page == 0 {
		page = 1
	}
	if perPage == 0 {
		perPage = 10
	}
	if order == "" {
		order = "id"
	}

	return &Filters{
		Page:    page,
		PerPage: perPage,
		Fields:  fields,
		Order:   order,
		Email:   email,
	}
}
