// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package refluture

import (
	"reflect"
)

type Future struct {
	c <-chan []reflect.Value
	r []reflect.Value
}

func New(fn interface{}) *Future {
	c := make(chan []reflect.Value)
	fut := &Future{c: c}

	go func() {
		// no compile-time safety.
		// TODO: add runtime checks.
		c <- reflect.ValueOf(fn).Call(nil)
	}()

	return fut

}

// assumes error exists as the last arg.
func (f *Future) Await(out ...interface{}) error {
	if f.r == nil {
		f.r = <-f.c
	}

	for i, o := range out {
		v := f.r[i]

		ov := reflect.ValueOf(o)
		if !ov.IsValid() {
			continue // skip setting nils
		}
		ov = ov.Elem()
		ov.Set(v)
	}

	err := f.r[len(f.r)-1].Interface()
	if et, ok := err.(error); ok {
		return et
	}

	return nil
}
