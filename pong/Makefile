.PHONY: build

build: 
	@go build -o ./build/pong *.go
	@echo "[OK] Pong binary was created!"

run: build
	@./build/pong

clean:
	rm -rf build