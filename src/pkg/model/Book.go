package Model

type Book struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Auther        string `json:"auther"`
	NumberOfPages int    `json:"total_pages"`
}
