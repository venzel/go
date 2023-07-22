/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Slices de objetos
 */

package src

import "fmt"

type Member struct {
	Name   string
	Active bool
}

func SlicesObjects() {
	members := []*Member{
		{Name: "Eneas", Active: true},
		{Name: "João", Active: false},
		{Name: "Maria", Active: true},
		{Name: "José", Active: false},
	}

	member0and1 := members[:2]  // Enéas e João
	member2and3 := members[2:]  // Maria e José
	member1and3 := members[1:3] // João e Maria

	for _, member := range member0and1 {
		fmt.Println(member.Name)
	}

	for _, member := range member2and3 {
		fmt.Println(member.Name)
	}

	for _, member := range member1and3 {
		fmt.Println(member.Name)
	}
}
