build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	GOARCH=wasm GOOS=js go build -o docs/web/app.wasm
	go build -o launcher-server

run: build
	./launcher-server