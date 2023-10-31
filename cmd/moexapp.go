package main

import (
	"fmt"
	"io"
	"log"
	"moexapplication/internal/requests"
	"net/http"
	"os"
)

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/s", securetiesHandler)
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

	resp, _ := doSitenewsReq(r.URL.RawQuery)
	var b []byte
	if resp.StatusCode != 200 {
		w.WriteHeader(resp.StatusCode)
	}
	var err error
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	bString := string(b)

	log.Printf(bString[:100])
	_, err = fmt.Fprintf(w, bString)
	if err != nil {
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func securetiesHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/s" {
		http.NotFound(w, r)
		return
	}

	resp, _ := doSecurityReq(r.URL.RawQuery)
	var b []byte
	if resp.StatusCode != 200 {
		w.WriteHeader(resp.StatusCode)
	}
	var err error
	defer resp.Body.Close()
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	bString := string(b)

	log.Printf(bString[:100])
	_, err = fmt.Fprintf(w, bString)
	if err != nil {
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func doSitenewsReq(q string) (*http.Response, error) {
	request := requests.New()
	request.NewSitenews()

	log.Printf(q)
	log.Printf(request.GetUrl())

	return request.ExexuteWithType(getResultType(q))
}

func doSecurityReq(q string) (*http.Response, error) {
	request := requests.New()
	request.NewSecurity("IMOEX")

	log.Printf(q)
	log.Printf(request.GetUrl())

	return request.ExexuteWithType(getResultType(q))
}

func getResultType(q string) requests.ResultType {
	var rt requests.ResultType
	switch q {
	case "j":
		rt = requests.JSON
	case "h":
		rt = requests.HTML
	case "c":
		rt = requests.CSV
	default:
		rt = requests.JSON
	}
	return rt
}
