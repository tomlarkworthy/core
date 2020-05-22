// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: vdl

package vdl

import (
	"fmt"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// WireRetryCode is the suggested retry behavior for the receiver of an error.
// If the receiver doesn't know how to handle the specific error, it should
// attempt the suggested retry behavior.
type WireRetryCode int

const (
	WireRetryCodeNoRetry WireRetryCode = iota
	WireRetryCodeRetryConnection
	WireRetryCodeRetryRefetch
	WireRetryCodeRetryBackoff
)

// WireRetryCodeAll holds all labels for WireRetryCode.
var WireRetryCodeAll = [...]WireRetryCode{WireRetryCodeNoRetry, WireRetryCodeRetryConnection, WireRetryCodeRetryRefetch, WireRetryCodeRetryBackoff}

// WireRetryCodeFromString creates a WireRetryCode from a string label.
//nolint:deadcode,unused
func WireRetryCodeFromString(label string) (x WireRetryCode, err error) {
	err = x.Set(label)
	return
}

// Set assigns label to x.
func (x *WireRetryCode) Set(label string) error {
	switch label {
	case "NoRetry", "noretry":
		*x = WireRetryCodeNoRetry
		return nil
	case "RetryConnection", "retryconnection":
		*x = WireRetryCodeRetryConnection
		return nil
	case "RetryRefetch", "retryrefetch":
		*x = WireRetryCodeRetryRefetch
		return nil
	case "RetryBackoff", "retrybackoff":
		*x = WireRetryCodeRetryBackoff
		return nil
	}
	*x = -1
	return fmt.Errorf("unknown label %q in vdl.WireRetryCode", label)
}

// String returns the string label of x.
func (x WireRetryCode) String() string {
	switch x {
	case WireRetryCodeNoRetry:
		return "NoRetry"
	case WireRetryCodeRetryConnection:
		return "RetryConnection"
	case WireRetryCodeRetryRefetch:
		return "RetryRefetch"
	case WireRetryCodeRetryBackoff:
		return "RetryBackoff"
	}
	return ""
}

func (WireRetryCode) VDLReflect(struct {
	Name string `vdl:"v.io/v23/vdl.WireRetryCode"`
	Enum struct{ NoRetry, RetryConnection, RetryRefetch, RetryBackoff string }
}) {
}

func (x WireRetryCode) VDLIsZero() bool {
	return x == WireRetryCodeNoRetry
}

func (x WireRetryCode) VDLWrite(enc Encoder) error {
	if err := enc.WriteValueString(__VDLType_enum_1, x.String()); err != nil {
		return err
	}
	return nil
}

func (x *WireRetryCode) VDLRead(dec Decoder) error {
	switch value, err := dec.ReadValueString(); {
	case err != nil:
		return err
	default:
		if err := x.Set(value); err != nil {
			return err
		}
	}
	return nil
}

// WireError is the wire representation for the built-in error type.  Errors and
// exceptions in each programming environment are converted to this type to
// ensure wire compatibility.  Generated code for each environment provides
// automatic conversions into idiomatic native representations.
type WireError struct {
	Id        string        // Error Id, used to uniquely identify each error.
	RetryCode WireRetryCode // Retry behavior suggested for the receiver.
	Msg       string        // Error message, may be empty.
	ParamList []*Value      // Variadic parameters contained in the error.
}

func (WireError) VDLReflect(struct {
	Name string `vdl:"v.io/v23/vdl.WireError"`
}) {
}

func (x WireError) VDLIsZero() bool {
	if x.Id != "" {
		return false
	}
	if x.RetryCode != WireRetryCodeNoRetry {
		return false
	}
	if x.Msg != "" {
		return false
	}
	if len(x.ParamList) != 0 {
		return false
	}
	return true
}

func (x WireError) VDLWrite(enc Encoder) error {
	if err := enc.StartValue(__VDLType_struct_2); err != nil {
		return err
	}
	if x.Id != "" {
		if err := enc.NextFieldValueString(0, StringType, x.Id); err != nil {
			return err
		}
	}
	if x.RetryCode != WireRetryCodeNoRetry {
		if err := enc.NextFieldValueString(1, __VDLType_enum_1, x.RetryCode.String()); err != nil {
			return err
		}
	}
	if x.Msg != "" {
		if err := enc.NextFieldValueString(2, StringType, x.Msg); err != nil {
			return err
		}
	}
	if len(x.ParamList) != 0 {
		if err := enc.NextField(3); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_1(enc, x.ParamList); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_1(enc Encoder, x []*Value) error {
	if err := enc.StartValue(__VDLType_list_3); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntry(false); err != nil {
			return err
		}
		if elem == nil {
			if err := enc.NilValue(AnyType); err != nil {
				return err
			}
		} else {
			if err := elem.VDLWrite(enc); err != nil {
				return err
			}
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *WireError) VDLRead(dec Decoder) error {
	*x = WireError{}
	if err := dec.StartValue(__VDLType_struct_2); err != nil {
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
		if decType != __VDLType_struct_2 {
			index = __VDLType_struct_2.FieldIndexByName(decType.Field(index).Name)
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
				x.Id = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				if err := x.RetryCode.Set(value); err != nil {
					return err
				}
			}
		case 2:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Msg = value
			}
		case 3:
			if err := __VDLReadAnon_list_1(dec, &x.ParamList); err != nil {
				return err
			}
		}
	}
}

func __VDLReadAnon_list_1(dec Decoder, x *[]*Value) error {
	if err := dec.StartValue(__VDLType_list_3); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]*Value, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, err := dec.NextEntry(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			var elem *Value
			elem = new(Value)
			if err := elem.VDLRead(dec); err != nil {
				return err
			}
			*x = append(*x, elem)
		}
	}
}

// Type-check native conversion functions.
var ()

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	__VDLType_enum_1   *Type
	__VDLType_struct_2 *Type
	__VDLType_list_3   *Type
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

	// Register native type conversions first, so that vdl.TypeOf works.

	// Register types.
	Register((*WireRetryCode)(nil))
	Register((*WireError)(nil))

	// Initialize type definitions.
	__VDLType_enum_1 = TypeOf((*WireRetryCode)(nil))
	__VDLType_struct_2 = TypeOf((*WireError)(nil)).Elem()
	__VDLType_list_3 = TypeOf((*[]*Value)(nil))

	return struct{}{}
}
