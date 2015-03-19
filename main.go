package main

import (
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
)

// panicMiddleware intercepts all panics and tries to return them as HTTP errors
func panicMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				err, ok := p.(error)
				if !ok {
					err = errors.New(p.(string))
				}

				rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
				rw.WriteHeader(http.StatusInternalServerError)

				rw.Write([]byte(err.Error()))
			}
		}()

		handler(rw, req)
	}
}

// GET /fibonacci?n=3
func fibonacciHandler(rw http.ResponseWriter, req *http.Request) {
	N, err := strconv.ParseInt(req.URL.Query().Get("n"), 10, 64)
	if err != nil {
		panic(err)
	}

	if N < 0 {
		panic("n can't be negative")
	}

	i1, i2, x := big.NewInt(0), big.NewInt(0), big.NewInt(0)

	rw.Header().Set("Content-Type", "text/plain")

	for n := int64(0); n < N; n++ {
		if n == 1 {
			i2.SetInt64(1)
		} else if n > 1 {
			x.Set(i1)
			i1.Set(i2)
			i2.Add(x, i2)
		}

		result := fmt.Sprintf("%d ", i2)

		_, err = rw.Write([]byte(result))
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	http.Handle("/fibonacci", panicMiddleware(fibonacciHandler))
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
