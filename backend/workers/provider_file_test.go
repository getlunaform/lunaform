package workers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProviderFile_RenderMap(t *testing.T) {
	for _, tt := range []struct {
		name    string
		key     string
		content map[string]interface{}
		indent  int
		want    string
	}{
		{
			name: "basic",
			key:  "key",
			content: map[string]interface{}{
				"one":   "1",
				"two":   "2",
				"three": "3",
			},
			indent: 0,
			want: `key = {
  one   = "1"
  three = "3"
  two   = "2"
}`,
		},
		{
			name: "basic-indented",
			key:  "key",
			content: map[string]interface{}{
				"one":   "1",
				"two":   "2",
				"three": "3",
			},
			indent: 2,
			want: `    key = {
      one   = "1"
      three = "3"
      two   = "2"
    }`,
		},
		{
			name: "slice-indented",
			key:  "key",
			content: map[string]interface{}{
				"one":  "1",
				"two":  []string{"2", "3", "4"},
				"five": "5",
			},
			indent: 2,
			want: `    key = {
      five = "5"
      one  = "1"
      two  = ["2", "3", "4"]
    }`,
		},
		{
			name: "slice-map-indented",
			key:  "key",
			content: map[string]interface{}{
				"one": "1",
				"two": []string{"2", "3", "4"},
				"five": map[string]interface{}{
					"six": "6",
				},
			},
			indent: 2,
			want: `    key = {
      five = {
        six = "6"
      }
      one  = "1"
      two  = ["2", "3", "4"]
    }`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			pf := &ProviderFile{}
			got := pf.RenderMap(tt.key, tt.content, tt.indent)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProviderFile_RenderString(t *testing.T) {
	for _, tt := range []struct {
		name    string
		content string
		key     string
		indent  int
		want    string
	}{
		{
			name:    "indent-0",
			content: "content",
			key:     "key",
			indent:  0,
			want:    `key = "content"`,
		},
		{
			name:    "indent-1",
			content: "content",
			key:     "key",
			indent:  1,
			want:    `  key = "content"`,
		},
		{
			name:    "indent-2",
			content: "content",
			key:     "key",
			indent:  2,
			want:    `    key = "content"`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			pf := &ProviderFile{}
			got := pf.RenderString(tt.key, tt.content, tt.indent, 0)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProviderFile_RenderSlice(t *testing.T) {
	for _, tt := range []struct {
		name    string
		key     string
		content []string
		indent  int
		want    string
	}{
		{
			name:   "basic",
			key:    "key",
			indent: 0,
			want:   `key = ["one", "two", "three"]`,
			content: []string{
				"one", "two", "three",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			pf := &ProviderFile{}
			got := pf.RenderSlice(tt.key, tt.content, tt.indent, 0)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestProviderFile_RenderFileContent(t *testing.T) {
	for _, tt := range []struct {
		name      string
		providers map[string]interface{}
		want      string
	}{
		{
			name: "basic",
			providers: map[string]interface{}{
				"provider-1": map[string]interface{}{
					"one": "1",
					"two": "2",
				},
				"provider-2": map[string]interface{}{
					"three": "3",
					"four":  "4",
				},
			},
			want: `provider "provider-1" {
  one = "1"
  two = "2"
}

provider "provider-2" {
  four  = "4"
  three = "3"
}

`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			pf := &ProviderFile{Providers: tt.providers}
			got := pf.RenderFileContent()
			assert.Equal(t, tt.want, got)
		})
	}
}
