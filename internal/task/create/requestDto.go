package create

type RequestDto struct {
	Title       string `json:"title"`
	Description *string `json:"description"`
}