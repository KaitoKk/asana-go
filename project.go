package asanago

type Project struct {
	Gid                 string
	ResourceType        string `json:"resource_type"`
	Archived            bool
	Color               string
	CreatedAt           string               `json:"created_at"`
	CurrentStatus       StatusUpdate         `json:"current_status"`
	CustomFieldSettings []CustomFieldSetting `json:"custom_field_settings"`
	DefaultView         string               `json:"default_view"`
	DueDate             string               `json:"due_date"`
	DueOn               string               `json:"due_on"`
	HtmlNotes           string               `json:"html_notes"`
	IsTemplate          bool                 `json:"is_template"`
	Members             []User
	ModifiedAt          string `json:"modified_at"`
	Name                string
	Notes               string
	Public              bool
	StartOn             string `json:"start_on"`
	Workspace           Workspace
	Completed           bool
	CompletedAt         string `json:"completed_at"`
	CompletedBy         User
	CreatedFromTemplate ProjectTemplate `json:"created_from_template"`
	CustomFields        []CustomField   `json:"custom_fields"`
	Followers           []User
	Icon                string
	Owner               User
	PermalinkUrl        string `json:"permalink_url"`
	Team                Team
}

type ProjectTemplate struct {
	Gid             string
	ResourceType    string `json:"resource_type"`
	Color           string
	Description     string
	HtmlDescription string `json:"html_description"`
	Name            string
	Owner           User
	Public          bool
	RequestedDates  string
	Team            Team
}

type ProjectCompact struct {
	Gid          string
	Name         string
	ResourceType string `json:"resource_type"`
}
