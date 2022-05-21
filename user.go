package asana

type User struct {
	Gid string
	Name string
	ResourceType string `json:"resource_type"`
	Workspaces []Workspace
}

type UserResponse struct {
	Data User
}
type UsersResponse struct {
	Data []User
}
