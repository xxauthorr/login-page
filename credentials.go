package main

type Credentials struct {
	Heading string
	Email   string
	ErrMsg  string
}

// for checking the api control using postman
type Article struct {
	Email   string `json:"email"`
	Name    string `json:"name"`
	Details string `json:"details"`
}

type Articles []Article
