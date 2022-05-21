package asana

import (
	"io"
	"fmt"
	"errors"
	"encoding/json"
	"net/http"
)

type Client struct {
	baseUrl string
	apiKey string
	HttpClient *http.Client
}

func BuildClient(apiKey string) (*Client, error) {
	if len(apiKey) == 0 {
		return nil, errors.New("apiKeyがありません")
	}

	c := Client{
		"https://app.asana.com/api/1.0/",
		apiKey,
		&http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}

	return &c, nil
}

// User API

func (c Client) GetUser(userGid string) {//(User, error) {
	userPath := "users/" + userGid

	req, _ := http.NewRequest("GET", c.baseUrl + userPath, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + c.apiKey)

	res, _ := c.HttpClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

	var user UserResponse
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(user)
}


// Users API
func (c Client) GetUsers() ([]User, error) {
	usersPath := "users"

	req, _ := http.NewRequest("GET", c.baseUrl + usersPath, nil)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + c.apiKey)

	res, _ := c.HttpClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var users UsersResponse
	err := json.Unmarshal(body, &users)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return users.Data, nil
}
