package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBodyEncodeToString(t *testing.T) {
	type Person struct {
		Name string `json:"name"`
	}

	b := newBody(Config{}, Person{Name: "Anson"})
	str := b.EncodeToString()
	assert.Contains(t, str, `%7B%22name%22%3A%22Anson%22%7D`)
}
