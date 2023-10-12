package validator_test

import (
	"testing"

	"github.com/tiryaki-repo/tiryaki_shared_models/services/validator"
)

func TestValidate(t *testing.T) {

	validator := validator.NewValidator()

	input := struct {
		A string `json:"a"`
		B string `json:"b"`
	}{A: "A", B: "B"}

	got := validator.Validate(input)
	var want error = nil

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
