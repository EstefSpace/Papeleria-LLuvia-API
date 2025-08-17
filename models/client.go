package models

type ClientAPI struct {
	ApiKey    string
	WebClient string
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
