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

	tests := []struct {
		given body
		want  string
	}{
		{
			given: body{"name": "anson", "age": "18"}, want: "name%3Danson%26age%3D18",
		},
	}

	for _, tc := range tests {
		assert.Equal(t, tc.given.EncodeToString(), tc.want)
	}
}
