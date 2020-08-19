package got

import (
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
)

// DefaultFileName is the fallback name for GetFilename.
var DefaultFileName = "got.output"

// GetFilename it returns default file name from a URL.
func GetFilename(URL string) string {

	if u, err := url.Parse(URL); err == nil && filepath.Ext(u.Path) != "" {

		return filepath.Base(u.Path)
	}

	res, err := http.Head(URL)
	if err == nil {
		header := res.Header
		if hcd, ok := header["Content-Disposition"]; ok && len(hcd) > 0 {
			hcds := strings.Split(hcd[0], "=")
			if len(hcds) > 1 {
				if filename := hcds[1]; filename != "" {
					return filepath.Base(filename)
				}
			}
		}
	}

	return DefaultFileName
}
