package Model

type PublicBook struct {
	ID        int        `json:"id"`
	Year      int        `json:"Year"`
	Title     string     `json:"Title"`
	Handle    string     `json:"handle"`
	Publisher string     `json:"Publisher"`
	ISBN      string     `json:"ISBN"`
	Notes     []string   `json:"Notes"`
	CreatedAt string     `json:"created_at"`
	Villains  []Villains `json:"villains"`
	Pages     int        `json:"Pages"`
}

type Villains struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}
