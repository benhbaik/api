package memory

import (
	"api/data/chat"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/tidwall/buntdb"
)

var database *buntdb.DB

// OpenDB Opens or creates a file that does not persist to disk.
func OpenDB() {
	fmt.Println("-----Opening DB-----")
	if database == nil {
		fmt.Println("DB is nil opening new DB")
		db, e := buntdb.Open(":memory:")
		if e != nil {
			fmt.Println("Error opening DB")
			log.Fatal(e)
		}
		database = db
		fmt.Println("DB open...")
	}
}

// Insert Either inserts or replaces a chatroom in database based on key.
func Insert(room chat.Room) error {
	e := database.Update(func(tx *buntdb.Tx) error {
		byteValue, e := json.Marshal(room)
		if e != nil {
			return e
		}
		strValue := string(byteValue)
		tx.Set(room.Name, strValue, nil)
		return nil
	})
	return e
}

// Fetch Gets chatroom struct by string.
func Fetch(roomName string) (fetched chat.Room, err error) {
	var marshalledRoom string
	var room chat.Room

	e := database.View(func(tx *buntdb.Tx) error {
		val, e := tx.Get(roomName)
		if e != nil {
			return e
		}
		marshalledRoom = val
		return nil
	})
	if e != nil {
		fmt.Println("Could not find key.")
		return room, e
	}

	e = json.Unmarshal([]byte(marshalledRoom), &room)
	if e != nil {
		fmt.Println("Unmarshall failed, check format.")
		return room, e
	}
	return room, e
}

// FetchList Gets a list of chatroom structs by a list of strings.
func FetchList(roomNames []string) (fetched []chat.Room, err error) {
	var rooms []chat.Room
	marshalledRooms := "["
	errorList := make([]string, 0)

	e := database.View(func(tx *buntdb.Tx) error {
		for i := 0; i < len(roomNames); i++ {
			val, e := tx.Get(roomNames[i])
			if e != nil {
				errorList = append(errorList, roomNames[i])
				return e
			}
			marshalledRooms += val + ", "
		}
		marshalledRooms = strings.TrimRight(marshalledRooms, ", ")
		marshalledRooms = marshalledRooms + "]"
		return nil
	})
	if e != nil {
		fmt.Println("Could not find the following keys...")
		for i := 0; i < len(errorList); i++ {
			fmt.Println(errorList[i])
		}
		return rooms, e
	}

	e = json.Unmarshal([]byte(marshalledRooms), &rooms)
	if e != nil {
		fmt.Println("Unmarshall failed, check format.")
		return rooms, e
	}
	return rooms, e
}
