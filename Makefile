build:
	GOARCH=wasm GOOS=js go build -o web/app.wasm
	go build -o launcher-server

run: build
	./launcher-server