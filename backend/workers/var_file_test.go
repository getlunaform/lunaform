package workers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableFile_Build(t *testing.T) {

	for _, tt := range []struct {
		name      string
		variables map[string]string
		want      string
	}{
		{
			name: "String",
			variables: map[string]string{
				"hello": "world",
			},
			want: `variable "hello" {
    default = "world"
}
`,
		},
		{
			name: "Slice",
			variables: map[string]string{
				"hello": "[foo,bar]",
			},
			want: `variable "hello" {
    default = ["foo", "bar"]
    type = "list"
}
`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			vf := &VariableFile{
				variables: tt.variables,
			}
			assert.Equal(t, tt.want, vf.Build())
		})
	}
}

func TestVariableFile_ParseSlice(t *testing.T) {
	for _, tt := range []struct {
		name      string
		raw       string
		wantSlice []string
	}{
		{
			name:      "basic",
			raw:       `["foo","bar"]`,
			wantSlice: []string{"foo", "bar"},
		},
		{
			name:      "escape_char",
			raw:       `["foo\"","bar"]`,
			wantSlice: []string{"foo\"", "bar"},
		},
		{
			name:      "ints",
			raw:       `[1,2,3]`,
			wantSlice: []string{"1", "2", "3"},
		},
		{
			name:      "mixed",
			raw:       `["1one",2,"three"]"`,
			wantSlice: []string{"1one", "2", "three"},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			vf := &VariableFile{}
			gotSlice := vf.ParseSlice(tt.raw)
			assert.Equal(t, tt.wantSlice, gotSlice)
		})
	}
}

func TestVariableFile_ParseMap(t *testing.T) {
	for _, tt := range []struct {
		name          string
		raw           string
		wantStringMap map[string]string
	}{
		{
			name: "basic",
			raw:  `{ foo = "bar", baz = "qux" }`,
			wantStringMap: map[string]string{
				"foo": "bar",
				"baz": "qux",
			},
		},
		{
			name: "no-spaces",
			raw:  `{foo="bar",baz="qux"}`,
			wantStringMap: map[string]string{
				"foo": "bar",
				"baz": "qux",
			},
		},
		{
			name: "ints",
			raw:  `{foo=1,baz="qux"}`,
			wantStringMap: map[string]string{
				"foo": "1",
				"baz": "qux",
			},
		},
		{
			name: "escape",
			raw:  `{foo="bar\""}`,
			wantStringMap: map[string]string{
				"foo": "bar\"",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			vf := &VariableFile{}
			assert.Equal(t, tt.wantStringMap, vf.ParseMap(tt.raw))
		})
	}
}
