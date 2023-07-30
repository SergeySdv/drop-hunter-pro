package task

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Marshal(m proto.Message) ([]byte, error) {
	return protojson.MarshalOptions{Multiline: true, UseEnumNumbers: false, UseProtoNames: false, EmitUnpopulated: true}.Marshal(m)
}
