package delete

type Input struct {
	ID int64 `uri:"id" binding:"required"`
}
