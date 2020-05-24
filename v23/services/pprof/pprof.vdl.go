// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: pprof

// Package pprof defines an interface for accessing runtime profiling data in
// the format expected by the pprof visualization tool.  For more information
// about pprof, see http://code.google.com/p/google-perftools/.
//nolint:golint
package pprof

import (
	"io"

	v23 "v.io/v23"
	"v.io/v23/context"
	"v.io/v23/rpc"
	"v.io/v23/security/access"
	"v.io/v23/vdl"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Interface definitions

// PProfClientMethods is the client interface
// containing PProf methods.
type PProfClientMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(*context.T, ...rpc.CallOpt) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(*context.T, ...rpc.CallOpt) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(_ *context.T, name string, debug int32, _ ...rpc.CallOpt) (PProfProfileClientCall, error)
	// CpuProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CpuProfile(_ *context.T, seconds int32, _ ...rpc.CallOpt) (PProfCpuProfileClientCall, error)
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(_ *context.T, programCounters []uint64, _ ...rpc.CallOpt) ([]string, error)
}

// PProfClientStub adds universal methods to PProfClientMethods.
type PProfClientStub interface {
	PProfClientMethods
	rpc.UniversalServiceMethods
}

// PProfClient returns a client stub for PProf.
func PProfClient(name string) PProfClientStub {
	return implPProfClientStub{name}
}

type implPProfClientStub struct {
	name string
}

func (c implPProfClientStub) CmdLine(ctx *context.T, opts ...rpc.CallOpt) (o0 []string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "CmdLine", nil, []interface{}{&o0}, opts...)
	return
}

func (c implPProfClientStub) Profiles(ctx *context.T, opts ...rpc.CallOpt) (o0 []string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Profiles", nil, []interface{}{&o0}, opts...)
	return
}

func (c implPProfClientStub) Profile(ctx *context.T, i0 string, i1 int32, opts ...rpc.CallOpt) (ocall PProfProfileClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "Profile", []interface{}{i0, i1}, opts...); err != nil {
		return
	}
	ocall = &implPProfProfileClientCall{ClientCall: call}
	return
}

func (c implPProfClientStub) CpuProfile(ctx *context.T, i0 int32, opts ...rpc.CallOpt) (ocall PProfCpuProfileClientCall, err error) {
	var call rpc.ClientCall
	if call, err = v23.GetClient(ctx).StartCall(ctx, c.name, "CpuProfile", []interface{}{i0}, opts...); err != nil {
		return
	}
	ocall = &implPProfCpuProfileClientCall{ClientCall: call}
	return
}

func (c implPProfClientStub) Symbol(ctx *context.T, i0 []uint64, opts ...rpc.CallOpt) (o0 []string, err error) {
	err = v23.GetClient(ctx).Call(ctx, c.name, "Symbol", []interface{}{i0}, []interface{}{&o0}, opts...)
	return
}

// PProfProfileClientStream is the client stream for PProf.Profile.
type PProfProfileClientStream interface {
	// RecvStream returns the receiver side of the PProf.Profile client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// PProfProfileClientCall represents the call returned from PProf.Profile.
type PProfProfileClientCall interface {
	PProfProfileClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implPProfProfileClientCall struct {
	rpc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implPProfProfileClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implPProfProfileClientCallRecv{c}
}

type implPProfProfileClientCallRecv struct {
	c *implPProfProfileClientCall
}

func (c implPProfProfileClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implPProfProfileClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implPProfProfileClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implPProfProfileClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// PProfCpuProfileClientStream is the client stream for PProf.CpuProfile.
type PProfCpuProfileClientStream interface {
	// RecvStream returns the receiver side of the PProf.CpuProfile client stream.
	RecvStream() interface {
		// Advance stages an item so that it may be retrieved via Value.  Returns
		// true iff there is an item to retrieve.  Advance must be called before
		// Value is called.  May block if an item is not available.
		Advance() bool
		// Value returns the item that was staged by Advance.  May panic if Advance
		// returned false or was not called.  Never blocks.
		Value() []byte
		// Err returns any error encountered by Advance.  Never blocks.
		Err() error
	}
}

// PProfCpuProfileClientCall represents the call returned from PProf.CpuProfile.
type PProfCpuProfileClientCall interface {
	PProfCpuProfileClientStream
	// Finish blocks until the server is done, and returns the positional return
	// values for call.
	//
	// Finish returns immediately if the call has been canceled; depending on the
	// timing the output could either be an error signaling cancelation, or the
	// valid positional return values from the server.
	//
	// Calling Finish is mandatory for releasing stream resources, unless the call
	// has been canceled or any of the other methods return an error.  Finish should
	// be called at most once.
	Finish() error
}

type implPProfCpuProfileClientCall struct {
	rpc.ClientCall
	valRecv []byte
	errRecv error
}

func (c *implPProfCpuProfileClientCall) RecvStream() interface {
	Advance() bool
	Value() []byte
	Err() error
} {
	return implPProfCpuProfileClientCallRecv{c}
}

type implPProfCpuProfileClientCallRecv struct {
	c *implPProfCpuProfileClientCall
}

func (c implPProfCpuProfileClientCallRecv) Advance() bool {
	c.c.errRecv = c.c.Recv(&c.c.valRecv)
	return c.c.errRecv == nil
}
func (c implPProfCpuProfileClientCallRecv) Value() []byte {
	return c.c.valRecv
}
func (c implPProfCpuProfileClientCallRecv) Err() error {
	if c.c.errRecv == io.EOF {
		return nil
	}
	return c.c.errRecv
}
func (c *implPProfCpuProfileClientCall) Finish() (err error) {
	err = c.ClientCall.Finish()
	return
}

// PProfServerMethods is the interface a server writer
// implements for PProf.
type PProfServerMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(*context.T, rpc.ServerCall) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(*context.T, rpc.ServerCall) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(_ *context.T, _ PProfProfileServerCall, name string, debug int32) error
	// CpuProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CpuProfile(_ *context.T, _ PProfCpuProfileServerCall, seconds int32) error
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(_ *context.T, _ rpc.ServerCall, programCounters []uint64) ([]string, error)
}

// PProfServerStubMethods is the server interface containing
// PProf methods, as expected by rpc.Server.
// The only difference between this interface and PProfServerMethods
// is the streaming methods.
type PProfServerStubMethods interface {
	// CmdLine returns the command-line arguments of the server, including
	// the name of the executable.
	CmdLine(*context.T, rpc.ServerCall) ([]string, error)
	// Profiles returns the list of available profiles.
	Profiles(*context.T, rpc.ServerCall) ([]string, error)
	// Profile streams the requested profile. The debug parameter enables
	// additional output. Passing debug=0 includes only the hexadecimal
	// addresses that pprof needs. Passing debug=1 adds comments translating
	// addresses to function names and line numbers, so that a programmer
	// can read the profile without tools.
	Profile(_ *context.T, _ *PProfProfileServerCallStub, name string, debug int32) error
	// CpuProfile enables CPU profiling for the requested duration and
	// streams the profile data.
	CpuProfile(_ *context.T, _ *PProfCpuProfileServerCallStub, seconds int32) error
	// Symbol looks up the program counters and returns their respective
	// function names.
	Symbol(_ *context.T, _ rpc.ServerCall, programCounters []uint64) ([]string, error)
}

// PProfServerStub adds universal methods to PProfServerStubMethods.
type PProfServerStub interface {
	PProfServerStubMethods
	// DescribeInterfaces the PProf interfaces.
	Describe__() []rpc.InterfaceDesc
}

// PProfServer returns a server stub for PProf.
// It converts an implementation of PProfServerMethods into
// an object that may be used by rpc.Server.
func PProfServer(impl PProfServerMethods) PProfServerStub {
	stub := implPProfServerStub{
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

type implPProfServerStub struct {
	impl PProfServerMethods
	gs   *rpc.GlobState
}

func (s implPProfServerStub) CmdLine(ctx *context.T, call rpc.ServerCall) ([]string, error) {
	return s.impl.CmdLine(ctx, call)
}

func (s implPProfServerStub) Profiles(ctx *context.T, call rpc.ServerCall) ([]string, error) {
	return s.impl.Profiles(ctx, call)
}

func (s implPProfServerStub) Profile(ctx *context.T, call *PProfProfileServerCallStub, i0 string, i1 int32) error {
	return s.impl.Profile(ctx, call, i0, i1)
}

func (s implPProfServerStub) CpuProfile(ctx *context.T, call *PProfCpuProfileServerCallStub, i0 int32) error {
	return s.impl.CpuProfile(ctx, call, i0)
}

func (s implPProfServerStub) Symbol(ctx *context.T, call rpc.ServerCall, i0 []uint64) ([]string, error) {
	return s.impl.Symbol(ctx, call, i0)
}

func (s implPProfServerStub) Globber() *rpc.GlobState {
	return s.gs
}

func (s implPProfServerStub) Describe__() []rpc.InterfaceDesc {
	return []rpc.InterfaceDesc{PProfDesc}
}

// PProfDesc describes the PProf interface.
var PProfDesc rpc.InterfaceDesc = descPProf

// descPProf hides the desc to keep godoc clean.
var descPProf = rpc.InterfaceDesc{
	Name:    "PProf",
	PkgPath: "v.io/v23/services/pprof",
	Methods: []rpc.MethodDesc{
		{
			Name: "CmdLine",
			Doc:  "// CmdLine returns the command-line arguments of the server, including\n// the name of the executable.",
			OutArgs: []rpc.ArgDesc{
				{Name: "", Doc: ``}, // []string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "Profiles",
			Doc:  "// Profiles returns the list of available profiles.",
			OutArgs: []rpc.ArgDesc{
				{Name: "", Doc: ``}, // []string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "Profile",
			Doc:  "// Profile streams the requested profile. The debug parameter enables\n// additional output. Passing debug=0 includes only the hexadecimal\n// addresses that pprof needs. Passing debug=1 adds comments translating\n// addresses to function names and line numbers, so that a programmer\n// can read the profile without tools.",
			InArgs: []rpc.ArgDesc{
				{Name: "name", Doc: ``},  // string
				{Name: "debug", Doc: ``}, // int32
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "CpuProfile",
			Doc:  "// CpuProfile enables CPU profiling for the requested duration and\n// streams the profile data.",
			InArgs: []rpc.ArgDesc{
				{Name: "seconds", Doc: ``}, // int32
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
		{
			Name: "Symbol",
			Doc:  "// Symbol looks up the program counters and returns their respective\n// function names.",
			InArgs: []rpc.ArgDesc{
				{Name: "programCounters", Doc: ``}, // []uint64
			},
			OutArgs: []rpc.ArgDesc{
				{Name: "", Doc: ``}, // []string
			},
			Tags: []*vdl.Value{vdl.ValueOf(access.Tag("Debug"))},
		},
	},
}

// PProfProfileServerStream is the server stream for PProf.Profile.
type PProfProfileServerStream interface {
	// SendStream returns the send side of the PProf.Profile server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// PProfProfileServerCall represents the context passed to PProf.Profile.
type PProfProfileServerCall interface {
	rpc.ServerCall
	PProfProfileServerStream
}

// PProfProfileServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements PProfProfileServerCall.
type PProfProfileServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes PProfProfileServerCallStub from rpc.StreamServerCall.
func (s *PProfProfileServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the PProf.Profile server stream.
func (s *PProfProfileServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implPProfProfileServerCallSend{s}
}

type implPProfProfileServerCallSend struct {
	s *PProfProfileServerCallStub
}

func (s implPProfProfileServerCallSend) Send(item []byte) error {
	return s.s.Send(item)
}

// PProfCpuProfileServerStream is the server stream for PProf.CpuProfile.
type PProfCpuProfileServerStream interface {
	// SendStream returns the send side of the PProf.CpuProfile server stream.
	SendStream() interface {
		// Send places the item onto the output stream.  Returns errors encountered
		// while sending.  Blocks if there is no buffer space; will unblock when
		// buffer space is available.
		Send(item []byte) error
	}
}

// PProfCpuProfileServerCall represents the context passed to PProf.CpuProfile.
type PProfCpuProfileServerCall interface {
	rpc.ServerCall
	PProfCpuProfileServerStream
}

// PProfCpuProfileServerCallStub is a wrapper that converts rpc.StreamServerCall into
// a typesafe stub that implements PProfCpuProfileServerCall.
type PProfCpuProfileServerCallStub struct {
	rpc.StreamServerCall
}

// Init initializes PProfCpuProfileServerCallStub from rpc.StreamServerCall.
func (s *PProfCpuProfileServerCallStub) Init(call rpc.StreamServerCall) {
	s.StreamServerCall = call
}

// SendStream returns the send side of the PProf.CpuProfile server stream.
func (s *PProfCpuProfileServerCallStub) SendStream() interface {
	Send(item []byte) error
} {
	return implPProfCpuProfileServerCallSend{s}
}

type implPProfCpuProfileServerCallSend struct {
	s *PProfCpuProfileServerCallStub
}

func (s implPProfCpuProfileServerCallSend) Send(item []byte) error {
	return s.s.Send(item)
}

var initializeVDLCalled bool

// initializeVDL performs vdl initialization.  It is safe to call multiple times.
// If you have an init ordering issue, just insert the following line verbatim
// into your source files in this package, right after the "package foo" clause:
//
//    var _ = initializeVDL()
//
// The purpose of this function is to ensure that vdl initialization occurs in
// the right order, and very early in the init sequence.  In particular, vdl
// registration and package variable initialization needs to occur before
// functions like vdl.TypeOf will work properly.
//
// This function returns a dummy value, so that it can be used to initialize the
// first var in the file, to take advantage of Go's defined init order.
func initializeVDL() struct{} {
	if initializeVDLCalled {
		return struct{}{}
	}
	initializeVDLCalled = true

	return struct{}{}
}
