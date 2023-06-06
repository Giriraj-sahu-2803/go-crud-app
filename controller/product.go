package controller

import (
	"encoding/json"
	"fmt"
	"gocrud/model"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

type ProductController struct {
	session *mgo.Session
}

func NewController(session *mgo.Session) *ProductController {
	p := ProductController{session: session}
	return &p
}

func (p *ProductController) Homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to home page")
	fmt.Println("Endpoint hit")
}

//Get The list of Products
func (p *ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	collection := p.session.DB("test").C("Product")
	var products []model.Product
	err := collection.Find(nil).All(&products)
	if err != nil {
		println(" error in fetching data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(products)
}

//Get Product by ID
func (p *ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]

	if !bson.IsObjectIdHex(query) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	var product model.Product
	collection := p.session.DB("test").C("Product")
	err := collection.FindId(bson.ObjectIdHex(query)).One(&product)
	if err != nil {
		println(" error in fetching data")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)

}

//Create an POST a model product
func (p *ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	collection := p.session.DB("test").C("Product")
	err = collection.Insert(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

//DELETE a product
func (p *ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]

	if !bson.IsObjectIdHex(query) {
		w.WriteHeader(http.StatusInternalServerError)
	}
	collection := p.session.DB("test").C("Product")
	err := collection.RemoveId(bson.M{"_id": query})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusNoContent)

}

//Update a Product by ID
func (p *ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	query := mux.Vars(r)["id"]
	var updatedProduct model.Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if !bson.IsObjectIdHex(query) {
		w.WriteHeader(http.StatusInternalServerError)
	}

	collection := p.session.DB("test").C("Product")
	err = collection.Update(bson.ObjectIdHex(query), updatedProduct)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
