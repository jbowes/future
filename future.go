// Copyright (c) 2021 James Bowes. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package future

import (
	"context"
	"fmt"
	"time"
)

type result[T, K any] struct {
	t T
	k K
}

type Future[T, K any] struct {
	c <-chan *result[T, K]
	r *result[T, K]
}

func New[T, K any](fn func() (T, K)) *Future[T, K] {
	c := make(chan *result[T, K])
	fut := &Future[T, K]{c: c}

	go func() {
		r := result[T, K]{}
		r.t, r.k = fn()
		c <- &r
	}()

	return fut

}

func (f *Future[T, K]) Await() (T, K) {
	if f.r == nil {
		f.r = <-f.c
	}

	return f.r.t, f.r.k
}

type result3[T, K, L any] struct {
	t T
	k K
	l L
}

type Future3[T, K, L any] struct {
	c <-chan *result3[T, K, L]
	r *result3[T, K, L]
}

func New3[T, K, L any](fn func() (T, K, L)) *Future3[T, K, L] {
	c := make(chan *result3[T, K, L])
	fut := &Future3[T, K, L]{c: c}

	go func() {
		r := result3[T, K, L]{}
		r.t, r.k, r.l = fn()
		c <- &r
	}()

	return fut

}

func (f *Future3[T, K]) Await() (T, K, L) {
	if f.r == nil {
		f.r = <-f.c
	}

	return f.r.t, f.r.k, f.r.l
}