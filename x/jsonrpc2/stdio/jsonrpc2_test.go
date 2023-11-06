// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package stdio_test

import (
	"context"
	"testing"

	"github.com/goplus/gop/x/jsonrpc2"
	"github.com/goplus/gop/x/jsonrpc2/jsonrpc2test/cases"
	"github.com/goplus/gop/x/jsonrpc2/stdio"
)

func TestStdio(t *testing.T) {
	ctx := context.Background()
	listener := stdio.Listener(true)
	cases.Test(t, ctx, listener, jsonrpc2.HeaderFramer())
}
