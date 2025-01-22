package model

type Request struct {
	Command    string `json:"command"`              // required
	Subcommand string `json:"subcommand,omitempty"` // optional
	Tan        *int   `json:"tan,omitempty"`        // optional

	Priority *int   `json:"priority,omitempty"` // optional | required
	Origin   string `json:"origin,omitempty"`   // optional | required
	Duration *int   `json:"duration,omitempty"` // optional
}

type Response struct {
	Command  string      `json:"command"`
	Instance int         `json:"instance"`
	Success  bool        `json:"success"`
	Error    string      `json:"error"`
	Tan      int         `json:"tan"`
	Info     interface{} `json:"info"` // generic data
}
