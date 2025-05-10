package headers

import (
	"bytes"
	"errors"
	"strings"
)

type Headers map[string]string

func NewHeaders() Headers {
	return Headers{}
}

func (h Headers) Get(key string) string {
	return h[strings.ToLower(key)]
}

func (h Headers) Set(key, value string) {
	h[key] = value
}

const crlf = "\r\n"

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	// incomplete line
	if idx == -1 {
		return 0, false, nil
	}
	for _, line := range bytes.Split(data, []byte(crlf)) {
		if len(line) == 0 {
			return n + 2, true, nil
		}
		parsed, err := h.parseHeaderFromString(string(line))
		// error parsing header line
		if err != nil {
			return 0, false, err
		}
		n += parsed
	}
	return n + 2, false, nil
}

func (h *Headers) parseHeaderFromString(s string) (int, error) {
	m := *h
	n := len(s)
	s = strings.TrimSpace(s)
	ErrInvalidHeader := errors.New("Invalid header")

	fieldParts := strings.SplitN(s, ":", 2)
	// not correct number of parts
	if len(fieldParts) != 2 {
		return 0, nil
	}

	// for field name
	if len(strings.Split(fieldParts[0], " ")) != 1 {
		return 0, ErrInvalidHeader
	}
	fieldName := strings.ToLower(strings.TrimSpace(fieldParts[0]))
	if !isValidToken(fieldName) {
		return 0, ErrInvalidHeader
	}

	fieldValue := strings.TrimSpace(fieldParts[1])
	prevValue, ok := m[fieldName]
	if !ok {
		m.Set(fieldName, fieldValue)
	} else {
		newValue := strings.Join([]string{prevValue, fieldValue}, ", ")
		m.Set(fieldName, newValue)
	}

	return n, nil
}

func isValidToken(s string) bool {
	const specials = "!#$%&'*+-.^_`|~"
	for _, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
		case r >= 'a' && r <= 'z':
		case r >= '0' && r <= '9':
		case strings.ContainsRune(specials, r):
		default:
			return false
		}
	}
	return true
}
