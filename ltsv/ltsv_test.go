/*
MIT License

Copyright (c) 2017 naoto yamaguchi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

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
