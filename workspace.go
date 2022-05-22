package asanago

type Workspace struct {
	Gid string
	Name string
	ResourceType string `json:"resource_type"`
	EmailDomains []string `json:"email_domains"`
	IsOrganization bool `json:"is_organization"`
}

type WorkspaceResponse struct {
	Data Workspace
}

type WorkspacesResponse struct {
	Data []Workspace
}
