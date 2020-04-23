package api

import (
	"io/ioutil"
	"net/http"
)

var (
	client = &http.Client{}
)

func get(url string) (string, error) {
	res, err := client.Get(url)
	if err != nil {
		return "", err
	}

	defer func() {
		_ = res.Body.Close()
	}()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
