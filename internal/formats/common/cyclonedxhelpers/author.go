package cyclonedxhelpers

import (
	"fmt"
	"strings"

	"github.com/anchore/syft/syft/pkg"
)

func Author(p pkg.Package) string {
	if hasMetadata(p) {
		switch metadata := p.Metadata.(type) {
		case pkg.NpmPackageJSONMetadata:
			return metadata.Author
		case pkg.PythonPackageMetadata:
			author := metadata.Author
			if metadata.AuthorEmail != "" {
				if author == "" {
					return metadata.AuthorEmail
				}
				author += fmt.Sprintf(" <%s>", metadata.AuthorEmail)
			}
			return author
		case pkg.GemMetadata:
			if len(metadata.Authors) > 0 {
				return strings.Join(metadata.Authors, ",")
			}
			return ""
		}
	}
	return ""
}
