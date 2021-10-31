.DEFAULT_GOAL := help

start: ## start gamedata server
	@echo "start gamedata server"
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

start_mongo: ## start mongodb using by docker
	@echo "start mongodb using by docker"
	@docker run -d -p 27017:27017 mongo

# restart_mongo: ## start mongodb using by docker
# 	@echo "start mongodb using by docker"
# 	@docker start $(docker ps -a | grep mongo | awk '{print $1}')

.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'