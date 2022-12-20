package codec

import "testing"

func TestJsonCodecWrite(t *testing.T) {
	codecFunc := NewCodecFuncMap[JsonType]
	codec := codecFunc(nil)
	codec.Write(nil, &User{"zhangsan", "123456"})
}
