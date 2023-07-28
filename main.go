package main

import (
	"fmt"

	"github.com/coscup_2023/code_gen/gen/coscuppb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func runCodeGen() {
	speaker := coscuppb.Speaker{Id: 2023, Name: "Rain"}
	fmt.Printf("speaker_id: %d, speaker_name: %s", speaker.Id, speaker.Name)
}

func truncatedField(v protoreflect.Value, limit int) protoreflect.Value {
	vStr := v.String()
	if limit <= 3 || len(vStr) <= limit {
		return v
	}
	vStr = fmt.Sprintf("%s...", vStr[:limit-3])
	return protoreflect.ValueOfString(vStr)
}

func runCustomOption(s *coscuppb.Session) {
	s.ProtoReflect().Range(
		func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
			opts := fd.Options().(*descriptorpb.FieldOptions)
			limit := proto.GetExtension(opts, coscuppb.E_LenLimit)
			s.ProtoReflect().Set(
				fd,
				// truncate the field if it exceeds the len limit
				truncatedField(v, int(limit.(int32))),
			)
			return true
		},
	)
}

func main() {
	// runCodeGen()
	speaker := coscuppb.Speaker{Id: 2023, Name: "Rain"}
	session := coscuppb.Session{Speaker: &speaker, Intro: "Protocol Buffer v3 in Go"}
	runCustomOption(&session)
}
