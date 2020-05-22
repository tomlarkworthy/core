// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: rpc

package rpc

import (
	"v.io/v23/security"
	"v.io/v23/vdl"
	vdltime "v.io/v23/vdlroot/time"
	"v.io/v23/verror"
	"v.io/v23/vtrace"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Request describes the request header sent by the client to the server.  A
// non-zero request header is sent at the beginning of the RPC call, followed by
// the positional args.  Thereafter a zero request header is sent before each
// streaming arg, terminated by a non-zero request header with EndStreamArgs set
// to true.
type Request struct {
	// Suffix of the name used to identify the object hosting the service.
	Suffix string
	// Method to invoke on the service.
	Method string
	// NumPosArgs is the number of positional arguments, which follow this message
	// (and any blessings) on the request stream.
	NumPosArgs uint64
	// EndStreamArgs is true iff no more streaming arguments will be sent.  No
	// more data will be sent on the request stream.
	//
	// NOTE(bprosnitz): We can support multiple stream values per request (+response) header
	// efficiently by adding a NumExtraStreamArgs (+NumExtraStreamResults to response) field
	// that is the uint64 (number of stream args to send) - 1. The request is then zero when
	// exactly one streaming arg is sent. Since the request and response headers are small,
	// this is only likely necessary for frequently streaming small values.
	// See implementation in CL: 3913
	EndStreamArgs bool
	// Deadline after which the request should be cancelled.  This is a hint to
	// the server, to avoid wasted work.
	Deadline vdltime.Deadline
	// GrantedBlessings are blessings bound to the principal running the server,
	// provided by the client.
	GrantedBlessings security.Blessings
	// TraceRequest maintains the vtrace context between clients and servers
	// and specifies additional parameters that control how tracing behaves.
	TraceRequest vtrace.Request
	// Language indicates the language of the instegator of the RPC.
	// By convention it should be an IETF language tag:
	// http://en.wikipedia.org/wiki/IETF_language_tag
	Language string
}

func (Request) VDLReflect(struct {
	Name string `vdl:"v.io/v23/rpc.Request"`
}) {
}

func (x Request) VDLIsZero() bool {
	if x.Suffix != "" {
		return false
	}
	if x.Method != "" {
		return false
	}
	if x.NumPosArgs != 0 {
		return false
	}
	if x.EndStreamArgs {
		return false
	}
	if !x.Deadline.Time.IsZero() {
		return false
	}
	if !x.GrantedBlessings.IsZero() {
		return false
	}
	if x.TraceRequest != (vtrace.Request{}) {
		return false
	}
	if x.Language != "" {
		return false
	}
	return true
}

func (x Request) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.Suffix != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Suffix); err != nil {
			return err
		}
	}
	if x.Method != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.Method); err != nil {
			return err
		}
	}
	if x.NumPosArgs != 0 {
		if err := enc.NextFieldValueUint(2, vdl.Uint64Type, x.NumPosArgs); err != nil {
			return err
		}
	}
	if x.EndStreamArgs {
		if err := enc.NextFieldValueBool(3, vdl.BoolType, x.EndStreamArgs); err != nil {
			return err
		}
	}
	if !x.Deadline.Time.IsZero() {
		if err := enc.NextField(4); err != nil {
			return err
		}
		var wire vdltime.WireDeadline
		if err := vdltime.WireDeadlineFromNative(&wire, x.Deadline); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if !x.GrantedBlessings.IsZero() {
		if err := enc.NextField(5); err != nil {
			return err
		}
		var wire security.WireBlessings
		if err := security.WireBlessingsFromNative(&wire, x.GrantedBlessings); err != nil {
			return err
		}
		if err := wire.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.TraceRequest != (vtrace.Request{}) {
		if err := enc.NextField(6); err != nil {
			return err
		}
		if err := x.TraceRequest.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Language != "" {
		if err := enc.NextFieldValueString(7, vdl.StringType, x.Language); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Request) VDLRead(dec vdl.Decoder) error {
	*x = Request{}
	if err := dec.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_1 {
			index = __VDLType_struct_1.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Suffix = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Method = value
			}
		case 2:
			switch value, err := dec.ReadValueUint(64); {
			case err != nil:
				return err
			default:
				x.NumPosArgs = value
			}
		case 3:
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.EndStreamArgs = value
			}
		case 4:
			var wire vdltime.WireDeadline
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := vdltime.WireDeadlineToNative(wire, &x.Deadline); err != nil {
				return err
			}
		case 5:
			var wire security.WireBlessings
			if err := wire.VDLRead(dec); err != nil {
				return err
			}
			if err := security.WireBlessingsToNative(wire, &x.GrantedBlessings); err != nil {
				return err
			}
		case 6:
			if err := x.TraceRequest.VDLRead(dec); err != nil {
				return err
			}
		case 7:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Language = value
			}
		}
	}
}

// Response describes the response header sent by the server to the client.  A
// zero response header is sent before each streaming arg.  Thereafter a
// non-zero response header is sent at the end of the RPC call, right before
// the positional results.
type Response struct {
	// Error in processing the RPC at the server. Implies EndStreamResults.
	Error error
	// EndStreamResults is true iff no more streaming results will be sent; the
	// remainder of the stream consists of NumPosResults positional results.
	EndStreamResults bool
	// NumPosResults is the number of positional results, which immediately follow
	// on the response stream.  After these results, no further data will be sent
	// on the response stream.
	NumPosResults uint64
	// TraceResponse maintains the vtrace context between clients and servers.
	// In some cases trace data will be included in this response as well.
	TraceResponse vtrace.Response
	// AckBlessings is true if the server successfully recevied the client's
	// blessings and stored them in the server's blessings cache.
	AckBlessings bool
}

func (Response) VDLReflect(struct {
	Name string `vdl:"v.io/v23/rpc.Response"`
}) {
}

func (x Response) VDLIsZero() bool {
	if x.Error != nil {
		return false
	}
	if x.EndStreamResults {
		return false
	}
	if x.NumPosResults != 0 {
		return false
	}
	if !x.TraceResponse.VDLIsZero() {
		return false
	}
	if x.AckBlessings {
		return false
	}
	return true
}

func (x Response) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_5); err != nil {
		return err
	}
	if x.Error != nil {
		if err := enc.NextField(0); err != nil {
			return err
		}
		if err := verror.VDLWrite(enc, x.Error); err != nil {
			return err
		}
	}
	if x.EndStreamResults {
		if err := enc.NextFieldValueBool(1, vdl.BoolType, x.EndStreamResults); err != nil {
			return err
		}
	}
	if x.NumPosResults != 0 {
		if err := enc.NextFieldValueUint(2, vdl.Uint64Type, x.NumPosResults); err != nil {
			return err
		}
	}
	if !x.TraceResponse.VDLIsZero() {
		if err := enc.NextField(3); err != nil {
			return err
		}
		if err := x.TraceResponse.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.AckBlessings {
		if err := enc.NextFieldValueBool(4, vdl.BoolType, x.AckBlessings); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Response) VDLRead(dec vdl.Decoder) error {
	*x = Response{}
	if err := dec.StartValue(__VDLType_struct_5); err != nil {
		return err
	}
	decType := dec.Type()
	for {
		index, err := dec.NextField()
		switch {
		case err != nil:
			return err
		case index == -1:
			return dec.FinishValue()
		}
		if decType != __VDLType_struct_5 {
			index = __VDLType_struct_5.FieldIndexByName(decType.Field(index).Name)
			if index == -1 {
				if err := dec.SkipValue(); err != nil {
					return err
				}
				continue
			}
		}
		switch index {
		case 0:
			if err := verror.VDLRead(dec, &x.Error); err != nil {
				return err
			}
		case 1:
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.EndStreamResults = value
			}
		case 2:
			switch value, err := dec.ReadValueUint(64); {
			case err != nil:
				return err
			default:
				x.NumPosResults = value
			}
		case 3:
			if err := x.TraceResponse.VDLRead(dec); err != nil {
				return err
			}
		case 4:
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.AckBlessings = value
			}
		}
	}
}

//////////////////////////////////////////////////
// Const definitions

// TODO(toddw): Rename GlobMethod to ReservedGlob.
const GlobMethod = "__Glob"
const ReservedSignature = "__Signature"
const ReservedMethodSignature = "__MethodSignature"

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_struct_2 *vdl.Type
	__VDLType_struct_3 *vdl.Type
	__VDLType_struct_4 *vdl.Type
	__VDLType_struct_5 *vdl.Type
	__VDLType_struct_6 *vdl.Type
)

var __VDLInitCalled bool

// __VDLInit performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = __VDLInit()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func __VDLInit() struct{} {
	if __VDLInitCalled {
		return struct{}{}
	}
	__VDLInitCalled = true

	// Register types.
	vdl.Register((*Request)(nil))
	vdl.Register((*Response)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*Request)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*vdltime.WireDeadline)(nil)).Elem()
	__VDLType_struct_3 = vdl.TypeOf((*security.WireBlessings)(nil)).Elem()
	__VDLType_struct_4 = vdl.TypeOf((*vtrace.Request)(nil)).Elem()
	__VDLType_struct_5 = vdl.TypeOf((*Response)(nil)).Elem()
	__VDLType_struct_6 = vdl.TypeOf((*vtrace.Response)(nil)).Elem()

	return struct{}{}
}
