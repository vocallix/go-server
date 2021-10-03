.DEFAULT_GOAL := help

start: ## start gamedata server
	@echo "start gamedata server"
	@go build
	@./scripts/run_server.sh

run-frontend-dev: webpack.PID ## Run the frontend and admin apps in dev (using webpack-dev-server)

webpack.PID:
	@./node_modules/.bin/babel-node ./node_modules/.bin/webpack-dev-server \
		--content-base=build \
		--devtool=cheap-module-inline-source-map \
		--hot \
		--inline \
		--progress \
		& echo "$$!" > webpack.PID

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'