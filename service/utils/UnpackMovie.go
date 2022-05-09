package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/oranges0da/entertainment-api/model"
)

func UnpackMovie(res *http.Response) model.Movie {
	body, err := ioutil.ReadAll(res.Body) // read res body

	if err != nil {
		panic(err)
	}

	var data model.Movie

	err = json.Unmarshal(body, &data) // unmarshal body to data

	if err != nil {
		panic(err)
	}

	return data
}