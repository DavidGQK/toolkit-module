package main

import (
	"github.com/davidgqk/toolkit"
	"log"
)

func main() {
	toSlug := "NOW!!? is the time 123"

	var tools toolkit.Tools

	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		log.Println(err)
	}

	log.Println(slugified)
}
