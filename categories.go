package main

type Category struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

var Categories = []Category{
	{Id: "AudioVideo", Name: "Audio & Video"},
	{Id: "Development", Name: "Developer Tools"},
	{Id: "Education", Name: "Education"},
	{Id: "Game", Name: "Games"},
	{Id: "Graphics", Name: "Graphics & Photography"},
	{Id: "Network", Name: "Communication & News"},
	{Id: "Office", Name: "Productivity"},
	{Id: "Science", Name: "Science"},
	{Id: "Settings", Name: "Settings"},
	{Id: "Utility", Name: "Utilities"},
}
