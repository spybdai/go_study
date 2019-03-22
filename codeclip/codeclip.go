package codeclip

import (
	"strings"
)

// render text with templates
// template like:
// "Your account #{userAccount} will expire soon"
// and content is a map[string]string, with same key as in the template

func RendText(template string, content map[string]string) string {
	for k, v := range content {
		template = strings.Replace(template, strings.Join([]string{"#{", k, "}"}, ""), v, -1)
	}
	return template
}
