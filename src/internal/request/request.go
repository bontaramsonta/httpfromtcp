package request

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	requestBytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(requestBytes, []byte("\r\n"))

	// read request line
	var RequestLine RequestLine
	requestLineParts := bytes.Split(lines[0], []byte(" "))
	if len(requestLineParts) != 3 {
		return nil, errors.New("Invalid request line")
	}

	for i, field := range requestLineParts {
		switch i {
		case 0:
			switch string(field) {
			case "GET":
				RequestLine.Method = "GET"
			case "POST":
				RequestLine.Method = "POST"
			case "PUT":
				RequestLine.Method = "PUT"
			case "DELETE":
				RequestLine.Method = "DELETE"
			case "HEAD":
				RequestLine.Method = "HEAD"
			case "OPTIONS":
				RequestLine.Method = "OPTIONS"
			case "TRACE":
				RequestLine.Method = "TRACE"
			case "CONNECT":
				RequestLine.Method = "CONNECT"
			default:
				return nil, errors.New("invalid method")
			}
		case 1:
			RequestLine.RequestTarget = string(field)
		case 2:
			parts := bytes.SplitN(field, []byte("/"), 2)
			if string(parts[0]) != "HTTP" || len(parts) != 2 {
				return nil, errors.New("Error HTTP Version")
			}
			if string(parts[1]) != "1.1" {
				return nil, errors.New("Unsupported HTTP Version")
			}

			RequestLine.HttpVersion = string(parts[1])
		}
	}

	// read headers
	headers := make(map[string]string)
	headerLines := lines[1 : len(lines)-2] // don't take last 2 lines (body line and empty line)
	for _, line := range headerLines {
		parts := bytes.SplitN(line, []byte(":"), 2)
		key := strings.ToLower(string(parts[0]))
		val := strings.TrimSpace(string(parts[1]))
		if len(parts) != 2 {
			return nil, errors.New("invalid header")
		}
		headers[key] = val
	}

	return &Request{RequestLine: RequestLine}, nil
}
