.PHONY: test pre-commit gen-model gen-migrate-sql

test:
	go test -v -cover ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run -E "bodyclose,dogsled,durationcheck,errorlint,exportloopref,forbidigo,forcetypeassert,gocritic,godox,gofmt,goprintffuncname,gosec,ifshort,makezero,misspell,nakedret,nestif,nilerr,noctx,predeclared,rowserrcheck,sqlclosecheck,tparallel,unconvert,wastedassign"

run:
	go run ./cmd/app