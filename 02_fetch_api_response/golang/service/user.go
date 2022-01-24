package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nahidhasan98/care-t/02_fetch_api_response/golang/model"
)

// GetUserList function for getting user list
func GetUserList() ([]model.User, error) {
	// setting api url
	apiURL := "https://gorest.co.in/public/v1/users"

	// sending api request
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// taking response body
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// parsing response vody to struct/variable
	var userResponse model.UserResponse
	err = json.Unmarshal(resp, &userResponse)
	if err != nil {
		return nil, err
	}

	// returning user data, ignoring meta part from response
	return userResponse.Data, nil
}

// CreateUser function for creating a new user
func CreateUser(user *model.User) (*model.User, error) {
	// setting api url
	apiURL := "https://gorest.co.in/public/v1/users"

	//setting api request method
	method := "POST"

	// preparing payload/post data
	ploadData, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	payload := bytes.NewReader(ploadData)

	// setting up new request
	req, err := http.NewRequest(method, apiURL, payload)
	if err != nil {
		return nil, err
	}

	// setting up request header
	// token should not be placed directly. it should be stored at some safe place/directory
	// but for now it is given directly
	req.Header.Add("Authorization", "Bearer d7c01847de4c083cb154e9a533294301e9f05f93dbae7d589e42ece63226c0a3")
	req.Header.Add("Content-Type", "application/json")

	// executing api request with new http client
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// taking response body
	resp, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// parsing response vody to struct/variable
	var apiResponse model.DynamicResponse
	err = json.Unmarshal(resp, &apiResponse)
	if err != nil {
		return nil, err
	}

	// if user created successfully, response are backed into model.User format
	// but if any error occurred, then response are backed into an array format(i.e. []model.User)
	// so struct field (Data) taken ase interfaces, so that it can be checked later and parsed accordingly
	// to do so, at first response type is checked, if it become a type of arry([]interface{}), that means some error occurred
	// and then error message is returned
	respType := fmt.Sprintf("%T", apiResponse.Data)

	if respType == "[]interface {}" { // error occurred
		// retrieving error messages
		errData := apiResponse.Data.([]interface{}) // interface type assertion
		errMsg := ""
		for _, val := range errData { // loop over all error and taking to a string
			msg := val.(map[string]interface{}) // typer assertion
			errMsg += fmt.Sprintf("%v %v\n", msg["field"], msg["message"])
		}
		// returning error to calling function
		return nil, errors.New(errMsg)
	}

	// at this point, response type is not an array, means no error occured.
	// user successfully created.
	// now for getting user id, response data is being parsed
	id := apiResponse.Data.(map[string]interface{})["id"].(float64) // interface type assertion
	user.ID = int(id)                                               // only taking inserted it, cause already have other information

	return user, nil
}
