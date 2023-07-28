/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo resolve o problema de chamar uma função com um map de funções
 */

package src

import (
	"fmt"
)

type fnx func(name string)

func a(name string) {
	fmt.Println("a ->", name)
}

func b(name string, size int) {
	fmt.Println("b ->", name, size)
}

func ExecuteFuncWithMap(fnName string) {
	var fa, fb fnx

	fa = func(name string) {
		a(name)
	}

	fb = func(name string) {
		b(name, 30)
	}

	mp := map[string]interface{}{
		"a": fa,
		"b": fb,
	}

	fn, exists := mp[fnName]

	if exists {
		fn.(fnx)("Tiago")
	} else {
		fmt.Println("função não encontrada")
	}
}
