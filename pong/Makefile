.PHONY: build

compile-assets:
	@go generate assets/generate.go
	@echo "[OK] Pong static assets compiled!"

build: compile-assets
	@go build -o ./build/pong *.go
	@echo "[OK] Pong binary was created!"

run: build
	@./build/pong

clean:
	rm -rf build