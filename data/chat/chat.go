package chat

import (
	"github.com/google/uuid"
)

// Room A chatroom.
type Room struct {
	ID       uuid.UUID
	Name     string
	Messages []string
}

// Group A chat group.
type Group struct {
	IDs      []uuid.UUID
	Names    []string
	Messages []string
}

// NewRoom Creates new chatroom.
func NewRoom(name string) Room {
	return Room{
		ID:       uuid.New(),
		Name:     name,
		Messages: make([]string, 0),
	}
}

// NewGroup Creates a new chat group.
func NewGroup(rooms []Room) *Group {
	newGroup := Group{}
	for i := 0; i < len(rooms); i++ {
		newGroup.IDs = append(newGroup.IDs, rooms[i].ID)
		newGroup.Names = append(newGroup.Names, rooms[i].Name)
		newGroup.Messages = append(newGroup.Messages, rooms[i].Messages...)
	}
	return &newGroup
}
