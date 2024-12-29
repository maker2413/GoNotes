package main

import (
	"log"

	"github.com/maker2413/GoNotes/BuildingGoModules/toolkit"
)

func main() {
	toSlug := "Now is the time 123!"

	var tools toolkit.Tools
	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Panicln(err)
	}

	log.Println(slugified)
}
