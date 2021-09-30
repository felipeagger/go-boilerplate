GIT_CURRENT_BRANCH := ${shell git symbolic-ref --short HEAD}

start:
	docker-compose up -d

stop:
	docker-compose down

export_envs:
	export $(cat .env-sample | xargs)

build:
	go build -o api cmd/api/main.go

run:
	go run cmd/api/main.go

doc:
	swag init -g cmd/api/main.go

test:
	@echo "Input Package Name (Ex: pkg/utils)"
	@read INPUT_PKG; go test -v -race -cover "github.com/felipeagger/go-boilerplate/$$INPUT_PKG"

release:
	@if [ "$(v)" == "" ]; then \
		echo "You need to specify the new release version. Ex: make release v=1.0.0"; \
		exit 1; \
	fi
	@echo "Creating a new release tag version: ${v}"
	@git tag ${v}
	@git push origin ${v}
	@git push --set-upstream origin "${GIT_CURRENT_BRANCH}"
	@git push origin