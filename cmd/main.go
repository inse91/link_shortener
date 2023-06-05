package main

import "link_shortener/app"

type user struct {
	Age   int
	Name  string
	Phone string
}

var phones = map[string]string{
	"dan":   "123",
	"mary":  "456",
	"john":  "789",
	"ross":  "1244",
	"trent": "1244",
}

func main() {
	app.Start()
}
