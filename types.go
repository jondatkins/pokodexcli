package main

type RootArray []Root

type Root struct {
	Name   string  `json:"name"`
	Email  string  `json:"email"`
	Gender string  `json:"gender"`
	Status string  `json:"status"`
	Id     float64 `json:"id"`
}
