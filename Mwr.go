package mini_website_render

import (
	"fmt"
	"log"
	"net/http"
)

type Mwr struct {
	host string
	port int
}

func helloMwr(writer http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(writer, "MWR is NB!")
}

func (mwr *Mwr) Run() {
	http.HandleFunc("/", helloMwr)
	err := http.ListenAndServe(":9000", nil)
	log.Print("Server started at: http://127.0.0.1:9000")
	fmt.Print("Server started at: http://127.0.0.1:9000")
	if err != nil {
		log.Fatal("Error", err)
	}
}
