// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

import "unsafe"

func float64touint32(a float64) uint32
func uint32tofloat64(a uint32) float64

// stackcheck checks that SP is in range [g->stack.lo, g->stack.hi).
func stackcheck()

// Called from assembly only; declared for go vet.
func setldt(slot uintptr, base unsafe.Pointer, size uintptr)
func emptyfunc()

//go:noescape
func asmcgocall_no_g(fn, arg unsafe.Pointer)

// getcallerfp returns the address of the frame pointer in the callers frame or 0 if not implemented.
func getcallerfp() uintptr { return 0 }
