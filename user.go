package asanago

type User struct {
	Gid string
	Name string
	Email string
	ResourceType string `json:"resource_type"`
	Photo []string
	Workspaces []Workspace
}

type UserResponse struct {
	Data User
}
type UsersResponse struct {
	Data []User
}
