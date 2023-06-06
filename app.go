package main

//
import (
	"encoding/json"
	"gocrud/controller"
	routes "gocrud/routes"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
)

type App struct {
	Router  *mux.Router
	Session *mgo.Session
}

//Initializing The App

func (app *App) Initialise() error {
	var err error
	app.Session, err = mgo.Dial("mongodb+srv://rootuser1:rootuser1@cluster0.y1xnpso.mongodb.net/")
	if err != nil {
		println("Error while connecting to atlas")
	}
	app.Router = mux.NewRouter().StrictSlash(true)

	return nil
}

//Running the App

func (app *App) Run(address string) {
	defer app.Session.Close()
	log.Fatal(http.ListenAndServe(address, app.Router))
}

//Handeling Requests
func SendResponse(w http.ResponseWriter, statuscode int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		println(err)
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statuscode)
	w.Write(response)
}

//HAndelling Errors

func HandleError(w http.ResponseWriter, statuscode int, err string) {
	errstring := map[string]string{"error": err}
	SendResponse(w, statuscode, errstring)
}

//Handelling Routes

func (app *App) handleRoutes() {
	router := app.Router
	pc := controller.NewController(app.Session)
	routes.RouteHandeler(router, pc)
}
