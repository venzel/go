/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Scan terminal
 */

package src

import (
	"fmt"
	"os"
)

func ScanTerminal() {
	fmt.Println("Digite seu nome: ")
	var name string
	fmt.Scan(&name)
	fmt.Println("Seu nome é: ", name)
}

func ScanTerminalWithOs() {
	fmt.Println("Digite seu nome: ")
	var name string
	fmt.Fscan(os.Stdin, &name)
	fmt.Println("Seu nome é: ", name)
}
