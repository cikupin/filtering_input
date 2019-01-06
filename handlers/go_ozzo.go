package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/cikupin/filtering_input/helpers"
	"github.com/cikupin/filtering_input/models"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type inputData models.InputData

// Validate using Go-Ozzo validator library
func (d inputData) Validate() error {
	err := validation.ValidateStruct(&d,
		validation.Field(&d.ID, validation.Required),
		validation.Field(&d.NamaAwal, validation.Required),
		validation.Field(&d.NamaAkhir, validation.Required),
		validation.Field(&d.Email, validation.Required),
		validation.Field(&d.Email, validation.Required, is.Email),
	)
	return err
}

// GoOzzoValidator - POST method
func GoOzzoValidator(w http.ResponseWriter, r *http.Request) {
	var data inputData

	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	err := data.Validate()
	if err != nil {
		errValidation, _ := json.Marshal(err)
		helpers.ResponseFromStruct(w, errValidation)
		return
	}

	// modifying data
	data.ID = 666
	data.NamaAwal = strings.Join([]string{data.NamaAwal, " -- modif awal"}, "")
	data.NamaAkhir = strings.Join([]string{data.NamaAkhir, " -- modif akhir"}, "")
	data.Email = "bodoamat@123.com"
	data.Kondisi = false

	newData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	helpers.ResponseFromStruct(w, newData)
}
