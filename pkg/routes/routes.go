package routes

import (
	"github.com/gorilla/mux"
	"Blogging_Platform_Api/pkg/controller"
)

func RegisterBlogPostRoutes(mux *mux.Router) {

	mux.HandleFunc("/posts",controller.CreatePost).Methods("POST")
	mux.HandleFunc("/posts/{post_id}",controller.EditPost).Methods("PUT")
	mux.HandleFunc("/posts/{post_id}",controller.GetPost).Methods("GET")
	mux.HandleFunc("/posts/{post_id}",controller.DeletePost).Methods("DELETE")
	mux.HandleFunc("/posts",controller.GetAllPosts).Methods("GET")

}
