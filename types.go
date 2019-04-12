package utils

// http
type (
	MultipartField struct {
		IsFile    bool
		Fieldname string
		Value     []byte
		Filename  string
	}
)
