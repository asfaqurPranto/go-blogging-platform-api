package controller

import (
	"Blogging_Platform_Api/pkg/models"
	"encoding/json"

	//"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) { //Done
	all_post := models.GetAllPost()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(all_post)
	w.WriteHeader(200)
}

func CreatePost(w http.ResponseWriter, r *http.Request) { //done
	w.Header().Set("Content-Type", "application/json")

	var NewPost models.Post
	decoder := json.NewDecoder(r.Body)

	//Handling bad request error [Can ignore this]
	err := decoder.Decode(&NewPost)
	//fmt.Print(NewPost)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) //400 bad request
		return
	}
	if NewPost.Title == "" {
		http.Error(w, "title Can not be empty", 400)
		return
	}
	if NewPost.Content == "" {
		http.Error(w, "Content Can not be empty", 400)
		return
	}
	if NewPost.Tags == "" {
		http.Error(w, "Tags Can not be empty", 400)
		return
	}
	if NewPost.Catagory == "" {
		http.Error(w, "Catagory Can not be empty", 400)
		return
	}

	//

	models.CreatePost(&NewPost) //it will automaticly get id,date and other info
	encoder := json.NewEncoder(w)
	encoder.Encode(NewPost)

	w.WriteHeader(http.StatusCreated)

	//
}
func GetPost(w http.ResponseWriter, r *http.Request) { //Done
	vars := mux.Vars(r)
	id := vars["post_id"]
	Id, _ := strconv.ParseInt(id, 0., 0)
	//fmt.Print(Id)
	post, err := models.GetPostById(Id)
	if err != nil {
		http.Error(w, "Post not found with given id", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(post)

}
func DeletePost(w http.ResponseWriter, r *http.Request) { //Done
	vars := mux.Vars(r)
	id := vars["post_id"]
	Id, _ := strconv.ParseInt(id, 0., 0)
	_, err := models.DeletePostById(Id)

	// if post not found return 404 not found and
	if err != nil {
		http.Error(w, "Post not found with given id", 404)
		return
	}
	w.WriteHeader(204)

}
func EditPost(w http.ResponseWriter, r *http.Request) { //done
	w.Header().Set("Content-Type", "application/json")

	var Edited_post models.Post
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&Edited_post)
	// should return validation return(400) if some field is empty

	vars := mux.Vars(r)
	id := vars["post_id"]
	Id, _ := strconv.ParseInt(id, 0., 0)

	post, err := models.GetPostById(Id)
	if err != nil {
		http.Error(w, "Post not found with given id", 404)
		return
	}
	post.Title = Edited_post.Title
	post.Content = Edited_post.Content
	post.Tags = Edited_post.Tags
	post.Catagory = Edited_post.Catagory

	models.SaveExisting(&post)
	w.WriteHeader(200)
	encoder := json.NewEncoder(w)
	encoder.Encode(post)

}
