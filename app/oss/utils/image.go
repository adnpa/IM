package utils

import "strings"

var magicTable = map[string]string{
	"\xff\xd8\xff":      "image/jpeg",
	"\x89PNG\r\n\x1a\n": "image/png",
	"GIF87a":            "image/gif",
	"GIF89a":            "image/gif",
}

// mimeFromIncipit returns the mime type of an image file from its first few
// bytes or the empty string if the file does not look like a known file type
func mimeFromIncipit(incipit []byte) string {
	incipitStr := []byte(incipit)
	for magic, mime := range magicTable {
		if strings.HasPrefix(string(incipitStr), magic) {
			return mime
		}
	}
	return ""
}
