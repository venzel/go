/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Buscar de objetos
 */

package src

import (
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func findPersonInSlice(persons []*Person, name string) (*Person, bool) {
	for _, person := range persons {
		if person.Name == name {
			return person, true
		}
	}

	return nil, false
}

func GetPersonByName(name string) *Person {
	persons := []*Person{
		{Name: "Tiago", Age: 31},
		{Name: "Eneas", Age: 31},
		{Name: "Maria", Age: 31},
	}

	person, exists := findPersonInSlice(persons, name)

	if !exists {
		fmt.Println("Pessoa não encontrada!")
		return nil
	}

	fmt.Println(person)

	return person
}

func findPersonsByNameContain(persons []*Person, partName string) ([]*Person, bool) {
	var aux []*Person
	var exists bool

	for _, person := range persons {
		if strings.Contains(strings.ToLower(person.Name), strings.ToLower(partName)) {
			exists = true
			aux = append(aux, person)
		}
	}

	return aux, exists
}

func GetPersonsByNameContain(nameContain string) []*Person {
	p := []*Person{
		{Name: "Tiago", Age: 31},
		{Name: "Eneas", Age: 31},
		{Name: "Maria", Age: 31},
		{Name: "jiraia", Age: 31},
	}

	persons, exists := findPersonsByNameContain(p, nameContain)

	if !exists {
		fmt.Println("Pessoa não encontrada!")
		return nil
	}

	for _, person := range persons {
		fmt.Println(person.Name, person.Age)
	}

	return persons
}
