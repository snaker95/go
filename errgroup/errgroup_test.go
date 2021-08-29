package errgroup

import (
	"context"
	"fmt"
	"testing"
	"time"
)

/*
@Time : 2021/8/29 下午4:17
@Author : snaker95
@File : errgroup_test.gone
@Software: GoLand
*/

func TestWithContext(t *testing.T) {
	// 设置 recover 函数处理
	InitErrgroupRecover(func(i interface{}) {
		fmt.Println(i)
	})
	ctx := context.Background()
	eg, egCtx := WithContext(ctx)
	eg.Go(func() error {
		a(egCtx)
		return nil
	})
	eg.Go(func() error {
		b(egCtx)
		return nil
	})
	err := eg.Wait()
	t.Logf("eg wait err = %v", err)
}

func a(ctx context.Context) {

}
func b(ctx context.Context) {
	time.Sleep(1 * time.Second)
	panic("b - xx")
}
