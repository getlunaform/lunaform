package workers

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func newProviderFile(filePath string) *ProviderFile {
	return &ProviderFile{
		filePath:  filePath,
		Providers: map[string]interface{}{},
	}
}

type ProviderFile struct {
	filePath  string
	Providers map[string]interface{}
}

func (pf *ProviderFile) String() string {
	return pf.RenderFileContent()
}

func (pf *ProviderFile) Byte() []byte {
	return []byte(pf.String())
}

func (pf *ProviderFile) WriteToFile() (err error) {
	return ioutil.WriteFile(pf.filePath, pf.Byte(), 0644)
}

func (pf *ProviderFile) RenderFileContent() string {
	provider := ""
	for key, value := range pf.Providers {
		provider = provider + `provider "` + key + `" {`
		switch v := value.(type) {
		case string:
			provider = provider + v
		case map[string]interface{}:
			provider = provider + pf.RenderMap("", v, 0)
		}
		provider = provider + "\n\n"
	}
	return provider
}

func (pf *ProviderFile) RenderString(key string, value string, indent int, equalIndent int) string {

	paddingFormat := fmt.Sprintf("%%-%ds", equalIndent)
	key = fmt.Sprintf(paddingFormat, key)

	return fmt.Sprintf(`%s%s = "%s"`,
		strings.Repeat(" ", indent*2),
		key, value)

}

func (pf *ProviderFile) RenderSlice(key string, content []string, indent int, equalIndent int) string {
	paddingFormat := fmt.Sprintf("%%-%ds", equalIndent)
	key = fmt.Sprintf(paddingFormat, key)

	return fmt.Sprintf(`%s%s = ["%s"]`,
		strings.Repeat(" ", indent*2),
		key, strings.Join(content, `", "`))

}

func (pf *ProviderFile) RenderMap(key string, c map[string]interface{}, indent int) string {
	padding := strings.Repeat(" ", indent*2)
	var response string
	if key == "" {
		response = "\n"
	} else {
		response = fmt.Sprintf("%s%s = {\n", padding, key)
	}

	longestKey := 0
	keys := []string{}
	for key, _ := range c {
		if len(key) > longestKey {
			longestKey = len(key)
		}
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		value := c[key]
		switch v := value.(type) {
		case string:
			line := pf.RenderString(key, v, indent+1, longestKey)
			response = response + line + "\n"
		case []string:
			line := pf.RenderSlice(key, v, indent+1, longestKey)
			response = response + line + "\n"
		case map[string]interface{}:
			line := pf.RenderMap(key, v, indent+1)
			response = response + line + "\n"
		}
	}
	return response + padding + "}"
}
