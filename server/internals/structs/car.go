package structs

type Car struct {
	ID      int    `json:"id"`
	Type    string `json:"type"`
	Color   string `json:"color"`
	Make    string `json:"make"`
	Model   string `json:"model"`
	Image   []byte `json:"-"`
	Caption string `json:"caption"`
}
