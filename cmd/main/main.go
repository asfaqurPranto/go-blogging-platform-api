package main

import (
	"Blogging_Platform_Api/pkg/routes"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)
func main(){
	mux:=mux.NewRouter() //create a object of mux.Router
	routes.RegisterBlogPostRoutes(mux)

	err:=http.ListenAndServe(":8080",mux)
	if err!=nil{
		fmt.Print(err.Error())
	}

}