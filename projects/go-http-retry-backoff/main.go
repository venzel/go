package main

import (
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

func elapsed(startTime, endTime time.Time) {
	elapsedTime := endTime.Sub(startTime).Seconds()

	elapsedTimeRounded := math.Round(elapsedTime*100) / 100

	fmt.Println("Tempo decorrido:", elapsedTimeRounded, "s")
}

func call(url string) (*http.Response, error) {
	client := &http.Client{}

	res, err := client.Get(url)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func main() {
	url := "https://www.google.com.brx"

	maxAttempts := 3

	bf := backoff.NewExponentialBackOff()
	bf.InitialInterval = 500 * time.Microsecond
	bf.MaxInterval = 1 * time.Second
	bf.MaxElapsedTime = 1 * time.Second

	var (
		res *http.Response
		err error
	)

	startTime := time.Now()

	for i := 0; i < maxAttempts; i++ {
		cb := func() error {
			res, err = call(url)

			if err != nil {
				fmt.Printf("Erro na chamada HTTP: %s. Tentativa #%d\n", err.Error(), i+1)
				return err
			}

			fmt.Println("Chamada HTTP bem-sucedida!")

			return nil
		}

		err = backoff.Retry(cb, bf)

		if err == nil {
			break
		}
	}

	endTime := time.Now()

	elapsed(startTime, endTime)

	if err != nil {
		fmt.Println("Falha na chamada HTTP após várias tentativas.")
		return
	}

	defer res.Body.Close()
}
