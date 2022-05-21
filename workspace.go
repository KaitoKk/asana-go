package asanago

type Workspace struct {
	Gid string
	Name string
}

type WorkspaceResponse struct {
	Data Workspace
}

type WorkspacesResponse struct {
	Data []Workspace
}
