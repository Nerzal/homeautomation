init:
	mkdir html

dashboard:
	rd /s /q ".\html"
	mkdir html
	xcopy /f ".\dashboard\wasm_exec.js" ".\html"
	xcopy /f ".\dashboard\wasm.js" ".\html"
	xcopy /f ".\dashboard\index.html" ".\html"
	xcopy /f ".\dashboard\styles\general.css" ".\html"
	tinygo build -o ".\html\wasm.wasm" -opt=s --no-debug ".\dashboard\wasm.go"
	go run ".\api\main.go"