package asanago

type Tag struct {
	Gid          string
	ResourceType string `json:"resource_type"`
	Color        string
	CreatedAt    string
	Followers    []User
	Name         string
	Notes        string
	PermalinkUrl string `json:"permalink_url"`
	Workspace    Workspace
}
