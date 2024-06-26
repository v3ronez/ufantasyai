run: build
	@./bin/ufantasyai

install:
	@go install github.com/a-h/templ/cmd/templ@latest
	@go get ./...
	@go mod vendor
	@go mod tidy
	@go mod download
	@npm install -D tailwindcss
	@npm install -D daisyui@latest

# css:
templ:
	@templ generate --watch --proxy=http://localhost:3000

tailwind:
	@npx tailwind -i view/css/app.css -o public/styles.css --watch

build:
	npx tailwindcss -i view/css/app.css -o public/styles.css
	@templ generate view
	@go build -o bin/ufantasyai main.go
# @go build -tags dev -o bin/ufantasyai main.go
up: ## Database migration up
	@go run cmd/migrate/main.go up

reset:
	@go run cmd/reset/main.go up

down: ## Database migration down
	@go run cmd/migrate/main.go down

migration: ## Migrations against the database
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

seed:
	@go run cmd/seed/main.go


# hot reload
# make tailwind > make templ > air