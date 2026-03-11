.PHONY: run-control-plane run-agent run-demo tidy

run-control-plane:
	go run ./apps/control-plane/cmd/server

run-agent:
	go run ./apps/agent/cmd/agent

run-demo:
	go run ./apps/demo-app/cmd/demo

tidy:
	cd apps/control-plane && go mod tidy
	cd apps/agent && go mod tidy
	cd apps/demo-app && go mod tidy
	cd packages/shared && go mod tidy
