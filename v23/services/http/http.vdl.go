// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: http

// Package HTTP defines an interface to send a http.Request from a client to a
// Vanadium server. This code is Go-specific since it is only used internally
// by Vanadium.
package http

import (
	"v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/vdl"
)

var _ = __VDLInit() // Must be first; see __VDLInit comments for details.

//////////////////////////////////////////////////
// Type definitions

// Url represents a url.URL struct.
// The User field is skipped since it is a struct with only unexported fields.
type Url struct {
	Scheme   string
	Opaque   string
	Host     string
	Path     string
	RawPath  string
	RawQuery string
	Fragment string
}

func (Url) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/http.Url"`
}) {
}

func (x Url) VDLIsZero() bool {
	return x == Url{}
}

func (x Url) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_1); err != nil {
		return err
	}
	if x.Scheme != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Scheme); err != nil {
			return err
		}
	}
	if x.Opaque != "" {
		if err := enc.NextFieldValueString(1, vdl.StringType, x.Opaque); err != nil {
			return err
		}
	}
	if x.Host != "" {
		if err := enc.NextFieldValueString(2, vdl.StringType, x.Host); err != nil {
			return err
		}
	}
	if x.Path != "" {
		if err := enc.NextFieldValueString(3, vdl.StringType, x.Path); err != nil {
			return err
		}
	}
	if x.RawPath != "" {
		if err := enc.NextFieldValueString(4, vdl.StringType, x.RawPath); err != nil {
			return err
		}
	}
	if x.RawQuery != "" {
		if err := enc.NextFieldValueString(5, vdl.StringType, x.RawQuery); err != nil {
			return err
		}
	}
	if x.Fragment != "" {
		if err := enc.NextFieldValueString(6, vdl.StringType, x.Fragment); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Url) VDLRead(dec vdl.Decoder) error {
	*x = Url{}
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
				x.Scheme = value
			}
		case 1:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Opaque = value
			}
		case 2:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Host = value
			}
		case 3:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Path = value
			}
		case 4:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.RawPath = value
			}
		case 5:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.RawQuery = value
			}
		case 6:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Fragment = value
			}
		}
	}
}

// Request represents the http.Request struct. The MultipartForm field is
// skipped since the docs indicate that it is ignored by clients in favor of
// the Body field.
type Request struct {
	Method           string
	Url              Url
	Proto            string
	ProtoMajor       int16
	ProtoMinor       int16
	Header           map[string][]string
	Body             []byte
	ContentLength    int64
	TransferEncoding []string
	Close            bool
	Host             string
	Form             map[string][]string
	PostForm         map[string][]string
	Trailer          map[string][]string
	RemoteAddr       string
	RequestUri       string
}

func (Request) VDLReflect(struct {
	Name string `vdl:"v.io/v23/services/http.Request"`
}) {
}

func (x Request) VDLIsZero() bool {
	if x.Method != "" {
		return false
	}
	if x.Url != (Url{}) {
		return false
	}
	if x.Proto != "" {
		return false
	}
	if x.ProtoMajor != 0 {
		return false
	}
	if x.ProtoMinor != 0 {
		return false
	}
	if len(x.Header) != 0 {
		return false
	}
	if len(x.Body) != 0 {
		return false
	}
	if x.ContentLength != 0 {
		return false
	}
	if len(x.TransferEncoding) != 0 {
		return false
	}
	if x.Close {
		return false
	}
	if x.Host != "" {
		return false
	}
	if len(x.Form) != 0 {
		return false
	}
	if len(x.PostForm) != 0 {
		return false
	}
	if len(x.Trailer) != 0 {
		return false
	}
	if x.RemoteAddr != "" {
		return false
	}
	if x.RequestUri != "" {
		return false
	}
	return true
}

func (x Request) VDLWrite(enc vdl.Encoder) error {
	if err := enc.StartValue(__VDLType_struct_2); err != nil {
		return err
	}
	if x.Method != "" {
		if err := enc.NextFieldValueString(0, vdl.StringType, x.Method); err != nil {
			return err
		}
	}
	if x.Url != (Url{}) {
		if err := enc.NextField(1); err != nil {
			return err
		}
		if err := x.Url.VDLWrite(enc); err != nil {
			return err
		}
	}
	if x.Proto != "" {
		if err := enc.NextFieldValueString(2, vdl.StringType, x.Proto); err != nil {
			return err
		}
	}
	if x.ProtoMajor != 0 {
		if err := enc.NextFieldValueInt(3, vdl.Int16Type, int64(x.ProtoMajor)); err != nil {
			return err
		}
	}
	if x.ProtoMinor != 0 {
		if err := enc.NextFieldValueInt(4, vdl.Int16Type, int64(x.ProtoMinor)); err != nil {
			return err
		}
	}
	if len(x.Header) != 0 {
		if err := enc.NextField(5); err != nil {
			return err
		}
		if err := __VDLWriteAnon_map_1(enc, x.Header); err != nil {
			return err
		}
	}
	if len(x.Body) != 0 {
		if err := enc.NextFieldValueBytes(6, __VDLType_list_4, x.Body); err != nil {
			return err
		}
	}
	if x.ContentLength != 0 {
		if err := enc.NextFieldValueInt(7, vdl.Int64Type, x.ContentLength); err != nil {
			return err
		}
	}
	if len(x.TransferEncoding) != 0 {
		if err := enc.NextField(8); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_2(enc, x.TransferEncoding); err != nil {
			return err
		}
	}
	if x.Close {
		if err := enc.NextFieldValueBool(9, vdl.BoolType, x.Close); err != nil {
			return err
		}
	}
	if x.Host != "" {
		if err := enc.NextFieldValueString(10, vdl.StringType, x.Host); err != nil {
			return err
		}
	}
	if len(x.Form) != 0 {
		if err := enc.NextField(11); err != nil {
			return err
		}
		if err := __VDLWriteAnon_map_1(enc, x.Form); err != nil {
			return err
		}
	}
	if len(x.PostForm) != 0 {
		if err := enc.NextField(12); err != nil {
			return err
		}
		if err := __VDLWriteAnon_map_1(enc, x.PostForm); err != nil {
			return err
		}
	}
	if len(x.Trailer) != 0 {
		if err := enc.NextField(13); err != nil {
			return err
		}
		if err := __VDLWriteAnon_map_1(enc, x.Trailer); err != nil {
			return err
		}
	}
	if x.RemoteAddr != "" {
		if err := enc.NextFieldValueString(14, vdl.StringType, x.RemoteAddr); err != nil {
			return err
		}
	}
	if x.RequestUri != "" {
		if err := enc.NextFieldValueString(15, vdl.StringType, x.RequestUri); err != nil {
			return err
		}
	}
	if err := enc.NextField(-1); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_map_1(enc vdl.Encoder, x map[string][]string) error {
	if err := enc.StartValue(__VDLType_map_3); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for key, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, key); err != nil {
			return err
		}
		if err := __VDLWriteAnon_list_2(enc, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func __VDLWriteAnon_list_2(enc vdl.Encoder, x []string) error {
	if err := enc.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if err := enc.SetLenHint(len(x)); err != nil {
		return err
	}
	for _, elem := range x {
		if err := enc.NextEntryValueString(vdl.StringType, elem); err != nil {
			return err
		}
	}
	if err := enc.NextEntry(true); err != nil {
		return err
	}
	return enc.FinishValue()
}

func (x *Request) VDLRead(dec vdl.Decoder) error {
	*x = Request{}
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
				x.Method = value
			}
		case 1:
			if err := x.Url.VDLRead(dec); err != nil {
				return err
			}
		case 2:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Proto = value
			}
		case 3:
			switch value, err := dec.ReadValueInt(16); {
			case err != nil:
				return err
			default:
				x.ProtoMajor = int16(value)
			}
		case 4:
			switch value, err := dec.ReadValueInt(16); {
			case err != nil:
				return err
			default:
				x.ProtoMinor = int16(value)
			}
		case 5:
			if err := __VDLReadAnon_map_1(dec, &x.Header); err != nil {
				return err
			}
		case 6:
			if err := dec.ReadValueBytes(-1, &x.Body); err != nil {
				return err
			}
		case 7:
			switch value, err := dec.ReadValueInt(64); {
			case err != nil:
				return err
			default:
				x.ContentLength = value
			}
		case 8:
			if err := __VDLReadAnon_list_2(dec, &x.TransferEncoding); err != nil {
				return err
			}
		case 9:
			switch value, err := dec.ReadValueBool(); {
			case err != nil:
				return err
			default:
				x.Close = value
			}
		case 10:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.Host = value
			}
		case 11:
			if err := __VDLReadAnon_map_1(dec, &x.Form); err != nil {
				return err
			}
		case 12:
			if err := __VDLReadAnon_map_1(dec, &x.PostForm); err != nil {
				return err
			}
		case 13:
			if err := __VDLReadAnon_map_1(dec, &x.Trailer); err != nil {
				return err
			}
		case 14:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.RemoteAddr = value
			}
		case 15:
			switch value, err := dec.ReadValueString(); {
			case err != nil:
				return err
			default:
				x.RequestUri = value
			}
		}
	}
}

func __VDLReadAnon_map_1(dec vdl.Decoder, x *map[string][]string) error {
	if err := dec.StartValue(__VDLType_map_3); err != nil {
		return err
	}
	var tmpMap map[string][]string
	if len := dec.LenHint(); len > 0 {
		tmpMap = make(map[string][]string, len)
	}
	for {
		switch done, key, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			*x = tmpMap
			return dec.FinishValue()
		default:
			var elem []string
			if err := __VDLReadAnon_list_2(dec, &elem); err != nil {
				return err
			}
			if tmpMap == nil {
				tmpMap = make(map[string][]string)
			}
			tmpMap[key] = elem
		}
	}
}

func __VDLReadAnon_list_2(dec vdl.Decoder, x *[]string) error {
	if err := dec.StartValue(__VDLType_list_5); err != nil {
		return err
	}
	if len := dec.LenHint(); len > 0 {
		*x = make([]string, 0, len)
	} else {
		*x = nil
	}
	for {
		switch done, elem, err := dec.NextEntryValueString(); {
		case err != nil:
			return err
		case done:
			return dec.FinishValue()
		default:
			*x = append(*x, elem)
		}
	}
}

//////////////////////////////////////////////////
// Interface definitions

// HttpClientMethods is the client interface
// containing Http methods.
type HttpClientMethods interface {
	// RawDo returns the server's response to req.
	RawDo(_ *context.T, req Request, _ ...rpc.CallOpt) (data []byte, _ error)
}

// HttpClientStub adds universal methods to HttpClientMethods.
type HttpClientStub interface {
	HttpClientMethods
	rpc.UniversalServiceMethods
}

// HttpClient returns a client stub for Http.
func HttpClient(name string) HttpClientStub {
	return implHttpClientStub{name}
}

type implHttpClientStub struct {
	name string
}

func (c implHttpClientStub) RawDo(ctx *context.T, i0 Request, opts ...rpc.CallOpt) (o0 []byte, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "RawDo", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

// HttpServerMethods is the interface a server writer
// implements for Http.
type HttpServerMethods interface {
	// RawDo returns the server's response to req.
	RawDo(_ *context.T, _ rpc.ServerCall, req Request) (data []byte, _ error)
}

// HttpServerStubMethods is the server interface containing
// Http methods, as expected by rpc.Server.
// There is no difference between this interface and HttpServerMethods
// since there are no streaming methods.
type HttpServerStubMethods HttpServerMethods

// HttpServerStub adds universal methods to HttpServerStubMethods.
type HttpServerStub interface {
	HttpServerStubMethods
	// Describe the Http interfaces.
	Describe__() []rpc.InterfaceDesc
}

// HttpServer returns a server stub for Http.
// It converts an implementation of HttpServerMethods into
// an object that may be used by rpc.Server.
func HttpServer(impl HttpServerMethods) HttpServerStub {
	stub := implHttpServerStub{
		impl: impl,
	}
	// Initialize GlobState; always check the stub itself first, to handle the
	// case where the user has the Glob method defined in their VDL source.
	if gs := rpc.NewGlobState(stub); gs != nil {
		stub.gs = gs
	} else if gs := rpc.NewGlobState(impl); gs != nil {
		stub.gs = gs
	}
	return stub
}

type implHttpServerStub struct {
	impl HttpServerMethods
	gs   *rpc.GlobState
}

func (s implHttpServerStub) RawDo(ctx *context.T, call rpc.ServerCall, i0 Request) ([]byte, error) {
	return s.impl.RawDo(ctx, call, i0)
}

func (s implHttpServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implHttpServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{HttpDesc}
}

// HttpDesc describes the Http interface.
var HttpDesc rpc.InterfaceDesc = descHttp

// descHttp hides the desc to keep godoc clean.
var descHttp = rpc.InterfaceDesc{
	Name:    "Http",
	PkgPath: "v.io/v23/services/http",
	Methods: []rpc.MethodDesc{
		{
			Name: "RawDo",
			Doc:  "// RawDo returns the server's response to req.",
			InArgs: []rpc.ArgDesc{
				{Name: "req", Doc: ``}, // Request
			},
			OutArgs: []rpc.ArgDesc{
				{Name: "data", Doc: ``}, // []byte
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
	},
}

// Hold type definitions in package-level variables, for better performance.
//nolint:unused
var (
	__VDLType_struct_1 *vdl.Type
	__VDLType_struct_2 *vdl.Type
	__VDLType_map_3    *vdl.Type
	__VDLType_list_4   *vdl.Type
	__VDLType_list_5   *vdl.Type
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
	vdl.Register((*Url)(nil))
	vdl.Register((*Request)(nil))

	// Initialize type definitions.
	__VDLType_struct_1 = vdl.TypeOf((*Url)(nil)).Elem()
	__VDLType_struct_2 = vdl.TypeOf((*Request)(nil)).Elem()
	__VDLType_map_3 = vdl.TypeOf((*map[string][]string)(nil))
	__VDLType_list_4 = vdl.TypeOf((*[]byte)(nil))
	__VDLType_list_5 = vdl.TypeOf((*[]string)(nil))

	return struct{}{}
}
