.PHONY: \
		help \
    	lint \
    	fmt \

all: imports fmt lint

help:
	@echo 'Usage: make <OPTIONS> ... <TARGETS>'
	@echo ''
	@echo 'Available targets are:'
	@echo ''
	@echo '    help               Show this help screen.'
	@echo '    lint               Run golint.'
	@echo '    fmt                Run go fmt.'

fmt:
	go install mvdan.cc/gofumpt@latest
	gofumpt -l -w -extra .

deps:
	go install golang.org/x/lint/golint

lint: deps
	golint ./...