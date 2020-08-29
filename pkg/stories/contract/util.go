package contract

type Story struct {
	ID    string `json:"id,omitempty"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
