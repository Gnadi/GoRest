package handler

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	"github.com/gnadlinger/Presentation/entities"
	"fmt"
	"net/http"
	"bytes"
	"encoding/json"
)

var db, err = gorm.Open("mysql", "Gnadlinger:admin@/presentation?charset=utf8&parseTime=True&loc=Local")
// region gin
func PostProduct(c *gin.Context)  {
	product:= entities.Product{
		Name:c.PostForm("name"),
		Quantity:c.PostForm("quantity"),
		Username:c.PostForm("username"),
	}
	db.Create(&product)
//region FCM
	url := "https://fcm.googleapis.com/fcm/send"

	fmt.Println("URL:>", url)
	foo:=entities.Message{
		Data:entities.Data{
			Body: c.PostForm("name"),
			Title:c.PostForm("username")+" Hat ein Produkt hinzugefügt",
		},
		To:"/topics/foo",
	}

	//var jsonStr = []byte(`{"data":{ "body":"i am the body","title":"I am the title"},"to":"/topics/foo"}`)
	str, err := json.Marshal(foo)
	//fmt.Println("JSON:>", jsonStr)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(str))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "key=AIzaSyDSrSrX4C4EGcu45HDja494foRekIvH--0")


	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//endregion
}
func GetAllProducts(c *gin.Context) {
	c.JSON(200,db.Find(&[]entities.Product{}))
}
//endregion
func GetAllProducts1(w http.ResponseWriter, r *http.Request) {

	j, _ := json.Marshal(db.Find(&[]entities.Product{}))
	w.Write(j)
}

func PostProduct1(w http.ResponseWriter, r *http.Request)  {
	var product = entities.Product{
		Name:r.FormValue("Name"),
		Quantity:r.FormValue("Quantity"),
		Username:r.FormValue("Username"),
	}
	db.Create(&product)

}
func DeleteProduct(w http.ResponseWriter, r *http.Request){
	db.Exec("Delete from  products where Id = "+r.FormValue("Id"))
}
func PutProduct(w http.ResponseWriter, r *http.Request){
	db.Exec("Update products set Username= '"+r.FormValue("Username")+"' where Id = "+r.FormValue("Id"))
}