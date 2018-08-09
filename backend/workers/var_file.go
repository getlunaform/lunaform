package workers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"text/template"
)

const (
	PARSER_STATE_SLICE_START = 0

	PARSER_STATE_MAP_START       = 1
	PARSER_STATE_MAP_KEY_START   = 2
	PARSER_STATE_MAP_KEY         = 3
	PARSER_STATE_MAP_VALUE_START = 4
	PARSER_STATE_MAP_VALUE       = 5

	PARSER_STATE_LOOKING_FOR_STRING_START = 10
	PARSER_STATE_LOOKING_FOR_STRING       = 11

	PARSER_STATE_ESCAPE_START  = 20
	PARSER_STATE_ESCAPE_STRING = 21

	VARIABLE_TYPE_STRING = "string"
	VARIABLE_TYPE_MAP    = "map"
	VARIABLE_TYPE_SLICE  = "list"

	VARIABLE_FILE_TYPE_TF     = "tf"
	VARIABLE_FILE_TYPE_TFVARS = "tfvars"
)

func newVariableFile(filePath string) *VariableFile {
	return newVariableFileWithType(filePath, VARIABLE_FILE_TYPE_TF)
}

func newVariableFileWithType(filePath string, fileType string) *VariableFile {
	return &VariableFile{
		filePath: filePath,
		fileType: fileType,
	}
}

type VariableFile struct {
	filePath  string
	fileType  string
	Variables map[string]*VariableFileEntry
}

type VariableFileEntry struct {
	Type   string
	Slice  []string
	Map    map[string]string
	String string
}

func (vfe *VariableFileEntry) Value() string {
	switch vfe.Type {
	case VARIABLE_TYPE_STRING:
		return fmt.Sprintf("\"%s\"", vfe.String)
	case VARIABLE_TYPE_MAP:
		return vfe.mapToString()
	case VARIABLE_TYPE_SLICE:
		return fmt.Sprintf("[\n    \"%s\"\n  ]", strings.Join(vfe.Slice, "\",\n    \""))
	}
	return ""
}

func (vfe *VariableFileEntry) mapToString() string {
	mapEntries := make([]string, 0)
	for key, value := range vfe.Map {
		mapEntries = append(mapEntries, fmt.Sprintf("%s = \"%s\"", key, value))
	}
	return fmt.Sprintf("{\n    %s\n  }", strings.Join(mapEntries, ","))
}

func (vf *VariableFile) Parse(variables map[string]string) {
	vf.Variables = make(map[string]*VariableFileEntry)
	for key, variable := range variables {
		if vf.IsMap(variable) {
			vf.Variables[key] = &VariableFileEntry{Type: VARIABLE_TYPE_MAP, Map: vf.ParseMap(variable)}
		} else if vf.IsSlice(variable) {
			vf.Variables[key] = &VariableFileEntry{Type: VARIABLE_TYPE_SLICE, Slice: vf.ParseSlice(variable)}
		} else if vf.IsString(variable) {
			vf.Variables[key] = &VariableFileEntry{Type: VARIABLE_TYPE_STRING, String: variable}
		}
	}
}

func (vf *VariableFile) Byte() []byte {
	templates := map[string]string{
		VARIABLE_FILE_TYPE_TF: `{{range $k,$v := .Variables}}variable "{{$k}}" {
  type = "{{$v.Type}}"

  default = {{$v.Value}}
}

{{end}}`,
		VARIABLE_FILE_TYPE_TFVARS: `{{range $k,$v := .Variables}}{{$k}} = {{$v.Value}}

{{end}}`,
	}

	templ, err := template.New("variables").Funcs(template.FuncMap{
		"last": func(x int, a interface{}) bool {
			return x == reflect.ValueOf(a).Len()-1
		},
	}).Parse(templates[vf.fileType])
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = templ.Execute(&tpl, vf)
	if err != nil {
		panic(err)
	}
	b := tpl.Bytes()
	return b
}

func (vf *VariableFile) String() string {
	return string(vf.Byte())
}

func (vf *VariableFile) WriteToFile() (err error) {
	return ioutil.WriteFile(vf.filePath, vf.Byte(), 0644)
}

func (vf *VariableFile) IsSlice(raw string) bool {
	return strings.HasPrefix(raw, "[")
}

func (vf *VariableFile) IsMap(raw string) bool {
	return strings.HasPrefix(raw, "{")
}

func (vf *VariableFile) IsString(raw string) bool {
	return !vf.IsSlice(raw) && !vf.IsMap(raw)
}

func (vf *VariableFile) ParseSlice(raw string) (slice []string) {
	slice = make([]string, 0)
	state := PARSER_STATE_SLICE_START
	escapeState := PARSER_STATE_ESCAPE_START

	var currentElement string
	for pos, char := range raw {
		if state == PARSER_STATE_SLICE_START {
			if (pos == 0) && char == '[' {
				state = PARSER_STATE_LOOKING_FOR_STRING_START
			}
		} else if state == PARSER_STATE_LOOKING_FOR_STRING_START {
			if char == '"' {
				state = PARSER_STATE_LOOKING_FOR_STRING
			} else if char == ',' {
				continue
			} else if char == ']' {
				state = PARSER_STATE_SLICE_START
			} else {
				slice = append(slice, string(char))
			}
		} else if state == PARSER_STATE_LOOKING_FOR_STRING {
			if escapeState == PARSER_STATE_ESCAPE_STRING {
				currentElement = currentElement + string(char)
				escapeState = PARSER_STATE_ESCAPE_START
			} else if char == '\\' {
				escapeState = PARSER_STATE_ESCAPE_STRING
			} else if char == '"' {
				state = PARSER_STATE_LOOKING_FOR_STRING_START
				slice = append(slice, currentElement)
				currentElement = ""
			} else {
				currentElement = currentElement + string(char)
			}
		}

	}
	return
}

func (vf *VariableFile) ParseMap(raw string) (stringMap map[string]string) {
	stringMap = make(map[string]string)
	state := PARSER_STATE_MAP_START
	escapeState := PARSER_STATE_ESCAPE_START

	var currentElement, currentKey string
	for pos, char := range raw {
		if state == PARSER_STATE_MAP_START {
			if (pos == 0) && char == '{' || char == ',' {
				state = PARSER_STATE_MAP_KEY_START
			}
		} else if state == PARSER_STATE_MAP_KEY_START {
			if char == ' ' {
				state = PARSER_STATE_MAP_KEY
			} else if char != '"' && char != ' ' {
				currentKey = currentKey + string(char)
				state = PARSER_STATE_MAP_KEY
			}
		} else if state == PARSER_STATE_MAP_KEY {
			if char != '"' && char != ' ' && char != '=' {
				currentKey = currentKey + string(char)
			} else {
				state = PARSER_STATE_MAP_VALUE_START
			}
		} else if state == PARSER_STATE_MAP_VALUE_START {
			if char == '=' {
				continue
			} else if char == '"' {
				state = PARSER_STATE_MAP_VALUE
			} else if char == ' ' {
				continue
			} else {
				stringMap[currentKey] = string(char)
				currentKey = ""
				currentElement = ""
				state = PARSER_STATE_MAP_START
			}
		} else if state == PARSER_STATE_MAP_VALUE {
			if escapeState == PARSER_STATE_ESCAPE_STRING {
				currentElement = currentElement + string(char)
				escapeState = PARSER_STATE_ESCAPE_START
			} else if char != '"' && char != ' ' && char != '\\' {
				currentElement = currentElement + string(char)
			} else if char == '\\' {
				escapeState = PARSER_STATE_ESCAPE_STRING
			} else if char == '"' {
				stringMap[currentKey] = currentElement
				currentKey = ""
				currentElement = ""
				state = PARSER_STATE_MAP_START
			}
		}
	}
	return
}
