FROM mcr.microsoft.com/devcontainers/go

# Install line-wrap formatter
RUN go install github.com/segmentio/golines@latest

# Install linter
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Install mock generation tool for testing
RUN go install go.uber.org/mock/mockgen@latest
