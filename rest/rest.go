package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yoonhero/initial_go_restapi_server/utils"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload"`
}

func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "See Documentation",
		},
		{
			URL:         url("/example"),
			Method:      "GET",
			Description: "Rest Api Get Exapmle",
		},
		{
			URL:         url("/post"),
			Method:      "POST",
			Description: "Rest Api POST Example",
		},
	}
	json.NewEncoder(rw).Encode(data)
}

func example(rw http.ResponseWriter, r *http.Request) {
	data := "HELLO WORLD"

	json.NewEncoder(rw).Encode(data)
}

type addPayload struct {
	Message string `json:"message"`
}

func post(rw http.ResponseWriter, r *http.Request) {
	var payload addPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))

	fmt.Println(payload)

	rw.WriteHeader(http.StatusCreated)
}

func jsonContentTypeMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		utils.AllowConnection(rw)
		rw.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(rw, r)
	})
}

func loggerMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL)
		next.ServeHTTP(rw, r)
	})
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()

	router.Use(jsonContentTypeMiddleWare, loggerMiddleWare)

	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/example", example).Methods("GET")
	router.HandleFunc("/post", post).Methods("POST")
	fmt.Printf("Listening on http://localhost%s\n", port)

	log.Fatal(http.ListenAndServe(port, router))
}
