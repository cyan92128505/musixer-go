package main

import (
	"context"
	"fmt"
	"log"
	"musixer/api/internal/app/initializers"
	"net/http"

	mxm "github.com/milindmadhukar/go-musixmatch"
	"github.com/milindmadhukar/go-musixmatch/params"
)

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(config.MusixmatchKey)

	client := mxm.New(config.MusixmatchKey, http.DefaultClient)

	artists, err := client.SearchArtist(context.Background(), params.QueryArtist("Martin Garrix"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println((*artists)[0])

}
