/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Ordenar objetos
 */

package src

import (
	"fmt"
	"sort"
)

type Manager struct {
	Name   string
	Age    int
	Active bool
}

func OrderAscStrings() {
	str := []string{"Eneas", "Amanda", "Maria", "José"}

	sort.Strings(str)

	for _, s := range str {
		fmt.Println(s)
	}
}

func OrderDescStrings() {
	str := []string{"Eneas", "Amanda", "Maria", "José"}

	sort.Sort(sort.Reverse(sort.StringSlice(str)))

	for _, s := range str {
		fmt.Println(s)
	}
}

func OrderAscObjectsManager() {
	managers := []*Manager{
		{Name: "Eneas", Active: true, Age: 30},
		{Name: "Marcos", Active: false, Age: 25},
		{Name: "Amanda", Active: true, Age: 20},
		{Name: "José", Active: false, Age: 40},
	}

	sort.Slice(managers, func(i, j int) bool {
		return managers[i].Name < managers[j].Name
	})

	for _, manager := range managers {
		fmt.Println(manager.Name)
	}
}

func OrderAscObjectManagerByNameAndAge() {
	managers := []*Manager{
		{Name: "Bruna", Active: true, Age: 30},
		{Name: "Amanda", Active: false, Age: 25},
		{Name: "Amanda", Active: true, Age: 20},
		{Name: "José", Active: false, Age: 40},
	}

	sort.Slice(managers, func(i, j int) bool {
		if managers[i].Name == managers[j].Name {
			return managers[i].Age < managers[j].Age
		}

		return managers[i].Name < managers[j].Name
	})

	for _, manager := range managers {
		fmt.Println(manager.Name, manager.Age)
	}
}
