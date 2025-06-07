package http2curl

import (
	"fmt"
	"io"
	"os/exec"

	"github.com/kamil-koziol/restree/pkg/httpparser"
)

func Convert(body io.Reader, args ...string) (string, error) {
	req, err := httpparser.Parse(body)
	if err != nil {
		return "", fmt.Errorf("failed to parse .http: %s", err)
	}

	args = append(args, "-X", req.Method)
	for k, v := range req.Headers {
		args = append(args, "-H", fmt.Sprintf("'%s: %s'", k, v))
	}
	if req.Body != "" {
		args = append(args, "-d", fmt.Sprintf("'%s'", req.Body))
	}
	args = append(args, req.URL.String())

	curl := exec.Command("curl", args...)
	return curl.String(), nil
}
