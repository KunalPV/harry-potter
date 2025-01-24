package models

type Character struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	AlternateNames  []string `json:"alternate_names"`
	Species         string   `json:"species"`
	Gender          string   `json:"gender"`
	House           string   `json:"house"`
	DateOfBirth     string   `json:"dateOfBirth"`
	Wizard          bool     `json:"wizard"`
	Ancestry        string   `json:"ancestry"`
	EyeColor        string   `json:"eyeColour"`
	HairColour      string   `json:"hairColour"`
	Patronus        string   `json:"patronus"`
	HogwartsStudent bool     `json:"hogwartsStudent"`
	HogwartsStaff   bool     `json:"hogwartsStaff"`
	Actor           string   `json:"actor"`
	Alive           bool     `json:"alive"`
	Wand            struct {
		Wood   string  `json:"wood"`
		Core   string  `json:"core"`
		Length float64 `json:"length"`
	} `json:"wand"`
}
