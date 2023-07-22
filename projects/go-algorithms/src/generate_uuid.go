/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description O algoritmo apenas utiliza a biblioteca do Google de gerar UUID
 */

package src

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUUID() {
	id := uuid.New().String()

	fmt.Println(id)
}
