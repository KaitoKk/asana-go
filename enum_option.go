package asanago

type EnumOption struct {
	Gid          string
	ResourceType string `json:"resource_type"`
	Color        string
	Enabled      bool
	Name         string
}
