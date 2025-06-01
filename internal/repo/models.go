package repo

// Repository represents a source of artifacts
type Repository struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string `json:"description,omitempty"`
}
