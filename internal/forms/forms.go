package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	errors errors
	form   url.Values
}

func NewForm(data url.Values) *Form {
	return &Form{
		errors: make(map[string][]string),
		form:   data,
	}
}

func (form *Form) HasValue(field string, r *http.Request) bool {
	value := r.Form.Get(field)

	if value == "" {
		return false
	}

	return true
}
