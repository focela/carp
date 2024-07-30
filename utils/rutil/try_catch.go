// Copyright (c) 2024 Focela Technologies. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rutil

import (
	"context"

	"github.com/focela/ratcatcher/errors/rcode"
	"github.com/focela/ratcatcher/errors/rerror"
)

// Throw throws out an exception, which can be caught be TryCatch or recover.
func Throw(exception interface{}) {
	panic(exception)
}

// Try implements try... logistics using internal panic...recover.
// It returns error if any exception occurs, or else it returns nil.
func Try(ctx context.Context, try func(ctx context.Context)) (err error) {
	if try == nil {
		return
	}
	defer func() {
		if exception := recover(); exception != nil {
			if v, ok := exception.(error); ok && rerror.HasStack(v) {
				err = v
			} else {
				err = rerror.NewCodef(rcode.CodeInternalPanic, "%+v", exception)
			}
		}
	}()
	try(ctx)
	return
}

// TryCatch implements `try...catch..`. logistics using internal `panic...recover`.
// It automatically calls function `catch` if any exception occurs and passes the exception as an error.
// If `catch` is given nil, it ignores the panic from `try` and no panic will throw to parent goroutine.
//
// But, note that, if function `catch` also throws panic, the current goroutine will panic.
func TryCatch(ctx context.Context, try func(ctx context.Context), catch func(ctx context.Context, exception error)) {
	if try == nil {
		return
	}
	if exception := Try(ctx, try); exception != nil && catch != nil {
		catch(ctx, exception)
	}
}