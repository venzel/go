/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Split strings
 */

package src

import (
	"fmt"
	"strings"
)

func SplitFullName(name string) {
	parts := strings.Split(name, " ")

	first, last := parts[0], parts[1]

	fmt.Println(first, last)
}
