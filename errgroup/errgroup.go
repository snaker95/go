// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errgroup provides synchronization, error propagation, and Context
// cancelation for groups of goroutines working on subtasks of a common task.
// 完全参考 sync/errgroup 包, 在此基础上增加 recover 方法接收 panic 问题, 防止因 goroutine
// 中出现异常(例如: 空指针) 导致整个应用程序退出
package errgroup

import (
	"context"
	"fmt"
	"runtime"
	"sync"
)

// A Group is a collection of goroutines working on subtasks that are part of
// the same overall task.
//
// A zero Group is valid and does not cancel on error.
type Group struct {
	cancel func()

	wg sync.WaitGroup

	errOnce sync.Once
	err     error
}

// errgroupRecover 封装 defer中recover()的处理方法
var errgroupRecover func(p interface{})

// InitErrgroupRecover errgroupRecover(), 可以在启动服务的做一次性初始化
func InitErrgroupRecover(r func(interface{})) {
	errgroupRecover = r
}

// defaultRecover call recover to hold back the case which has a panic
func defaultRecover(p interface{}) {
	if p != nil {
		var st [4096]byte
		n := runtime.Stack(st[:], false)
		fmt.Printf("errgroup panic, err=%+v，stack trace:\n%v", p, string(st[:n]))
	}
}

// WithContext returns a new Group and an associated Context derived from ctx.
//
// The derived Context is canceled the first time a function passed to Go
// returns a non-nil error or the first time Wait returns, whichever occurs
// first.
func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}

func (g *Group) recover() {
	if p := recover(); p != nil {
		if errgroupRecover == nil {
			errgroupRecover = defaultRecover
		}
		errgroupRecover(p)
		g.errCancel(fmt.Errorf("panic: %+v", p))
	}
}

func (g *Group) errCancel(err error) {
	if err != nil {
		g.errOnce.Do(func() {
			g.err = err
			if g.cancel != nil {
				g.cancel()
			}
		})
	}
}

// Wait blocks until all function calls from the Go method have returned, then
// returns the first non-nil error (if any) from them.
func (g *Group) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}

	return g.err
}

// Go calls the given function in a new goroutine.
//
// The first call to return a non-nil error cancels the group; its error will be
// returned by Wait.
func (g *Group) Go(f func() error) {
	g.wg.Add(1)

	go func() {
		// defer 堆的特性, 所以 recover 先执行
		defer g.wg.Done()
		defer g.recover()

		if err := f(); err != nil {
			g.errCancel(err)
		}
	}()
}
