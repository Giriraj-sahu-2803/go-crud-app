package main

func main() {
	myApp := App{}
	myApp.Initialise()
	myApp.Run("localhost:5000")
	myApp.handleRoutes()
}
