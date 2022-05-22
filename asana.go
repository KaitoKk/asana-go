package asanago

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

func (c Client) buildRequest(method string, path string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, c.baseUrl + path, body)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Bearer " + c.apiKey)

	return req, nil
}

// User API
func (c Client) GetUser(userGid string) (User, error) {
	userPath := "users/" + userGid

	req, _ := c.buildRequest("GET", userPath, nil)

	res, _ := c.HttpClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var user UserResponse
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println("error:", err)
		var empty User
		return empty, err
	}

	return user.Data, nil
}


// Users API
func (c Client) GetUsers() ([]User, error) {
	usersPath := "users"

	req, _ := c.buildRequest("GET", usersPath, nil)

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



// Woekspace API
func (c Client) GetWorkspace(workspaceGid string) (Workspace, error) {
	workspacePath := "workspaces/" + workspaceGid

	req, _ := c.buildRequest("GET", workspacePath, nil)

	res, _ := c.HttpClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var workspace WorkspaceResponse
	err := json.Unmarshal(body, &workspace)
	if err != nil {
		fmt.Println("error:", err)
		var empty Workspace
		return empty, err
	}

	return workspace.Data, nil
}

// Woekspaces API
func (c Client) GetWorkspaces() ([]Workspace, error) {
	workspacesPath := "workspaces"

	req, _ := c.buildRequest("GET", workspacesPath, nil)

	res, _ := c.HttpClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))
	var workspaces WorkspacesResponse
	err := json.Unmarshal(body, &workspaces)
	if err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return workspaces.Data, nil
}
