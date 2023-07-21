/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Realiza a leitura e escrita de arquivos
 */

package src

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile() {
	file, err := os.Create("test.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	file.WriteString("Tiago")

	fmt.Println("Criou o arquivo")
}

func ReadFile() {
	file, err := os.ReadFile("test.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println("Conteudo do arquivo:", string(file))
}

func OpenFileWithBuffer() {
	file, err := os.Open("test.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	buffer := make([]byte, 1)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println("Conteudo:", string(buffer[:n]))
	}
}
