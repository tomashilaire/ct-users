package domain

type Test struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

func NewTest(id string, name string) *Test {
	return &Test{Id: id, Name: name}
}
