package workers

import (
	"text/template"
	"bytes"
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
)

type VariableFile struct {
	//filePath  string
	variables map[string]string
}

func (vf *VariableFile) Build() string {
	templateStr := `{{range $k,$v := .}}variable "{{$k}}" {
    default = "{{$v}}"
}
{{end}}`
	templ, err := template.New("variables").Parse(templateStr)
	if err != nil {
		panic(err)
	}
	var tpl bytes.Buffer
	err = templ.Execute(&tpl, vf.variables)
	if err != nil {
		panic(err)
	}
	return tpl.String()
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
