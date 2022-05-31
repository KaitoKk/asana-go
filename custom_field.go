package asanago

type CustomField struct {
	Gid                     string
	ResourceType            string `json:"resource_type"`
	CreatedBy               User   `json:"created_by"`
	CurrencyType            string `json:"currency_type"`
	CustomLabel             string `json:"custom_label"`
	CustomLabelPosition     string `json:"custom_label_position"`
	Description             string
	DisplayValue            string `json:"display_value"`
	Enabled                 bool
	EnumOptions             []EnumOption `json:"enum_options"`
	EnumValue               EnumOption   `json:"enum_value"`
	format                  string
	HasNotificationsEnabled bool         `json:"has_notifications_enabled"`
	IsGlobalToWorkspace     bool         `json:"is_global_to_workspace"`
	MultiEnumValues         []EnumOption `json:"multi_enum_values"`
	Name                    string
	NumberValue             int
	Precision               int
	ResourceSubtype         string `json:"resource_subtype"`
	TextValue               string `json:"text_value"`
	Type                    string
}
