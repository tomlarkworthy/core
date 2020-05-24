// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: reserved

//nolint:golint
package reserved

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Error definitions

var (

	// ErrGlobMaxRecursionReached indicates that the Glob request exceeded the
	// max recursion level.
	ErrGlobMaxRecursionReached = verror.Register("v.io/v23/rpc/reserved.GlobMaxRecursionReached", verror.NoRetry, "{1:}{2:} max recursion level reached{:_}")
	// ErrGlobMatchesOmitted indicates that some of the Glob results might
	// have been omitted due to access restrictions.
	ErrGlobMatchesOmitted = verror.Register("v.io/v23/rpc/reserved.GlobMatchesOmitted", verror.NoRetry, "{1:}{2:} some matches might have been omitted")
	// ErrGlobNotImplemented indicates that Glob is not implemented by the
	// object.
	ErrGlobNotImplemented = verror.Register("v.io/v23/rpc/reserved.GlobNotImplemented", verror.NoRetry, "{1:}{2:} Glob not implemented")
)

// NewErrGlobMaxRecursionReached returns an error with the ErrGlobMaxRecursionReached ID.
func NewErrGlobMaxRecursionReached(ctx *context.T) error {
	return verror.New(ErrGlobMaxRecursionReached, ctx)
}

// NewErrGlobMatchesOmitted returns an error with the ErrGlobMatchesOmitted ID.
func NewErrGlobMatchesOmitted(ctx *context.T) error {
	return verror.New(ErrGlobMatchesOmitted, ctx)
}

// NewErrGlobNotImplemented returns an error with the ErrGlobNotImplemented ID.
func NewErrGlobNotImplemented(ctx *context.T) error {
	return verror.New(ErrGlobNotImplemented, ctx)
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

	// Set error format strings.
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrGlobMaxRecursionReached.ID), "{1:}{2:} max recursion level reached{:_}")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrGlobMatchesOmitted.ID), "{1:}{2:} some matches might have been omitted")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrGlobNotImplemented.ID), "{1:}{2:} Glob not implemented")

	return struct{}{}
}
