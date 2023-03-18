package pb_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tuanliang/rpc/protobuf/pb"
	"google.golang.org/protobuf/proto"
)

func TestMarshal(t *testing.T) {
	should := assert.New(t)
	str := &pb.String{Value: "hello"}

	// object--protobuf--[]byte
	pbBytes, err := proto.Marshal(str)
	if should.NoError(err) {
		fmt.Println(pbBytes)
	}

	// []byte--protobuf--object
	obj := pb.String{}
	err = proto.Unmarshal(pbBytes, &obj)
	if should.NoError(err) {
		fmt.Println(obj.Value)
	}
}
