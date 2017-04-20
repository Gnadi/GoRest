package entities

type Product struct {
	Id int
	Name string
	Quantity string
	Username string
}
type Message struct {
	Data Data `json:"data"`
	To string `json:"to"`
}
type Data struct {
	Body string `json:"body"`
	Title string `json:"title"`
}

