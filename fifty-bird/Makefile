.PHONY: build

build: 
	@go build -o ./build/fifty-bird *.go
	@echo "[OK] Fifty Bird binary was created!"

run: build
	@./build/fifty-bird

clean:
	rm -rf build