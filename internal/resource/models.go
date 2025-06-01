package resource

type Resource struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	Type       string   `json:"type"`
	Repository string   `json:"repository"`
	Tags       []string `json:"tags,omitempty"`
	Created    string   `json:"created,omitempty"`
	Size       string   `json:"size,omitempty"`
	Digest     string   `json:"digest,omitempty"`
	Owner      string   `json:"owner,omitempty"`
}
