package models

import(
	"github.com/jinzhu/gorm"
	"Blogging_Platform_Api/pkg/config"
	"errors"
)

var db *gorm.DB

type Post struct{
	gorm.Model
	Title string 	`json:"title"`
	Content string 	`json:"content"`
	Catagory string `json:"catagory"`
	Tags string 	`json:"tags"`

}

func init(){
	db=config.Connect_Mysql_server_and_ReturnDB()
	db.AutoMigrate(&Post{})
}

func CreatePost(post *Post){
	db.Create(post)
}
func GetAllPost() []Post{
	var posts []Post
	db.Find(&posts)
	return posts
}
func DeletePostById(id int64) (Post, error){
	var post Post
	info:=db.Where("ID=?",id).Delete(&post)
	if info.RowsAffected==0{
		err := errors.New("post not found with given id")
   	 	return post, err
	}
	return post,info.Error
}
func GetPostById(id int64) (Post , error){
	var post Post
	info:=db.Where("ID=?",id).Find(&post)
	if info.RowsAffected==0{
		//can not found in db
		err := errors.New("post not found with given id")
   	 	return post, err
	}
	return post,info.Error
}

func SaveExisting(post *Post){
	db.Save(post)
}