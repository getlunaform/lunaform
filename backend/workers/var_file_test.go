package workers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariableFile_Build(t *testing.T) {

	for _, tt := range []struct {
		name      string
		varType   string
		variables map[string]*VariableFileEntry
		want      string
	}{
		{
			name:    "string-tf",
			varType: VARIABLE_FILE_TYPE_TF,
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_STRING, String: "world"},
			},

			want: `variable "hello" {
  type = "string"

  default = "world"
}

`,
		},
		{
			name:    "slice-tf",
			varType: VARIABLE_FILE_TYPE_TF,
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_SLICE, Slice: []string{"foo", "bar"}},
			},
			want: `variable "hello" {
  type = "list"

  default = [
    "foo",
    "bar"
  ]
}

`,
		},
		{
			name:    "map-tf",
			varType: VARIABLE_FILE_TYPE_TF,
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_MAP, Map: map[string]string{"foo": "bar"}},
			},
			want: `variable "hello" {
  type = "map"

  default = {
    foo = "bar"
  }
}

`,
		},
		{
			name:    "mixed-tf",
			varType: VARIABLE_FILE_TYPE_TF,
			variables: map[string]*VariableFileEntry{
				"map":    {Type: VARIABLE_TYPE_MAP, Map: map[string]string{"foo": "bar"}},
				"slice":  {Type: VARIABLE_TYPE_SLICE, Slice: []string{"foo", "bar"}},
				"string": {Type: VARIABLE_TYPE_STRING, String: "world"},
			},
			want: `variable "map" {
  type = "map"

  default = {
    foo = "bar"
  }
}

variable "slice" {
  type = "list"

  default = [
    "foo",
    "bar"
  ]
}

variable "string" {
  type = "string"

  default = "world"
}

`,
		},
		{
			name:    "string-tfvars",
			varType: VARIABLE_FILE_TYPE_TFVARS,
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_STRING, String: "world"},
			},

			want: `hello = "world"

`,
		},
		{
			name:    "slice-tfvars",
			varType: VARIABLE_FILE_TYPE_TFVARS,
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_SLICE, Slice: []string{"foo", "bar"}},
			},
			want: `hello = [
    "foo",
    "bar"
  ]

`,
		},
		{
			name:    "map-tfvars",
			varType: VARIABLE_FILE_TYPE_TFVARS,
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_MAP, Map: map[string]string{"foo": "bar"}},
			},
			want: `hello = {
    foo = "bar"
  }

`,
		},
		{
			name:    "mixed-tfvars",
			varType: VARIABLE_FILE_TYPE_TFVARS,
			variables: map[string]*VariableFileEntry{
				"map":    {Type: VARIABLE_TYPE_MAP, Map: map[string]string{"foo": "bar"}},
				"slice":  {Type: VARIABLE_TYPE_SLICE, Slice: []string{"foo", "bar"}},
				"string": {Type: VARIABLE_TYPE_STRING, String: "world"},
			},
			want: `map = {
    foo = "bar"
  }

slice = [
    "foo",
    "bar"
  ]

string = "world"

`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			vf := &VariableFile{
				Variables: tt.variables,
				fileType:  tt.varType,
			}
			assert.Equal(t, tt.want, string(vf.Byte()))
		})
	}
}

func TestVariableFile_BuildTfVars(t *testing.T) {

	for _, tt := range []struct {
		name      string
		variables map[string]*VariableFileEntry
		want      string
	}{
		{
			name: "string",
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_STRING, String: "world"},
			},
			want: `variable "hello" {
  type = "string"

  default = "world"
}

`,
		},
		{
			name: "slice",
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_SLICE, Slice: []string{"foo", "bar"}},
			},
			want: `variable "hello" {
  type = "list"

  default = [
    "foo",
    "bar"
  ]
}

`,
		},
		{
			name: "map",
			variables: map[string]*VariableFileEntry{
				"hello": {Type: VARIABLE_TYPE_MAP, Map: map[string]string{"foo": "bar"}},
			},
			want: `variable "hello" {
  type = "map"

  default = {
    foo = "bar"
  }
}

`,
		},
		{
			name: "mixed",
			variables: map[string]*VariableFileEntry{
				"map":    {Type: VARIABLE_TYPE_MAP, Map: map[string]string{"foo": "bar"}},
				"slice":  {Type: VARIABLE_TYPE_SLICE, Slice: []string{"foo", "bar"}},
				"string": {Type: VARIABLE_TYPE_STRING, String: "world"},
			},
			want: `variable "map" {
  type = "map"

  default = {
    foo = "bar"
  }
}

variable "slice" {
  type = "list"

  default = [
    "foo",
    "bar"
  ]
}

variable "string" {
  type = "string"

  default = "world"
}

`,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			vf := &VariableFile{
				Variables: tt.variables,
			}
			assert.Equal(t, tt.want, string(vf.Byte()))
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
