package ini

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/go-zoox/tag"
	"github.com/go-zoox/tag/datasource"
)

// Marshal returns the ini data of the given struct pointer.
func Marshal(v interface{}) ([]byte, error) {
	// 1. any(struct / map) => json
	jsonv, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	// 2. json => map[string]interface{}
	var m map[string]interface{}
	if err := json.Unmarshal(jsonv, &m); err != nil {
		return nil, err
	}

	// 3. map[string]interface{} => string
	lines := []string{}

	// why this ? ensure key sorted
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		key := strings.ToLower(k)
		value := m[k]

		if value == nil {
			continue
		}

		switch value.(type) {
		case string:
			lines = append(lines, fmt.Sprintf("%s = %s", key, value))
		case bool:
			lines = append(lines, fmt.Sprintf("%s = %t", key, value))
		case int:
			lines = append(lines, fmt.Sprintf("%s = %d", key, value.(int)))
		case int64:
			lines = append(lines, fmt.Sprintf("%s = %d", key, value.(int64)))
		case float64:
			// lines = append(lines, fmt.Sprintf("%s = %f", k, v.(float64)))
			lines = append(lines, fmt.Sprintf("%s = %d", key, int64(value.(float64))))
		case map[string]interface{}:
			// empty line
			lines = append(lines, "")
			//
			lines = append(lines, "["+key+"]")

			m2 := value.(map[string]interface{})
			// why this ? ensure key sorted
			keys2 := []string{}
			for k := range m2 {
				keys2 = append(keys2, k)
			}
			sort.Strings(keys2)

			for _, k2 := range keys2 {
				kk, vv := strings.ToLower(k2), m2[k2]
				switch vv.(type) {
				case string:
					lines = append(lines, fmt.Sprintf("%s = %s", kk, vv))
				case bool:
					lines = append(lines, fmt.Sprintf("%s = %t", kk, vv))
				case int:
					lines = append(lines, fmt.Sprintf("%s = %d", kk, vv.(int)))
				case int64:
					lines = append(lines, fmt.Sprintf("%s = %d", kk, vv.(int64)))
				case float64:
					// lines = append(lines, fmt.Sprintf("%s = %f", kk, vv.(float64)))
					lines = append(lines, fmt.Sprintf("%s = %d", kk, int64(vv.(float64))))
				default:
					return nil, fmt.Errorf("unsupported type: %T", vv)
				}
			}
		default:
			return nil, fmt.Errorf("unsupported type: %T", value)
		}
	}

	return []byte(strings.Join(lines, "\n")), nil
}

// Unmarshal parses the ini data and stores the result in the value pointed to by v.
func Unmarshal(data []byte, v interface{}) error {
	if ds, err := Parse(data); err != nil {
		return err
	} else {
		tg := tag.New("ini", datasource.NewMapDataSource(ds))
		return tg.Decode(v)
	}
}
