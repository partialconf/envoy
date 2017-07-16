package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	apiKey = flag.String("key", os.Getenv("API_KEY"), "the Meetup API ket")

	text     = flag.String("text", "", "filter by text")
	location = flag.String("location", "", "filter by location")
)

func main() {
	flag.Parse()

	if *apiKey == "" {
		panic("Require an Meetup API key present in the API_KEY env")
	}

	client := NewClient(*apiKey)
	groups, err := client.FindGroups(*location, *text)
	if err != nil {
		panic(err)
	}

	fmt.Println(groups)
}
