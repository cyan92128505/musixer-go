package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"net/http"
	"os"

	"github.com/bobertlo/go-mpg123/mpg123"
	"github.com/gordonklaus/portaudio"
	"github.com/gorilla/mux"
)

func main() {
	buffer := microphone()
	router := mux.NewRouter()
	router.HandleFunc("/stream", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Content-Type", "audio/mpeg")

			audioChunks := bytes.NewBuffer(buffer)

			if _, err := audioChunks.WriteTo(w); err != nil {
				log.Println(err)
			}
		}
	})

	var ServerPort = os.Getenv("PORT")

	if ServerPort == "" {
		ServerPort = "8080"
	}

	log.Printf("The server is streaming on http://localhost:%s/stream", ServerPort)
	log.Fatal(http.ListenAndServe(":"+ServerPort, router), nil)
}

func microphone() []byte {
	decoder, err := mpg123.NewDecoder("")
	chk(err)

	// get audio format information
	rate, channels, _ := decoder.GetFormat()

	portaudio.Initialize()
	defer portaudio.Terminate()
	buffer := make([]int16, 8192)
	byteBuffer := make([]byte, 8192*2)
	stream, err := portaudio.OpenDefaultStream(0, channels, float64(rate), len(buffer), &buffer)

	for _, i := range buffer {
		binary.LittleEndian.PutUint16(byteBuffer, uint16(i))
	}

	chk(err)
	chk(stream.Start())
	defer stream.Close()
	return byteBuffer
}

func chk(err error) {
	if err != nil {
		panic(err)
	}
}
