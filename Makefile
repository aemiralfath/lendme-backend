create-mock:
	mockery --dir=./internal/auth --name=UseCase --output=./internal/auth/mocks
	mockery --dir=./internal/user --name=UseCase --output=./internal/user/mocks
	mockery --dir=./internal/admin --name=UseCase --output=./internal/admin/mocks

.PHONY: test-coverage
test-coverage:
	go test -failfast -tags=integration -coverprofile=coverage.out -covermode=count ./internal/...
	go tool cover -func coverage.out
