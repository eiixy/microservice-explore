# 基本项目模板
* interface: 对外的 BFF 服务，接受来自用户的请求，比如暴露了 HTTP/gRPC 接口。
* service: 对内的微服务，仅接受来自内部其他服务或者网关的请求，比如暴露了gRPC 接口只对内服务。
* admin：区别于 service，更多是面向运营测的服务，通常数据权限更高，隔离带来更好的代码级别安全。
* job: 流式任务处理的服务，上游一般依赖 message broker。
* task: 定时任务，类似 cronjob，部署到 task 托管平台中。