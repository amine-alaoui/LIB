// Copyright 2012 Aaron Jacobs. All Rights Reserved.
// Author: aaronjjacobs@gmail.com (Aaron Jacobs)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package siv

import (
	"strconv"
	"testing"

	. "github.com/jacobsa/oglematchers"
	. "github.com/jacobsa/ogletest"
)

func TestXorend(t *testing.T) { RunTests(t) }

////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////

// Like xorend, but doesn't require an existing buffer.
func easyXorend(a, b []byte) []byte {
	dst := make([]byte, len(a))
	xorend(dst, a, b)
	return dst
}

func fromBinary(s string) byte {
	AssertEq(8, len(s), "%s", s)

	u, err := strconv.ParseUint(s, 2, 8)
	AssertEq(nil, err, "%s", s)

	return byte(u)
}

type XorendTest struct{}

func init() { RegisterTestSuite(&XorendTest{}) }

////////////////////////////////////////////////////////////////////////
// Tests
////////////////////////////////////////////////////////////////////////

func (t *XorendTest) AIsShorterThanB() {
	a := []byte{0xde}
	b := []byte{0xde, 0xad}

	f := func() { easyXorend(a, b) }
	ExpectThat(f, Panics(HasSubstr("length")))
}

func (t *XorendTest) BothAreNil() {
	a := []byte(nil)
	b := []byte(nil)

	expected := []byte{}
	ExpectThat(easyXorend(a, b), DeepEquals(expected))
}

func (t *XorendTest) BIsNil() {
	a := []byte{fromBinary("10101010"), fromBinary("00000000")}
	b := []byte(nil)

	expected := a
	ExpectThat(easyXorend(a, b), DeepEquals(expected))
}

func (t *XorendTest) BothAreEmpty() {
	a := []byte{}
	b := []byte{}

	expected := []byte{}
	ExpectThat(easyXorend(a, b), DeepEquals(expected))
}

func (t *XorendTest) BIsEmpty() {
	a := []byte{fromBinary("10101010"), fromBinary("00000000")}
	b := []byte{}

	expected := a
	ExpectThat(easyXorend(a, b), DeepEquals(expected))
}

func (t *XorendTest) BIsNonEmpty() {
	a := []byte{
		fromBinary("11110000"),
		fromBinary("10101010"),
		fromBinary("00000000"),
	}

	b := []byte{
		fromBinary("11110000"),
		fromBinary("00001111"),
	}

	expected := []byte{
		fromBinary("11110000"),
		fromBinary("01011010"),
		fromBinary("00001111"),
	}

	ExpectThat(easyXorend(a, b), DeepEquals(expected))
}

func (t *XorendTest) DoesntClobberInputData() {
	a := []byte{
		fromBinary("11110000"),
		fromBinary("10101010"),
		fromBinary("00000000"),
	}
	aCopy := dup(a)

	b := []byte{
		fromBinary("11110000"),
		fromBinary("00001111"),
	}
	bCopy := dup(b)

	easyXorend(a, b)
	ExpectThat(a, DeepEquals(aCopy))
	ExpectThat(b, DeepEquals(bCopy))
}
