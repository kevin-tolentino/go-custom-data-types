package organization

import "fmt"

type Identifiable interface {
	ID() string
}

type Person struct {
		firstName string
		lastName string
}

//create a new function called new plus the type
func NewPerson(firstName, lastName string) Person {
	return Person{
		firstName: firstName,
		lastName:  lastName,
	}
}

func (p Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)

}

func (p Person) ID() string { //implicit implementation of interface of type Identifiable
	return "12345"
}
