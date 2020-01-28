package main

import (
	"fmt"
	"log"
	"os"

	"cuvva/internal/web/crawler"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("please enter the website you wish to crawl immediately after the command e.g. crawler https://cuvva.com")
	}

	pages, err := crawler.Handle(args[0])
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range pages {
		fmt.Println(p.String())
	}
}
