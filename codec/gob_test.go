package codec_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tuanliang/rpc/codec"
)

type TestStruct struct {
	F1 string
	F2 int
}

func TestGob(t *testing.T) {
	should := assert.New(t)
	gobbytes, err := codec.GobEncode(&TestStruct{F1: "test_f1", F2: 11})
	if should.NoError(err) {
		fmt.Println(gobbytes)
	}

	obj := TestStruct{}
	err = codec.GobDecode(gobbytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj)
	}
}
