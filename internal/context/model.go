package context

import "time"

type Resource struct {
	Value string `json:"value"`
	Notes string `json:"notes"`
}

type ProjectContext struct {
	Name      string     `json:"name"`
	Branch    string     `json:"branch"`
	Directory string     `json:"directory"`
	Links     []Resource `json:"links"`
	Terminals []Resource `json:"terminals"`
	CreatedAt time.Time  `json:"created_at"`
}
