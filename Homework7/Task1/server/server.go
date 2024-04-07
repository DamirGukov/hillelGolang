package server

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
)

func Start() error {
	http.HandleFunc("/", Handle)

	logrus.Info("prepared to listen :1888 port")
	err := http.ListenAndServe(":1888", nil)
	if err != nil {
		logrus.Error("failed to start listening :1888 port", err)
		return err
	}

	return nil
}

func Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getHandle(w, r)
		return
	}

	if r.Method == http.MethodPost {
		postHandle(w, r)
		return
	}

	if r.Method == http.MethodPut {
		putHandle(w, r)
		return
	}

	if r.Method == http.MethodDelete {
		deleteHandle(w, r)
		return
	}

	fmt.Println("method is not get ot post")
}

func getHandle(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("param1")
	param2 := r.URL.Query().Get("param2")
	fmt.Printf("GET request. Param1: %s, Param2: %s\n", param1, param2)
	fmt.Fprint(w, "GET request received")
}

type RequestBody struct {
	BodyParam1 string `json:"bodyParam1"`
	BodyParam2 string `json:"bodyParam2"`
}

func postHandle(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("POST request. BodyParam1: %s, BodyParam2: %s\n", requestBody.BodyParam1, requestBody.BodyParam2)
	fmt.Fprint(w, "POST request received")
}

func putHandle(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("PUT request. BodyParam1: %s, BodyParam2: %s\n", requestBody.BodyParam1, requestBody.BodyParam2)
	fmt.Fprint(w, "PUT request received")
}

func deleteHandle(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	fmt.Printf("DELETE request. BodyParam1: %s, BodyParam2: %s\n", requestBody.BodyParam1, requestBody.BodyParam2)
	fmt.Fprint(w, "DELETE request received")
}
