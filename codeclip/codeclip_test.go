package codeclip

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRendText(t *testing.T) {

	// test 1
	template := "Your account #{account} will expired, but #{password} will not"
	content := map[string]string{
		"account": "test",
	}
	expected := "Your account test will expired, but #{password} will not"
	actual := RendText(template, content)
	assert.Equal(t, expected, actual)

	// test 2
	template = "Your account #{account} will expired, but #{password} will not."
	content = map[string]string{
		"account":  "test",
		"password": "test2",
	}
	expected = "Your account test will expired, but test2 will not."
	actual = RendText(template, content)
	assert.Equal(t, expected, actual)

	// test 3
	template = "Your account #{account} will expired, but #{password} will not. It's #{account}"
	content = map[string]string{
		"account": "test",
	}
	expected = "Your account test will expired, but #{password} will not. It's test"
	actual = RendText(template, content)
	assert.Equal(t, expected, actual)

	// test 4
	template = "Your account #{account} will expired, but #{password} will not. It's #{account}"
	content = map[string]string{
		"account": "test",
		"acount2": "test2",
	}
	expected = "Your account test will expired, but #{password} will not. It's test"
	actual = RendText(template, content)
	assert.Equal(t, expected, actual)

	// test 5
	template = "你的账户#{account}将要过期, 但是#{password}不会."
	content = map[string]string{
		"account": "test",
	}
	expected = "你的账户test将要过期, 但是#{password}不会."
	actual = RendText(template, content)
	assert.Equal(t, expected, actual)

}
