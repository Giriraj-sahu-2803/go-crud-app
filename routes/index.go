package routes

import (
	controller "gocrud/controller"

	"github.com/gorilla/mux"
)

func RouteHandeler(router *mux.Router, pc *controller.ProductController) {
	router.HandleFunc("/product/{id}", pc.DeleteProduct).Methods("DELETE")
	router.HandleFunc("/product/{id}", pc.UpdateProduct).Methods("PUT")
	router.HandleFunc("/product/{id}", pc.GetProduct).Methods("GET")
	router.HandleFunc("/products", pc.GetProducts).Methods("GET")
	router.HandleFunc("/product", pc.CreateProduct).Methods("POST")
	router.HandleFunc("/", pc.Homepage)
}
