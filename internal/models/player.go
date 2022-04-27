package models

type Player struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Nation  string `json:"nation"`
	Team    string `json:"team"`

	Position string `json:"position"`
	Height   int    `json:"height"`
	Weight   int    `json:"weight"`
}

type Coach struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}
