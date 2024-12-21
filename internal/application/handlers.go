package application

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

func QueryHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprint(w, request.Expression)

}
