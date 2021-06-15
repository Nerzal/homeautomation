init:
	mkdir html

dashboard-windows:
	rd /s /q ".\html"
	mkdir html
	xcopy /f ".\dashboard\wasm_exec.js" ".\html"
	xcopy /f ".\dashboard\wasm.js" ".\html"
	xcopy /f ".\dashboard\index.html" ".\html"
	xcopy /f ".\dashboard\styles\general.css" ".\html"
	mkdir ".\html\assets"
	xcopy /f ".\assets" ".\html\assets"
	tinygo build -o ".\html\wasm.wasm" -opt=s --no-debug ".\dashboard\wasm.go"
	go run ".\api\main.go"

dashboard-linux:
	rm -rf "html"
	mkdir html
	cp "dashboard/wasm_exec.js" "html/wasm_exec.js"
	cp "dashboard/wasm.js" "html/wasm.js"
	cp "dashboard/mqtt.js" "html/mqtt.js"
	cp "dashboard/index.html" "html/index.html"
	mkdir html/styles
	cp "dashboard/styles/general.css" "html/styles/general.css"
	mkdir "html/assets"
	cp -r assets html
	tinygo build --target=wasm -o html/wasm.wasm -opt=s --no-debug dashboard/wasm.go
	go run "api/main.go"

build-go-wasm:
	GOOS=js GOARCH=wasm go build -o html/wasmgo.wasm dashboard/wasm.go 

bedroom-client:
	tinygo flash --target=arduino-nano33 clients/bedroom/main.go

start-mqtt-broker:
	docker start mosquitto