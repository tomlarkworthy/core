#!/bin/bash
# Copyright 2016 The Vanadium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

# Generate assets.go by running go-bindata.
# The author did not expect this script to be invoked directly,
# instead it is expected to be invoked via:
# jiri go generate v.io/x/ref/services/ben/benarchd/internal

set -euf -o pipefail

# Install go-bindata.
go get "github.com/cosnicolaou/go-bindata/v3/...@v3.0.8"

DIR=$(dirname $0)
cd "${DIR}/assets"
OUT="assets.go"
TMP="assets.tmp.go"

go run "github.com/cosnicolaou/go-bindata/v3/go-bindata" \
    -o "${TMP}" -pkg assets -ignore '\.go$' -nometadata -mode 0644 .
go fmt "${TMP}" >/dev/null
cat - "${TMP}" > "${OUT}" << EOF
// Copyright 2016 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package assets contains template strings and other assets for the benarchd web interface.
//
// This package is auto-generated by "go generate v.io/x/ref/services/ben/benarchd/internal"
// which in-turn uses https://github.com/cosnicolaou/go-bindata/
EOF
rm "${TMP}"
