// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated by the vanadium vdl tool.
// Package: version

//nolint:golint
package version

import (
	"v.io/v23/context"
	"v.io/v23/i18n"
	"v.io/v23/verror"
)

var _ = initializeVDL() // Must be first; see initializeVDL comments for details.

//////////////////////////////////////////////////
// Error definitions

var (
	ErrNoCompatibleVersion = verror.Register("v.io/v23/rpc/version.NoCompatibleVersion", verror.NoRetry, "{1:}{2:} There were no compatible versions between ({3},{4}) and ({5},{6}).")
)

// NewErrNoCompatibleVersion returns an error with the ErrNoCompatibleVersion ID.
func NewErrNoCompatibleVersion(ctx *context.T, lmin uint64, lmax uint64, rmin uint64, rmax uint64) error {
	return verror.New(ErrNoCompatibleVersion, ctx, lmin, lmax, rmin, rmax)
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
	i18n.Cat().SetWithBase(i18n.LangID("en"), i18n.MsgID(ErrNoCompatibleVersion.ID), "{1:}{2:} There were no compatible versions between ({3},{4}) and ({5},{6}).")

	return struct{}{}
}
