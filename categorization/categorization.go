package categorization

type Categorization struct {
	Version string            `json:"version"`
	Mapping map[string]string `json:"mapping"`
}
