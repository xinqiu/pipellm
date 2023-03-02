package types

// Document struct
type Document struct {
	PageContent string
	LookupStr   string
	LookupIndex int64
	Metadata    map[string]interface{}
}

// NewDocument creates a new document
func NewDocument(pageContent string, metadata map[string]interface{}) *Document {
	return &Document{
		PageContent: pageContent,
		Metadata:    metadata,
	}
}
