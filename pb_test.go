package pbx

import (
	"testing"

	"github.com/libgo/pbx/internal"
)

func TestMarshal(t *testing.T) {
	var tp *internal.TestProto
	data, err := Marshal(tp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))

	tp = &internal.TestProto{
		Id:     "uuid",
		Name:   "elvizlai",
		Age:    18,
		NextId: "next-uuid",
		Filter: map[string]string{
			"k1": "v1",
		},
	}
	data, err = Marshal(tp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

const dataStr = `{"id":"uuid","name":"elvizlai","age":"18","nextId":"next-uuid","filter":{"k1":"v1"}}`

func TestUnmarshal(t *testing.T) {
	tp := &internal.TestProto{}

	err := Unmarshal([]byte(dataStr), tp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tp)

	err = Unmarshal([]byte(dataStr), tp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tp)

	tp.Reset()

	err = Unmarshal([]byte("null"), tp)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tp == nil)
}

func TestMarshalAny(t *testing.T) {
	list := make([]Message, 3)
	list[1] = &internal.TestProto{}

	as := MarshalAny(list[1])
	t.Log(as, len(as))

	as = MarshalAny(list)
	t.Log(as, len(as))
}

func TestUnmarshalAny(t *testing.T) {
	pb1 := &internal.TestProto{}
	a1 := MarshalAny(pb1)
	err := UnmarshalAny(a1[0], pb1)
	if err != nil {
		t.Fatal(err)
	}

	a2 := MarshalAny([]*internal.TestProto{
		{
			Name: "n1",
		}, {
			Name: "n2",
		},
	})
	pb2 := make([]*internal.TestProto, 2)
	err = UnmarshalAny(a2, pb2)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(a2)
}
