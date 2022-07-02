package asanago

type Section struct {
	Gid          string
	ResourceType string `json:"resource_type"`
	Name         string
	CreatedAt    string `json:"created_at"`
	Project      Project
}
