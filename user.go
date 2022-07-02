package main

type User struct {
	ID       int
	Name     string `json:"name"`
	Password string `json:"pass"`
}
