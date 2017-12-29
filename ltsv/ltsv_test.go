package ltsv

import (
	"testing"
)

type User struct {
	Name      string `ltsv:"name"`
	Age       int    `ltsv:"age"`
	BirthYear int    `ltsv:"year"`
}

func TestMarshal(t *testing.T) {
	user := User{
		Name:      "naoto",
		Age:       27,
		BirthYear: 1990,
	}
	ret := Marshal(user)
	want := "name:naoto\tage:27\tyear:1990"
	if ret != want {
		t.Error("Error TestMarshal")
	}
}

func TestPairArray(t *testing.T) {
	pa := pairArray{}
	pa.append("hoge", "foo")
	pa.append("hogehoge", "foofoo")
	if pa.len() != 2 {
		t.Error("Error TestPairArray len()")
	}

	ret := pa.join()
	want := "hoge:foo\thogehoge:foofoo"
	if ret != want {
		t.Error("Error TestPairArray join()")
	}
}

func TestPair(t *testing.T) {
	p := pair{key: "hoge", value: "foo"}
	ret := p.join()
	want := "hoge:foo"
	if ret != want {
		t.Error("Error TestPair join()")
	}
}
