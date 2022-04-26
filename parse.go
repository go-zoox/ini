package ini

import (
	"regexp"
	"strings"
)

var (
	// spaceRe   = regexp.MustCompile(`\s+`)
	sectionRe = regexp.MustCompile(`^\[(.+)\]$`)
)

// Parse parses the ini data and returns a map.
func Parse(bytes []byte) (map[string]interface{}, error) {
	container := make(map[string]interface{})

	linesX := strings.Split(string(bytes), "\n")
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
			container[key] = value
		} else {
			if _, ok := container[currentSection]; !ok {
				container[currentSection] = make(map[string]interface{})
			}

			container[currentSection].(map[string]interface{})[key] = value
		}
	}

	return container, nil
}
