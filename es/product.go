package es

type Product struct {
	ID        string `json:"-" faker:"uuid_hyphenated"`
	Title     string `json:"title"`
	Published bool   `json:"published"`
}
