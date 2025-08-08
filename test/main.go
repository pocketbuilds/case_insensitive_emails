package main

import (
	"log"

	plugin "github.com/PocketBuilds/case_insensitive_emails"
	"github.com/pocketbase/pocketbase"
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
