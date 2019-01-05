package models

// InputDatastruct for go-ozzo validator
type InputData struct {
	ID        int32  `json:"id"`
	NamaAwal  string `json:"first_name"`
	NamaAkhir string `json:"last_name"`
	Email     string `json:"email"`
	Kondisi   bool   `json:"condition"`
}
