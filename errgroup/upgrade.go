package errgroup

/*
@Time : 2021/11/20 下午9:50
@Author : snaker95
@File : upgrade
@Software: GoLand
*/

// Run 简化errgroup使用, 避免在代码中被打断
// example:
// eg, egCtx := WithContext(ctx)
// eg.Run(
// 	func() (err error) {
// 		err = func1(egCtx)
// 		return err
// 	},
// 	func() (err error) {
// 		param, err = func1(egCtx)
// 		return err
// 	},
// 	func() (err error) {
// 		func1(egCtx)
// 		return nil
// 	},
// 	func() (err error) {
// 		func1(ctx)
// 		return nil
// 	},
// )
// 注意 egCtx, 若传递到实现函数中, 其中有一个并发报错, 会通知其他传入了egCtx的未完成的线程取消,
// 而没传入 ctx 的函数仍然继续
func (g *Group) Run(f ...func() error) error {
	if f == nil {
		return nil
	}
	for i := range f {
		g.Go(f[i])
	}
	if err := g.Wait(); err != nil {
		return err
	}
	return nil
}
