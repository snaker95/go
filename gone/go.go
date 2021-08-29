package gone

import (
	"fmt"
	"runtime"
)

/*
@Time : 2021/8/29 下午4:25
@Author : snaker95
@File : gone
@Software: GoLand
*/

func Go(f func()) {
	go func() {
		if goneRecover == nil {
			goneRecover = defaultRecover
		}
		defer goneRecover()
		f()
	}()
}

// goneRecover 封装 defer中recover()的处理方法
var goneRecover func()

// InitGoneRecover goneRecover(), 可以在启动服务的做一次性初始化
func InitGoneRecover(r func()) {
	goneRecover = r
}

// defaultRecover call recover to hold back the case which has a panic
func defaultRecover() {
	if p := recover(); p != nil {
		var st [4096]byte
		n := runtime.Stack(st[:], false)
		fmt.Printf("gone panic, err=%+v，stack trace:\n%v", p, string(st[:n]))
	}
}
