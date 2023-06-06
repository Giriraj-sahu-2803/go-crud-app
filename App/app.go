package App

//
import (
	"encoding/json"
	"gocrud/controller"
	routes "gocrud/routes"
	"log"
	"net/http"
	"os"

	"gopkg.in/mgo.v2"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type App struct {
	Router  *mux.Router
	Session *mgo.Session
}

//Initializing The App

func (app *App) Initialise() error {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Session, err = mgo.Dial(os.Getenv("CONNECTION_STRING"))
	if err != nil {
		log.Fatalf("%v", err)
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

func (app *App) HandleRoutes() {
	router := app.Router
	pc := controller.NewController(app.Session)
	routes.RouteHandeler(router, pc)
}
