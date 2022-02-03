.PHONY: test
test:
	for ((i=1; i <= 20; ++i)) do go test -v; done