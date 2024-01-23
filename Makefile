# I am a -*-Makefile-*-
#
# --------------------------------------------------
#

GO_FLAGS =

.PHONY: all
all: test lint

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:

.PHONY: clean
clean:
	go clean -cache
	go clean ./...
