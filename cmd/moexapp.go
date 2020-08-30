package main

import (
	"fmt"
	"log"
	"moexapplication/cmd/requests"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

	//fmt.Println(newend.MakeUrl())
	//fmt.Println(request.GetUrlWithType(requests.HTML))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	resp, _ := doReq()
	_, err := fmt.Fprint(w, resp.Proto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func doReq() (*http.Response, error) {
	request := requests.New()
	request.NewSitenews()
	//request.AddEndpoint("URL")
	//request.NewSecurities()
	//request.AddQueryParam("MOEX")

	return request.ExexuteWithType(requests.HTML)
}
