package folder

type CreateInput struct {
	Name     string `json:"name" binding:"required"`
	ParentId int64  `json:"parent_id"`
}

type CreateOutput struct {
	Id int64 `json:"id"`
}

type UpdateInput struct {
	ID   int64  `uri:"id"`
	Name string `json:"name" binding:"required,min=3"`
}

type UpdateOutput struct {
	Success bool
}
