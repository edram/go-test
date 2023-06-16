package test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"
	"testing"
)

type S struct {
	Name string
	Age  int
}

// 从 struct 编码成 buf
func TestEncodeGob(t *testing.T) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	s := S{Name: "name", Age: 18}

	enc.Encode(s)

	if err := enc.Encode(s); err != nil {
		t.Fatal(err)
	}

	print(buf.String())
}

// 从 buf 解码成 struct
func TestDecodeGob(t *testing.T) {
	str := " \xff\x81\x03\x01\x01\x01S\x01\xff\x82\x00\x01\x02\x01\x04Name\x01\f\x00\x01\x03Age\x01\x04\x00\x00\x00\v\xff\x82\x01\x04name\x01$\x00\v\xff\x82\x01\x04name\x01$\x00"

	buf := bytes.NewBuffer([]byte(str))
	dec := gob.NewDecoder(buf)

	var s S

	var check = S{Name: "name", Age: 18}

	if err := dec.Decode(&s); err != nil {
		t.Fatalf("解码失败 %v", err)
	}

	if !reflect.DeepEqual(s, check) {
		t.Fatalf("解码出来的%#v和检查的%#v不一致", s, check)
	}

	fmt.Printf("%#v", s)
}
