package utils

import (
	"bookstore/items-api/utils/errors"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func AuthenticateToken(r *http.Request) (int64, *errors.RestErr) {


	if _, ok := r.Header["Authorization"]; !ok {
		return -1, errors.NewBadRequestError("authentication failed")
	}

	token := r.Header["Authorization"][0]

	client := &http.Client{}
	// req, err := http.NewRequest("GET", "http://localhost:9092/token/login", nil)
	req, err := http.NewRequest("GET", "http://users:9092/token/login", nil)
	if err != nil {
		return -1, errors.NewBadRequestError(err.Error())
	}
	req.Header.Set("Authorization", token) 
	res, err := client.Do(req)
	if err != nil {
		return -1, errors.NewBadRequestError(err.Error())
	}
	body, _ := ioutil.ReadAll(res.Body)
	x := make(map[string]interface{})

	if err := json.Unmarshal(body, &x); err != nil {
		return -1, errors.NewBadRequestError(err.Error())
	}

	if _, ok := x["error"] ; ok {
		return -1, errors.NewBadRequestError(x["message"].(string))
	}

	return int64(x["userId"].(float64)), nil
}