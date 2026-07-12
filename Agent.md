# back-service Agent 说明

## 项目定位
- back-service 是分布式后端业务服务，承载核心业务逻辑。
- 采用 tRPC 风格工程结构，面向服务间通信与业务处理。

## 目录职责
- app/controller：接口层与请求入口处理。
- app/logic：核心业务编排与领域逻辑。
- app/models：数据模型与结构定义。
- deploy/helm：Kubernetes Helm 部署清单。
- provider/conf：运行时配置（custom.yaml、trpc_go.yaml）。

## 与其他子项目关系
- 被 back-gateway 调用，提供业务能力。
- 协议定义依赖 proto-test 维护的 proto 与生成代码。
- 配置来源可对接 server-conf，对应环境目录进行挂载或注入。
- 服务治理能力由 polaris-* 提供，不在本项目重复实现。

## 配置与部署
- 本地与容器构建入口在 provider 目录。
- 集群部署优先使用 deploy/helm 下模板。
- 不同环境通过 server-conf 进行配置分层管理。

## 变更边界
- 允许改动：业务逻辑、接口编排、模型演进。
- 慎改动：基础设施耦合逻辑（注册发现、网格注入机制）。
- 若出现流量治理异常，先核查 Polaris 侧状态，再回到业务代码。