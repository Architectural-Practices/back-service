# Helm Makefile for polaris-provider
# 使用方式: make <target>

RELEASE_NAME      ?= polaris-provider
NAMESPACE         ?= default
CHART_DIR         ?= .
VALUES_FILE       ?= values.yaml
TIMEOUT           ?= 300s

## -------------------- 部署相关 --------------------

.PHONY: install
install: ## 安装 Helm chart 到集群
	helm install $(RELEASE_NAME) $(CHART_DIR) \
		--namespace $(NAMESPACE) \
		--values $(VALUES_FILE) \
		--timeout $(TIMEOUT)

.PHONY: upgrade
upgrade: ## 升级 Helm release
	helm upgrade $(RELEASE_NAME) $(CHART_DIR) \
		--namespace $(NAMESPACE) \
		--values $(VALUES_FILE) \
		--timeout $(TIMEOUT) \
		--install

.PHONY: deploy
deploy: upgrade ## 部署/更新（install 或 upgrade，幂等）

.PHONY: rollback
rollback: ## 回滚到上一个版本
	helm rollback $(RELEASE_NAME) $(NAMESPACE) \
		--namespace $(NAMESPACE)

.PHONY: uninstall
uninstall: ## 卸载 release
	helm uninstall $(RELEASE_NAME) --namespace $(NAMESPACE)

## -------------------- 查看与调试 --------------------

.PHONY: template
template: ## 渲染模板并输出到终端
	helm template $(RELEASE_NAME) $(CHART_DIR) \
		--namespace $(NAMESPACE) \
		--values $(VALUES_FILE)

.PHONY: render
render: ## 渲染模板并保存到 rendered.yaml
	helm template $(RELEASE_NAME) $(CHART_DIR) \
		--namespace $(NAMESPACE) \
		--values $(VALUES_FILE) > rendered.yaml

.PHONY: lint
lint: ## 校验 chart 配置
	helm lint $(CHART_DIR) --values $(VALUES_FILE)

.PHONY: status
status: ## 查看 release 状态
	helm status $(RELEASE_NAME) --namespace $(NAMESPACE)

.PHONY: list
list: ## 列出所有 release
	helm list --namespace $(NAMESPACE)

.PHONY: history
history: ## 查看 release 变更历史
	helm history $(RELEASE_NAME) --namespace $(NAMESPACE)

## -------------------- 打包 --------------------

.PHONY: package
package: ## 打包 chart 为 .tgz
	helm package $(CHART_DIR) --destination ./dist

## -------------------- 组合流程 --------------------

.PHONY: redeploy
redeploy: uninstall deploy ## 重新部署（先卸载再安装）

.PHONY: full-deploy
full-deploy: lint template deploy ## 校验 -> 预览 -> 部署

## -------------------- 帮助 --------------------

.PHONY: help
help: ## 显示帮助信息
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| sort \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'
