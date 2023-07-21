/**
 * @author Enéas Almeida <eneas.eng@yahoo.com>
 * @description Http com contexto
 */

package src

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetRequestWithContext() {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*5)

	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://google.com", nil)

	if err != nil {
		panic(err)
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}
