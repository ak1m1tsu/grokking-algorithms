.PHONY: tests
tests:
	@go test -v -race -coverprofile=cover.out ./... \
	&& go tool cover -html=cover.out \
	&& rm cover.out
