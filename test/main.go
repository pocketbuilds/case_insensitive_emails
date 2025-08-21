package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	plugin "github.com/pocketbuilds/case_insensitive_emails"
)

func main() {
	app := pocketbase.New()

	(&plugin.Plugin{
		// test config will go here
	}).Init(app)

	err := app.Start()
	if err != nil {
		log.Fatal(err)
	}
}
