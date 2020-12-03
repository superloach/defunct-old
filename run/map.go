package run

import "strings"

type Map map[string]Funct

func (m Map) String() string {
	ss := make([]string, len(m))

	for k, v := range m {
		ss = append(ss, k+": "+string(TextOf(v)))
	}

	return "map " + strings.Join(ss, ", ")
}

func (m Map) Call(_ *Under, _ Queue) (Funct, error) {
	panic("map call stub")
}
