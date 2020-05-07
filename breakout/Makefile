.PHONY: build build_darwin_x64 build_windows_x64

game = breakout

compile-assets:
	@go generate assets/generate.go
	@echo "[OK] $(game) static assets compiled!"

build: compile-assets build_darwin_x64 build_windows_x64
	@echo "[OK] $(game) binary was created!"

build_darwin_x64:
	@GOOS=darwin GOARCH=amd64 go build -o ./build/macos-x64/$(game)

build_windows_x64:
	@GOOS=windows GOARCH=amd64 go build -o ./build/windows-x64/$(game).exe

run:
	go run *.go

clean:
	rm -rf build
	rm -rf assets/**/*.generated.go