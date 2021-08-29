# sync-go
golang 开发包, 针对golang语言中goroutine进行安全封装
## errgroup 包
errgroup 为 google/x/sync/errgroup 的二次开发, 支持 recover(), 有效防止并发时, 
由于并发协程产生 panic错误造成主协程退出, 尤其对于web 服务, 由并发造成主服务退出, 从而
影响提供对外服务
* 例如:
    * InitErrgroupRecover(i interface{}) 主协程初始化时, 设置使用该包并发时产生panic,
    recover()的处理方式,

## gone 包
gone为 关键字 go 进行二次封装, 支持 recover()统一处理; 避免影响主协程
* 例如:
    * InitGoneRecover(i interface{}) 主协程初始化时, 设置使用该包并发时产生panic,
    recover()的处理方式, 