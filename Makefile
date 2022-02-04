SHELL = bash
.PHONY: test
test:
	for ((i=0; i <= 100; i++)) do go clean -testcache && go test -race; done