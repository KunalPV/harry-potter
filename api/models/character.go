package models

type Character struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	AlternateNames  []string `json:"alternate_names"`
	House           string   `json:"house"`
	Ancestry        string   `json:"ancestry"`
	Patronus        string   `json:"patronus"`
	HairColour      string   `json:"hairColour"`
	HogwartsStudent bool     `json:"hogwartsStudent"`
	HogwartsStaff   bool     `json:"hogwartsStaff"`
	Actor           string   `json:"actor"`
	Wand            struct {
		Wood   string  `json:"wood"`
		Core   string  `json:"core"`
		Length float64 `json:"length"`
	} `json:"wand"`
}
