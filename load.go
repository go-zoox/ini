package ini

import (
	"regexp"
	"strings"
)

var (
	spaceRe   = regexp.MustCompile(`\s+`)
	sectionRe = regexp.MustCompile(`^\[(.+)\]$`)
)

type datasource struct {
	bytes     []byte
	structure map[string]interface{}
}

func newDataSource(data []byte) *datasource {
	return &datasource{
		structure: make(map[string]interface{}),
		bytes:     data,
	}
}

func (ds *datasource) Parse() error {
	linesX := strings.Split(string(ds.bytes), "\n")
	lines := []string{}
	for _, line := range linesX {
		// emty line
		if len(strings.TrimSpace(line)) == 0 {
			continue
		} else if strings.HasPrefix(line, "#") {
			continue
		}

		lines = append(lines, line)
	}

	currentSection := ""
	for _, line := range lines {
		if sectionRe.MatchString(line) {
			currentSection = sectionRe.FindStringSubmatch(line)[1]
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		key, value := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
		if currentSection == "" {
			ds.structure[key] = value
		} else {
			if _, ok := ds.structure[currentSection]; !ok {
				ds.structure[currentSection] = make(map[string]interface{})
			}

			ds.structure[currentSection].(map[string]interface{})[key] = value
		}
	}

	return nil
}

func (ds *datasource) Get(key string) interface{} {
	if strings.Contains(key, ".") {
		keys := strings.Split(key, ".")
		keyLength := len(keys)
		tmp := ds.structure
		for index, k := range keys {
			if v, ok := tmp[k]; ok {
				if index == keyLength-1 {
					return v
				}

				tmp = v.(map[string]interface{})
			} else {
				return nil
			}
		}
	} else {
		if v, ok := ds.structure[key]; ok {
			return v
		}
	}

	return nil
}
