package domain

type Entity struct {
	Id     string `bson:"_id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Action string `bson:"action" json:"action"`
}

func NewEntity(id string, name string, action string) *Entity {
	return &Entity{Id: id, Name: name, Action: action}
}
