package application

import "net/http"

func Run() {
	http.HandleFunc("/api/v1/calculate", QueryHandler)
	http.ListenAndServe(":8080", nil)
}
