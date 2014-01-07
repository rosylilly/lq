package main

import (
	"fmt"
	"strings"
)

type LTSVEntry struct {
	keys   []string
	values map[string]string
}

func ParseLTSV(line string) *LTSVEntry {
	fields := strings.Split(line, "\t")
	keys := make([]string, len(fields))
	values := map[string]string{}

	for i, field := range fields {
		l := strings.Index(field, ":")
		keys[i] = field[0:l]
		values[field[0:l]] = field[l+1 : len(field)]
	}

	return &LTSVEntry{keys, values}
}

func (e *LTSVEntry) Filter(filters []string) *LTSVEntry {
	values := map[string]string{}

	for _, key := range filters {
		values[key] = e.values[key]
	}

	return &LTSVEntry{filters, values}
}

func (e *LTSVEntry) String() string {
	s := make([]string, len(e.keys))

	for i, key := range e.keys {
		s[i] = fmt.Sprintf("%s:%s", key, e.values[key])
	}

	return strings.Join(s, "\t")
}
