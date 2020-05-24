// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: pattern

//nolint:golint
package pattern

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrIllegalEscapeChar = verror.Register("v.io/v23/query/pattern.IllegalEscapeChar", verror.NoRetry, "{1:}{2:} '%' and '_' cannot be used as escape characters")
	ErrInvalidEscape     = verror.Register("v.io/v23/query/pattern.InvalidEscape", verror.NoRetry, "{1:}{2:} only '%', '_', and the escape character are allowed to be escaped, found '\\{3}'")
)

// NewErrIllegalEscapeChar returns an error with the ErrIllegalEscapeChar ID.
func NewErrIllegalEscapeChar(ctx *context.T) error {
	return verror.New(ErrIllegalEscapeChar, ctx)
}

// NewErrInvalidEscape returns an error with the ErrInvalidEscape ID.
func NewErrInvalidEscape(ctx *context.T, escaped string) error {
	return verror.New(ErrInvalidEscape, ctx, escaped)
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
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrIllegalEscapeChar.ID), "{1:}{2:} '%' and '_' cannot be used as escape characters")
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrInvalidEscape.ID), "{1:}{2:} only '%', '_', and the escape character are allowed to be escaped, found '\\{3}'")

	return struct{}{}
}
