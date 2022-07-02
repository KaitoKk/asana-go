package asanago

type StatusUpdate struct {
	Gid             string
	ResourceType    string `json:"resource_type"`
	HtmlText        string `json:"html_text"`
	ResourceSubtype string `json:"resource_subtype"`
	StatusType      string `json:"status_type"`
	Text            string
	Title           string
	Author          User
	CreatedAt       string `json:"created_at"`
	CreatedBy       User   `json:"created_by"`
	Liked           bool
	Likes           []Like
	ModifiedAt      string `json:"modified_at"`
	NumLikes        int
	Parent          ProjectCompact
}
