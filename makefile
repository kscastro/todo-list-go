SHELL := /bin/bash
dev-server:
	source .env && go run main.go server --profile development