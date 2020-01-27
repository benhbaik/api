package main

import (
	"api/controllers/chatcontroller"
	"api/controllers/usercontroller"
	"api/memory"
	"api/util"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Opening DB.")
	memory.OpenDB()

	a := &app{}
	fmt.Println("Starting app.")
	fmt.Println("Listening on port 8000...")
	http.ListenAndServe(":8000", a)
}

type app struct {
}

func (chatApp *app) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var head string
	head, req.URL.Path = util.ShiftPath(req.URL.Path)

	switch head {
	case "chat":
		chatcontroller.HandleHTTP(res, req)
	case "user":
		usercontroller.HandleHTTP(res, req)
	default:
		http.Error(res, "Not found.", http.StatusNotFound)
	}
}
