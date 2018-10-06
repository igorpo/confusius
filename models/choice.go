package models

// Choices is a list of choice
type Choices []Choice

// Choice embodies the weights and options of one binary choice, that is, two options are weighted against each other
type Choice struct {
	FirstOption  Option `json:"option_1"`
	SecondOption Option `json:"option_2"`
}

// Option encodes one option
type Option struct {
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}
