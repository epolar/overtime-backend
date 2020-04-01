package request

type AddUser struct {
	Name  string `json:"name" validate:"required,max=8"`
	Label string `json:"label" validate:"required,max=16"`
	Nick  string `json:"nick" validate:"required,max=16"`
}
