package request

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSign(t *testing.T) {
	tests := []struct {
		given []string
		want  string
	}{
		{given: []string{"hello1234", "key123"}, want: "ufYU7rvXhHY3IDyZgyt6SA=="},
		{given: []string{`{"name":"anson","age":18}`, "key123"}, want: "znNaBxySyki5gNhjSzsJsA=="},
	}

	for _, tc := range tests {
		content := tc.given[0]
		key := tc.given[1]
		signature := sign(content, key)
		assert.Equal(t, tc.want, signature)
	}
}

func TestValidateSignature(t *testing.T) {
	tests := []struct {
		given []string
		want  bool
	}{
		{given: []string{"logistics_interface%3D%7B%22name%22%3A%22anson%22%2C%22age%22%3A18%7D%26data_digest%3DznNaBxySyki5gNhjSzsJsA%3D%3D", "key123"}, want: true},
		{given: []string{"logistics_interface%3D%7B%22name%22%3A%22anson%22%2C%22age%22%3A18%7D%26data_digest%3DFAKE_znNaBxySyki5gNhjSzsJsA%3D%3D", "key123"}, want: false},
	}

	for _, tc := range tests {
		content := tc.given[0]
		key := tc.given[1]
		signature := ValidateSignature(content, key)
		assert.Equal(t, tc.want, signature)
	}
}
