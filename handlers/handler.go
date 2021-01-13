package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sample-aws-app/api"
	abc "net/http"
	"os"
)

func ServeHttp() {

	router := mux.NewRouter()

	router.HandleFunc("/aws/listbucket", api.ListBucket).Methods("POST")
	router.HandleFunc("/aws/createBucket", api.CreateBucket).Methods("POST")
	router.HandleFunc("/aws/deleteBucket", api.DeleteBucket).Methods("POST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8001" //localhost
	}
	err := abc.ListenAndServe(":"+port, router) //Launch the app, visit localhost:8899/api
	if err != nil {
		fmt.Print(err)
	}
}
