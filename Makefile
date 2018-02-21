PACKAGE="github.com/hinshun/genplugin"

GOFILES=$(shell find . -type f -name '*.go')

DOCKERFILES=$(shell find . -type f -name 'Dockerfile')

.DEFAULT: all
all: genplugin

genplugin: $(GOFILES) $(DOCKERFILES)
	@echo "+ $@"
	@docker build -t hinshun/genplugin -f cmd/genplugin/Dockerfile .

.PHONY: all genplugin
