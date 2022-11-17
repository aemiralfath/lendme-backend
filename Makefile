create-mock:
	mockery --dir=./internal/auth --name=UseCase --output=./internal/auth/mocks
	mockery --dir=./internal/auth --name=Repository --output=./internal/auth/mocks
	mockery --dir=./internal/transaction --name=UseCase --output=./internal/transaction/mocks
	mockery --dir=./internal/transaction --name=Repository --output=./internal/transaction/mocks

.PHONY: test-coverage
test-coverage:
	go test -failfast -tags=integration -coverprofile=coverage.out -covermode=count ./internal/...
	go tool cover -func coverage.out
