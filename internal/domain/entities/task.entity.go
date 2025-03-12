package entities

type Task struct {
	Uuid    string `json:"uuid"`
	User    User   `json:"user"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
}
