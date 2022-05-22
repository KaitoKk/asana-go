package asanago

type Workspace struct {
	Gid string
	Name string
	ResourceType string `json:"resource_type"`
}

type WorkspaceResponse struct {
	Data Workspace
}

type WorkspacesResponse struct {
	Data []Workspace
}
