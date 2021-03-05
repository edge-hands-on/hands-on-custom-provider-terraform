package types

type Dashboard struct {
	Title             string             `json:"title"`
	Description       string             `json:"description"`
	IsReadOnly        bool               `json:"is_read_only"`
	LayoutType        string             `json:"layout_type"`
	TemplateVariables []TemplateVariable `json:"template_variables"`
	GroupWidget       []GroupWidget      `json:"widgets"`
}

type TemplateVariable struct {
	Name    string `json:"name"`
	Default string `json:"default"`
	Prefix  string `json:"prefix"`
}

type GroupWidget struct {
	Definition GroupWidgetDefinition `json:"definition"`
}

type GroupWidgetDefinition struct {
	LayoutType string   `json:"layout_type"`
	Title      string   `json:"title"`
	Type       string   `json:"type"`
	Widgets    []Widget `json:"widgets"`
}

type Widget struct {
	Definition interface{} `json:"definition"`
}

type TimeseriesWidget struct {
	Type     string              `json:"type"`
	Title    string              `json:"title"`
	Requests []TimeseriesRequest `json:"requests"`
}

type TimeseriesRequest struct {
	Q           string `json:"q"`
	DisplayType string `json:"display_type"`
}

type ToplistWidget struct {
	Type     string           `json:"type"`
	Title    string           `json:"title"`
	Requests []ToplistRequest `json:"requests"`
}

type ToplistRequest struct {
	Q string `json:"q"`
}

type QueryValueWidget struct {
	Type      string              `json:"type"`
	Title     string              `json:"title"`
	Autoscale bool                `json:"autoscale"`
	Precision int                 `json:"precision"`
	Requests  []QueryValueRequest `json:"requests"`
}

type QueryValueRequest struct {
	Q          string `json:"q"`
	Aggregator string `json:"aggregator"`
}
