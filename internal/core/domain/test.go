package domain

type Test struct {
	Id     string `bson:"_id"`
	Name   string `bson:"name"`
	Action string `bson: "action"`
}

func NewTest(Id string, Name string, Action string) Test {
	return Test{
		Id:     Id,
		Name:   Name,
		Action: Action}
}
