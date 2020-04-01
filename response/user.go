package response

type User struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Nick  string `json:"nick"`
}
