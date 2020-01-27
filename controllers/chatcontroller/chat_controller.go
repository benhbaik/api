package chatcontroller

import (
	"api/data/chat"
	"api/memory"
	"api/util"
	"encoding/json"
	"net/http"
	"strings"
)

// HandleHTTP Serve for /chat
func HandleHTTP(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		handleGet(res, req)
	case "POST":
		handlePost(res, req)
	default:
		http.Error(res, "Check method type.", http.StatusMethodNotAllowed)
	}
}

func handleGet(res http.ResponseWriter, req *http.Request) {
	var rooms []chat.Room
	var chatGroup chat.Group
	var err error

	roomNamesStr := req.FormValue("room")
	roomNameList := strings.Split(roomNamesStr, ",")
	roomLen := len(roomNameList)

	if roomLen <= 0 {
		respondWithError(res, "Please specify a chatroom.")
	} else if roomLen == 1 {
		room, err := memory.Fetch(roomNameList[0])
		if err != nil {
			respondWithError(res, "Could not find specified chatroom.")
			return
		}
		rooms = append(rooms, room)
	} else {
		rooms, err = memory.FetchList(roomNameList)
		if err != nil {
			respondWithError(res, "Could not find specified chatroom.")
			return
		}
		chatGroup = chat.NewGroup(rooms)
	}

	if rooms != nil {
		js, _ := json.Marshal(chatGroup)
		res.Write(js)
	} else {
		js, _ := json.Marshal(rooms)
		res.Write(js)
	}
}

func handlePost(res http.ResponseWriter, req *http.Request) {
	var room chat.Room
	var err error
	roomName := req.FormValue("room")
	message := req.Form.Get("message")

	if roomName != "" && message != "" {
		room, err = memory.Fetch(roomName)
		if err != nil {
			respondWithError(res, "Failed to find chatroom.")
		}
		room.Messages = append(room.Messages, message)
		err = memory.Insert(room)
	} else if roomName == "" {
		room = chat.NewRoom(util.RandStringRunes(10))
		err = memory.Insert(room)
		if err != nil {
			respondWithError(res, "Failed to create chatroom.")
		}
	} else {
		room = chat.NewRoom(roomName)
		err = memory.Insert(room)
		if err != nil {
			respondWithError(res, "Failed to create chatroom.")
		}
	}

	js, _ := json.Marshal(room)
	res.Write(js)
}

func respondWithError(res http.ResponseWriter, errorMessage string) {
	js, _ := json.Marshal(errorMessage)
	res.Write(js)
}
