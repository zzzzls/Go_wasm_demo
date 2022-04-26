.PHONY: run build server openurl status stop

PORT := 9091
SERVER_PID := $(shell cat fileServer.pid)

run: build openurl

build:
	@GOOS=js GOARCH=wasm go build -o test.wasm goWasmDemo.go
	@# 使用 brotli 压缩, 7m => 1.4m
	@# @brotli -f -o test.wasm.br test.wasm
	@go build fileServer.go
	@echo "===Build complete"

server:
ifneq ($(SERVER_PID),)
	kill -9 $(SERVER_PID)
endif
	@nohup ./fileServer >/dev/null 2>&1 & echo $$! >fileServer.pid
	@echo "===Server started at :$(PORT)..."

openurl: server
	@open http://127.0.0.1:$(PORT)/index.html

status:
ifneq ($(SERVER_PID),)
	@echo "Running"
else
	@echo "Stopped"
endif

stop:
ifneq ($(SERVER_PID),)
	@kill -9 $(SERVER_PID)
	@echo "" >fileServer.pid
endif
	@echo "===Server stopped"