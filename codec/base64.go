package codec

import (
	"encoding/base64"
)

type Base64Codec struct{}

func (c *Base64Codec) Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (c *Base64Codec) Decode(s string) (string, error) {
	if v, err := base64.StdEncoding.DecodeString(s); err != nil {
		return "", err
	} else {
		return string(v), nil
	}
}
