// Protocol Buffers for Go with Gadgets
//
// Copyright (c) 2015, The GoGo Authors. All rights reserved.
// http://github.com/gogo/protobuf
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//     * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package fuzztests

import (
	"github.com/yangshengBE/protobuf/proto"
	"testing"
)

func TestFuzzUnrecognized(t *testing.T) {
	msg := &Nil{}
	input := []byte{0x8, 0xaf, 0x81, 0xc9, 0xb3, 0x97, 0xd1, 0xb5, 0xc2, 0x4f, 0x1a, 0x4a, 0x52, 0x48, 0x4e, 0x44, 0x65, 0x51, 0x4b, 0x46, 0x44, 0x33, 0x5a, 0x44, 0x72, 0x38, 0x58, 0x4c, 0x58, 0x70, 0x59, 0x45, 0x71, 0x45, 0x4f, 0x6d, 0x45, 0x4d, 0x54, 0x59, 0x4c, 0x6b, 0x55, 0x7a, 0x6f, 0x6b, 0x5a, 0x69, 0x56, 0x64, 0x46, 0x45, 0x56, 0x4d, 0x70, 0x6a, 0x39, 0x7a, 0x4b, 0x43, 0x4d, 0x6d, 0x76, 0x63, 0x46, 0x4f, 0x31, 0x4a, 0x5a, 0x6b, 0x66, 0x4a, 0x75, 0x51, 0x38, 0x54, 0x54, 0x30, 0x53, 0x61, 0x36, 0x6e, 0x4f, 0x6b, 0x35, 0x54, 0x95, 0x0, 0x0, 0x0, 0x0, 0x12, 0x38, 0x52, 0x36, 0x66, 0x76, 0x41, 0x74, 0x73, 0x7a, 0x39, 0x43, 0x6a, 0x4f, 0x64, 0x59, 0x77, 0x33, 0x30, 0x36, 0x58, 0x75, 0x65, 0x46, 0x4b, 0x46, 0x55, 0x56, 0x71, 0x6d, 0x49, 0x73, 0x4a, 0x4b, 0x78, 0x76, 0x41, 0x65, 0x42, 0x61, 0x5a, 0x30, 0x41, 0x37, 0x45, 0x76, 0x72, 0x31, 0x30, 0x4e, 0x78, 0x6d, 0x33, 0x63, 0x65, 0x66, 0x6b, 0x30}
	if err := proto.Unmarshal(input, msg); err == nil {
		t.Fatal("expected error")
	}
}

func DisabledTestFuzzPackedIsNotIdempotent(t *testing.T) {
	msg := &NinRepPackedNative{}
	//original := []byte{0x9, 0xa3, 0xae, 0xab, 0xd2, 0xbe, 0x1c, 0xed, 0xbf, 0x15, 0x22, 0x1, 0x6e, 0x3f, 0x22, 0x81, 0x1, 0x9, 0x21, 0x84, 0x36, 0x21, 0x6a, 0xff, 0xd3, 0x3f, 0x15, 0x15, 0x71, 0x4b, 0xbd, 0x52, 0x70, 0x49, 0x6e, 0x48, 0x54, 0x6a, 0x61, 0x37, 0x63, 0x78, 0x47, 0x58, 0x31, 0x7a, 0x43, 0x4d, 0x7a, 0x48, 0x58, 0x56, 0xcb, 0x9c, 0x34, 0xdf, 0xc6, 0x3c, 0xa4, 0x33, 0xac, 0xba, 0xa7, 0xeb, 0x4, 0xa8, 0x8a, 0x48, 0x75, 0x67, 0x71, 0x31, 0x4a, 0x5b, 0xe1, 0xcf, 0x21, 0x88, 0xd3, 0xec, 0xac, 0x13, 0x28, 0xec, 0xa9, 0x51, 0xc8, 0xe9, 0x5e, 0xca, 0xbe, 0xea, 0x9c, 0x0, 0x6b, 0x44, 0x63, 0xc4, 0x32, 0xaa, 0x36, 0x4e, 0xfc, 0xbd, 0x7, 0xef, 0x5e, 0x47, 0x2, 0xfc, 0xd8, 0x83, 0x85, 0x9c, 0xca, 0x7c, 0xd2, 0xdb, 0xf5, 0x5d, 0xcc, 0x5a, 0x72, 0x1f, 0x66, 0x55, 0x74, 0x46, 0x47, 0x73, 0x75, 0x54, 0x30, 0x39, 0x53, 0x34, 0x4c, 0x61, 0x78, 0x59, 0x31, 0x51, 0x44, 0x30, 0x53, 0x51, 0x71, 0x44, 0x65, 0x6f, 0x53, 0x30, 0x44, 0x6a, 0x58, 0x7a, 0x1b, 0x65, 0x62, 0x9c, 0x95, 0xc5, 0x41, 0xcb, 0x48, 0xa, 0x47, 0xf6, 0xd8, 0xd2, 0xd5, 0x8d, 0x6, 0x69, 0x8f, 0xbe, 0x7c, 0xf3, 0xe9, 0x79, 0x3c, 0xca, 0x6, 0x5b}
	input := []byte{0x9, 0xa3, 0xae, 0xab, 0xd2, 0xbe, 0x1c, 0xed, 0xbf, 0x15, 0x22, 0x1, 0x6e, 0x3f, 0x22, 0x81, 0x1, 0x9, 0x21, 0x84, 0x36, 0x21, 0x6a, 0xff, 0xd3, 0x3f, 0x15, 0x15, 0x71, 0x4b, 0xbd, 0x52, 0x70, 0x49, 0x6e, 0x48, 0x54, 0x6a, 0x61, 0x37, 0x63, 0x78, 0x47, 0x58, 0x31, 0x7a, 0x43, 0x4d, 0x7a, 0x48, 0x58, 0x56, 0xcb, 0x9c, 0x34, 0xdf, 0xc6, 0x3c, 0xa4, 0x33, 0xac, 0xba, 0xa7, 0xeb, 0x4, 0xa8, 0x8a, 0x48, 0x75, 0x67, 0x71, 0x31, 0x4a, 0x5b, 0xe1, 0xcf, 0x21, 0x88, 0xd3, 0xec, 0xac, 0x13, 0x28, 0xec, 0xa9, 0x51, 0xc8, 0xe9, 0x5e, 0xca, 0xbe, 0xea, 0x9c, 0x0, 0x6b, 0x44, 0x63, 0xc4, 0x32, 0xaa, 0x36, 0x4e, 0xfc, 0xbd, 0x7, 0xef, 0x5e, 0x47, 0x2, 0xfc, 0xd8, 0x83, 0x85, 0x9c, 0xca, 0x7c, 0xd2, 0xdb, 0xf5, 0x5d, 0xcc, 0x5a, 0x72, 0x1f, 0x66, 0x55, 0x74, 0x46, 0x47, 0x73, 0x75, 0x54, 0x30, 0x39, 0x53, 0x34, 0x4c, 0x61, 0x78, 0x59, 0x31, 0x51, 0x44, 0x30, 0x53, 0x51}
	if err := proto.Unmarshal(input, msg); err == nil {
		t.Fatal("expected error")
	}
}

func DisabledTestFuzzFieldOrder(t *testing.T) {
	msg := &NinOptStruct{}
	input := []byte{0x52, 0x57, 0x52, 0x6a, 0x33, 0x56, 0x43, 0x76, 0x32, 0x54, 0x49, 0x4a, 0x55, 0x66, 0x39, 0x52, 0x32, 0x32, 0x73, 0x69, 0x4f, 0x67, 0x66, 0x79, 0x4b, 0x79, 0x5a, 0x55, 0x42, 0x53, 0x38, 0x68, 0x6c, 0x46, 0x79, 0x6b, 0x54, 0x43, 0x63, 0x66, 0x30, 0x6a, 0x33, 0x35, 0x33, 0x7a, 0x41, 0x66, 0x68, 0x57, 0x61, 0x78, 0x51, 0x37, 0x76, 0x52, 0x78, 0x34, 0x56, 0x43, 0x54, 0x31, 0x73, 0x6a, 0x77, 0x63, 0x45, 0x62, 0x62, 0x67, 0x34, 0x6f, 0x64, 0x35, 0x6c, 0x41, 0x45, 0x50, 0x64, 0x6f, 0x46, 0x38, 0x41, 0x4b, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30}
	if err := proto.Unmarshal(input, msg); err != nil {
		t.Fatal(err)
	}
}

func TestFuzzSint64Overflow(t *testing.T) {
	msg := &NinOptNative{}
	//original := []byte{0x9, 0x65, 0xb4, 0xfd, 0xbc, 0x5, 0xc7, 0xee, 0x3f, 0x15, 0x48, 0xec, 0x67, 0x3f, 0x18, 0xca, 0xa4, 0xe0, 0xa9, 0x5, 0x20, 0x8e, 0xb7, 0x9f, 0xf5, 0xcf, 0xe9, 0xea, 0xad, 0xfd, 0x1, 0x28, 0xc9, 0xf1, 0xbc, 0x88, 0xc, 0x30, 0xeb, 0x99, 0xbd, 0xa8, 0xe, 0x38, 0xc0, 0xd4, 0xb7, 0xba, 0x7, 0x40, 0xc8, 0xe4, 0xf6, 0xe2, 0xb8, 0xdd, 0xa7, 0xf2, 0x82, 0xba, 0x16, 0x9d, 0x59, 0xf9, 0x31, 0xe0, 0x99, 0x0, 0x0, 0x0, 0x0, 0x61, 0x59, 0x5b, 0xb5, 0x57, 0x56, 0x93, 0x70, 0xde, 0x68, 0x0, 0x72, 0x40, 0x64, 0x5a, 0x5a, 0x61, 0x57, 0x78, 0x68, 0x53, 0x65, 0x66, 0x67, 0x38, 0x38, 0x61, 0x48, 0x44, 0x32, 0x6c, 0x36, 0x50, 0x31, 0x4d, 0x43, 0x39, 0x31, 0x6d, 0x37, 0x34, 0x32, 0x48, 0x6b, 0x4d, 0x70, 0x31, 0x45, 0x73, 0x48, 0x71, 0x4a, 0x69, 0x37, 0x56, 0x53, 0x44, 0x6b, 0x48, 0x45, 0x50, 0x4b, 0x7a, 0x52, 0x49, 0x4c, 0x50, 0x69, 0x44, 0x72, 0x42, 0x56, 0x50, 0x78, 0x62, 0x56, 0x55, 0x7a, 0x5b, 0xb3, 0x6c, 0x59, 0x4c, 0xf1, 0x31, 0xeb, 0xb6, 0x25, 0x1a, 0x26, 0x67, 0x66, 0x97, 0x79, 0xb8, 0x37, 0x8, 0xe1, 0x32, 0x45, 0x6e, 0x6, 0x90, 0x4f, 0xde, 0x26, 0x7a, 0xc6, 0x29, 0x65, 0x4a, 0x69, 0xa7, 0x21, 0xfb, 0x42, 0xda, 0x43, 0x89, 0x27, 0x70, 0x71, 0xde, 0x66, 0xa4, 0x75, 0x2b, 0x5c, 0x96, 0x9f, 0x25, 0x3b, 0xc1, 0x64, 0x14, 0x4, 0x60, 0x8c, 0x58, 0x7e, 0xa1, 0x59, 0x7b, 0x47, 0x18, 0xc, 0x5b, 0x18, 0x63, 0x9, 0xb4, 0xc9, 0x7, 0xf9, 0xae, 0x33, 0xae, 0x2, 0x4a, 0x8b, 0x34, 0x92, 0x40, 0xb, 0xd7, 0x80, 0x60, 0xdb, 0x44, 0x5}
	input := []byte{0x40, 0xc8, 0xe4, 0xf6, 0xe2, 0xb8, 0xdd, 0xa7, 0xf2, 0x82, 0xba, 0x16}
	if err := proto.Unmarshal(input, msg); err != nil {
		return
	}
}

func DisabledTestFuzzOverrideField(t *testing.T) {
	msg := &NinOptNative{}
	//original := []byte{0x9, 0x73, 0x78, 0x5a, 0xf2, 0xb4, 0x66, 0xe8, 0x3f, 0x15, 0x71, 0xdc, 0x4, 0x3f, 0x18, 0xe5, 0x8e, 0xab, 0xdb, 0x3, 0x20, 0xbe, 0xed, 0xe6, 0xc0, 0xb9, 0xb8, 0xa7, 0xb5, 0x12, 0x28, 0xcb, 0x8c, 0x91, 0xef, 0xc, 0x30, 0x9a, 0xc1, 0xc3, 0xc0, 0xf, 0x38, 0xe8, 0x9b, 0xf0, 0xca, 0x5, 0x40, 0xd2, 0xd7, 0xdd, 0xa3, 0xea, 0xab, 0xec, 0xc2, 0xaa, 0x1, 0x4d, 0xc9, 0x15, 0x0, 0xea, 0x55, 0x72, 0x3e, 0x92, 0xa8, 0x59, 0x3e, 0x87, 0x7d, 0xf5, 0x0, 0x0, 0x0, 0x0, 0x61, 0xca, 0xe7, 0xdb, 0x57, 0xa1, 0xb6, 0x41, 0xf4, 0x72, 0x3e, 0x63, 0x36, 0x43, 0x73, 0x32, 0x68, 0x64, 0x75, 0x75, 0x70, 0x4f, 0x39, 0x67, 0x77, 0x42, 0x6a, 0x78, 0x63, 0x57, 0x64, 0x77, 0x6b, 0x74, 0x44, 0x4d, 0x79, 0x36, 0x30, 0x68, 0x38, 0x53, 0x31, 0x79, 0x33, 0x38, 0x4b, 0x7a, 0x76, 0x36, 0x48, 0x4a, 0x35, 0x37, 0x59, 0x48, 0x6c, 0x74, 0x72, 0x61, 0x33, 0x4c, 0x74, 0x45, 0x4a, 0x51, 0x68, 0x71, 0x31, 0x70, 0x50, 0x70, 0x6a, 0x4d, 0x15, 0x51, 0xce, 0xea, 0x82, 0x1, 0x23, 0xed, 0x7a, 0x3, 0x78, 0xee, 0x56, 0x46, 0xd0, 0xe1, 0x17, 0x18, 0x30, 0x9d, 0x2f, 0xac, 0x1c, 0xa, 0x30, 0xa9, 0x8d, 0x10, 0xed, 0xb5, 0x44, 0x36, 0x5e, 0x84, 0x73, 0x5d, 0x38, 0x51, 0x2b, 0x6e, 0xc6, 0xb5}
	input := []byte{0x4d, 0xc9, 0x15, 0x0, 0xea, 0x72, 0x3e, 0x63, 0x36, 0x43, 0x73, 0x32, 0x68, 0x64, 0x75, 0x75, 0x70, 0x4f, 0x39, 0x67, 0x77, 0x42, 0x6a, 0x78, 0x63, 0x57, 0x64, 0x77, 0x6b, 0x74, 0x44, 0x4d, 0x79, 0x36, 0x30, 0x68, 0x38, 0x53, 0x31, 0x79, 0x33, 0x38, 0x4b, 0x7a, 0x76, 0x36, 0x48, 0x4a, 0x35, 0x37, 0x59, 0x48, 0x6c, 0x74, 0x72, 0x61, 0x33, 0x4c, 0x74, 0x45, 0x4a, 0x51, 0x68, 0x71, 0x31, 0x70, 0x50, 0x70, 0x6a, 0x4d, 0x15, 0x51, 0xce, 0xea}
	if err := proto.Unmarshal(input, msg); err != nil {
		panic(err)
	}
	output, err := proto.Marshal(msg)
	if err != nil {
		t.Fatal(err)
	}
	if len(input) != len(output) {
		t.Logf("%#v", msg)
		msg2 := &NinOptNative{}
		if err := proto.Unmarshal(output, msg2); err == nil {
			t.Logf("%#v", msg2)
		}
		t.Errorf("expected %#v got %#v", input, output)
	}
}

//Generated code is correct, non generated returns an incorrect error
func DisabledTestFuzzBadWireType(t *testing.T) {
	msg := &NinRepPackedNative{}
	//input := []byte("j\x160\xfc0000\xf6\xfa000\xc1\xaf\xf5000\xcf" + "00\xb90z\r0\x850\xd30000'0000")
	input := []byte{0x6a, 0x16, 0x30, 0xfc, 0x30, 0x30, 0x30, 0x30, 0xf6, 0xfa, 0x30, 0x30, 0x30, 0xc1, 0xaf, 0xf5, 0x30, 0x30, 0x30, 0xcf, 0x30, 0x30, 0xb9, 0x30, 0x7a, 0xd, 0x30, 0x85, 0x30, 0xd3, 0x30, 0x30, 0x30, 0x30, 0x27, 0x30, 0x30, 0x30, 0x30}
	if err := proto.Unmarshal(input, msg); err == nil {
		t.Fatalf("expected bad wiretype for Field4 error got %#v", msg)
	} else {
		t.Log(err)
	}
}

func TestFuzzIntegerOverflow(t *testing.T) {
	msg := &Nil{}
	//input := []byte("\x1500000\x8b\x9b\xa3\xa8\xb6\xe1\xe1\xfe\u061c0")
	input := []byte{0x15, 0x30, 0x30, 0x30, 0x30, 0x30, 0x8b, 0x9b, 0xa3, 0xa8, 0xb6, 0xe1, 0xe1, 0xfe, 0xd8, 0x9c, 0x30}
	if err := proto.Unmarshal(input, msg); err == nil {
		t.Fatalf("expected integer overflow error %#v", msg)
	} else {
		t.Log(err)
	}
}

//Generated code is correct, non generated returns an incorrect error
func DisabledTestFuzzUnexpectedEOF(t *testing.T) {
	msg := &NinRepPackedNative{}
	//input := []byte("j\x16000000000000000000" + "00\xb90")
	input := []byte{0x6a, 0x16, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0x30, 0xb9, 0x30}
	if err := proto.Unmarshal(input, msg); err == nil {
		t.Fatalf("expected unexpected eof error got %#v", msg)
	} else {
		t.Log(err)
	}
}

//Generated code is correct, non generated returns an incorrect error
func DisabledTestFuzzCantSkipWireType(t *testing.T) {
	msg := &NinRepPackedNative{}
	//input := []byte("j\x160\xfc0000\xf6\xfa000\xc1\xaf\xf5000\xcf" + "00\xb90z\r0\x850\xd3000\xa80\xa7000")
	input := []byte{0x6a, 0x16, 0x30, 0xfc, 0x30, 0x30, 0x30, 0x30, 0xf6, 0xfa, 0x30, 0x30, 0x30, 0xc1, 0xaf, 0xf5, 0x30, 0x30, 0x30, 0xcf, 0x30, 0x30, 0xb9, 0x30, 0x7a, 0xd, 0x30, 0x85, 0x30, 0xd3, 0x30, 0x30, 0x30, 0xa8, 0x30, 0xa7, 0x30, 0x30, 0x30}
	if err := proto.Unmarshal(input, msg); err == nil {
		t.Fatalf("expected cant skip wiretype error got %#v", msg)
	} else {
		t.Log(err)
	}
}
