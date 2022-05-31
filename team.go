package asanago

type Team struct {
	Gid             string
	ResourceType    string `json:"resource_type"`
	Name            string
	Description     string
	HtmlDescription string `json:"html_description"`
	Organization    Workspace
	PermalinkUrl    string `json:"permalink_url"`
	Visibility      string
}
