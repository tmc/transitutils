package transitutils

import (
	"fmt"
)

// ToGo converts a transit type into standard go types.
func ToGo(t interface{}) (interface{}, error) {
	switch t := t.(type) {
	case map[interface{}]interface{}:
		r := make(map[string]interface{}, 0)
		for k, v := range t {
			v2, err := ToGo(v)
			if err != nil {
				// return a partial result when possible
				return r, err
			}
			r[fmt.Sprint(k)] = v2
		}
		return r, nil
	case []interface{}:
		r := make([]interface{}, 0)
		for _, v := range t {
			v2, err := ToGo(v)
			if err != nil {
				// return a partial result when possible
				return r, err
			}
			r = append(r, v2)
		}
		return r, nil
	default:
		return t, nil
	}
}
