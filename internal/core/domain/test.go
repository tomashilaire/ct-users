package domain

type Test struct {
	Id     string `bson:"_id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Action string `bson:"action" json:"action"`
}

func NewTest(id string, name string, action string) *Test {
	return &Test{Id: id, Name: name, Action: action}
}
