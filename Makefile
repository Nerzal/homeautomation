init:
	mkdir html

dashboard:
	rd /s /q ".\html"
	mkdir html
	xcopy /f ".\dashboard\wasm_exec.js" ".\html"
	xcopy /f ".\dashboard\wasm.js" ".\html"
	xcopy /f ".\dashboard\index.html" ".\html"
	tinygo build -o ".\html\wasm.wasm" --no-debug ".\dashboard\wasm.go"
	go run ".\api\main.go"