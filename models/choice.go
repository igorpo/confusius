package models

// Choices is a list of choice
type Choices []Choice

// Choice embodies the weights and options of one binary choice, that is, two options are weighted against each other
type Choice struct {
	Option1       string `json:"option-1"`
	Option2       string `json:"option-2"`
	Option1Weight int    `json:"option-1-weight"`
	Option2Weight int    `json:"option-2-weight"`
}
