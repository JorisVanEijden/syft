package cyclonedxhelpers

import (
	"testing"

	"github.com/anchore/syft/syft/pkg"
	"github.com/stretchr/testify/assert"
)

func Test_CPE(t *testing.T) {
	testCPE := pkg.MustCPE("cpe:2.3:a:name:name:3.2:*:*:*:*:*:*:*")
	testCPE2 := pkg.MustCPE("cpe:2.3:a:name:name2:3.2:*:*:*:*:*:*:*")
	tests := []struct {
		name     string
		input    pkg.Package
		expected string
	}{
		{
			// note: since this is an optional field, no value is preferred over NONE or NOASSERTION
			name: "no metadata",
			input: pkg.Package{
				CPEs: []pkg.CPE{},
			},
			expected: "",
		},
		{
			name: "single CPE",
			input: pkg.Package{
				CPEs: []pkg.CPE{
					testCPE,
				},
			},
			expected: "cpe:2.3:a:name:name:3.2:*:*:*:*:*:*:*",
		},
		{
			name: "multiple CPEs",
			input: pkg.Package{
				CPEs: []pkg.CPE{
					testCPE2,
					testCPE,
				},
			},
			expected: "cpe:2.3:a:name:name2:3.2:*:*:*:*:*:*:*",
		},
		{
			// note: since this is an optional field, no value is preferred over NONE or NOASSERTION
			name:     "empty",
			input:    pkg.Package{},
			expected: "",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, CPE(test.input))
		})
	}
}
