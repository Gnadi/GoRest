package main

import (
	//"github.com/gnadlinger/Presentation/routes"
	"github.com/jinzhu/gorm"
	"github.com/gnadlinger/Presentation/entities"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/gnadlinger/Presentation/handler"
)
var db, err = gorm.Open("mysql", "Gnadlinger:admin@/presentation?charset=utf8&parseTime=True&loc=Local")

func main(){
	prepareDb()
	//routes.CreateRoutes()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/foo", handler.PostProduct1).Methods("POST")
	router.HandleFunc("/get",handler.GetAllProducts1).Methods("GET")
	router.HandleFunc("/delete",handler.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/put",handler.PutProduct).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
func prepareDb(){
	defer db.Close()
	if err != nil {
		panic("failed to connect database")
	}
	db.DropTable(&entities.Product{})
	db.AutoMigrate(&entities.Product{})
	product:= entities.Product{
		Name:"Popcorn",
		Quantity:"10",
	}
	product1:= entities.Product{
		Name:"Schnitzel",
		Quantity:"10",
	}
	db.Create(&product)
	db.Create(&product1)
}