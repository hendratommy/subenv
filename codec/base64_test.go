package codec

import "testing"

var (
	in  = "hello world"
	out = "aGVsbG8gd29ybGQ="

	b64 = &Base64Codec{}
)

func TestBase64Codec_Encode(t *testing.T) {
	res := b64.Encode(in)
	if res != out {
		t.Errorf("expect %s equals %s", res, out)
	}
}

func TestBase64Codec_Decode(t *testing.T) {
	if res, err := b64.Decode(b64.Encode(in)); err != nil {
		t.Errorf("decode should not return error: %v", err)
	} else {
		if res != in {
			t.Errorf("expect %s equals %s", res, in)
		}
	}

	if res, err := b64.Decode(in); err == nil {
		t.Errorf("decode should return error, got %s", res)
	}
}
