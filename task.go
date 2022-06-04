package asanago

import (
	"encoding/json"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type Task struct {
	Gid             string
	ResourceType    string `json:"resource_type"`
	AssigneeStatus  string `json:"assignee_status"`
	Completed       bool
	CompletedAt     string `json:"completed_at"`
	CompletedBy     User   `json:"completed_by"`
	CreatedAt       string `json:"created_at"`
	Dependencies    []Task
	Dependents      []Task
	DueAt           string `json:"due_at"`
	DueIb           string `json:"due_on"`
	External        External
	HtmlNotes       string `json:"html_notes"`
	Liked           bool
	Likes           []Like
	ModifiedAt      string `json:"modified_at"`
	Name            string
	Notes           string
	NumLikes        int    `json:"num_likes"`
	NumSubtasks     int    `json:"num_tasks"`
	ResourceSubtype string `json:"resource_subtype"`
	StartAt         string `json:"start_at"`
	StartOn         string `json:"start_on"`
	Assignee        User
	AssigneeSection Section       `json:"assignee_section"`
	CustomFields    []CustomField `json:"custom_fields"`
	Followers       []User
	Parent          TaskCompact
	PermalinkUrl    string `json:"permalink_url"`
	Projects        []Project
	Tags            []Tag
	Workspace       Workspace
}

type TaskCompact struct {
	Gid          string
	ResourceType string `json:"resource_type"`
	Name         string
}

type TasksResponse struct {
	Data []TaskCompact
}

type GetTasksConfig struct {
	Assignee       string
	Project        string
	Section        string
	Workspace      string
	CompletedSince string
	ModifiedSince  string
}

func (op GetTasksConfig) BuildQueryParams() string {
	prefix := "?"
	rt := reflect.TypeOf(op)
	rv := reflect.ValueOf(op)

	var queries []string

	for i := 0; i < rt.NumField(); i++ {
		if rv.Field(i).String() == "" {
			continue
		}
		queries = append(queries, strings.ToLower(rt.Field(i).Name)+"="+rv.Field(i).String())
	}

	if len(queries) == 0 {
		return ""
	}
	return prefix + strings.Join(queries, "&")
}

func (c Client) GetTasks(config GetTasksConfig) ([]TaskCompact, error) {
	taskPath := "tasks" + config.BuildQueryParams()
	fmt.Println("path", taskPath)
	req, _ := c.buildRequest("GET", taskPath, nil)

	res, _ := c.HttpClient.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	fmt.Println(string(body))

	if res.StatusCode != 200 {
		fmt.Println("StatusCode: ", res.StatusCode)
		c.printError(body)
	}

	var tasks TasksResponse
	if err := json.Unmarshal(body, &tasks); err != nil {
		fmt.Println("error:", err)
		return nil, err
	}

	return tasks.Data, nil
}
