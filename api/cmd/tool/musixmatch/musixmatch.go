/*
 * Musixmatch Go Client
 * @uses Swagger generated JavaScript client
 * @author Loreto Parisi (loreto at musixmatch dot com)
 * @see https://developer.musixmatch.com/documentation
 *  @copy 2016 Musixmatch Spa
 */
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"musixer/api/internal/app/initializers"
	"musixer/api/internal/pkg/musixmatch"
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}
func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Musixmatch Go Client")

	albumApi := musixmatch.NewAlbumApi()

	configuration := albumApi.Configuration
	configuration.SetDebug(true)
	configuration.APIKey["apikey"] = config.MusixmatchKey

	var albumId = "14250417"
	jsonResponse, response, _ := albumApi.AlbumGetGet(albumId, "json", "")

	// read http.response
	defer response.Body.Close()

	// unmarshall this struct to json
	b, err := json.Marshal(jsonResponse.Message)
	if err != nil {
		fmt.Println(err)
		return
	}
	var c = jsonPrettyPrint(string(b))
	fmt.Println(c)

}
