package pkg

type Reminder struct {
	Id       string   `json:"id,omitempty"`
	Title    string   `json:"title,omitempty"`
	Status   string   `json:"status,omitempty"`
	Notes    string   `json:"notes,omitempty"`
	Category string   `json:"category,omitempty"`
	Priority string   `json:"priority,omitempty"`
	Flag     bool     `json:"flag,omitempty"`
	Tags     []string `json:"tags,omitempty"`
	DueOn    string   `json:"dueOn,omitempty"`
}
