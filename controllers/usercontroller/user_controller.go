package usercontroller

import (
	"encoding/json"
	"net/http"
)

// HandleHTTP Serve for /user
func HandleHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		handlePost(res)
	default:
		http.Error(res, "Check method type.", http.StatusMethodNotAllowed)
	}
}

func handlePost(res http.ResponseWriter) {
	js, _ := json.Marshal("User POST")
	res.Write(js)
}
