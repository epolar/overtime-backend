package request

type AddUser struct {
	Name string `json:"name" validate:"required,max=8"`
}
