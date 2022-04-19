package domain

type Test struct {
	Id   string `bson:"_id"`
	Name string `bson:"name"`
}

func NewTest(Id string, Name string) Test {
	return Test{Id: Id, Name: Name}
}
