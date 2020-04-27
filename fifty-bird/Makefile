.PHONY: build

compile-assets:
	@go generate assets/generate.go
	@echo "[OK] Fifty Bird static assets compiled!"

build: compile-assets
	@go build -o ./build/fifty-bird *.go
	@echo "[OK] Fifty Bird binary was created!"

run: build
	@./build/fifty-bird

clean:
	rm -rf build
	rm -rf assets/**/*.generated.go