package internal

//go:generate mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ocp-docs-api/internal/repo Repo
//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ocp-docs-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ocp-docs-api/internal/producer Producer
