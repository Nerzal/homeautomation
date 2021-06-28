package main

import "encoding/base64"

func main() {
	base64String := base64.StdEncoding.EncodeToString([]byte("thisIs.A.JWT"))
	decodedData, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		panic(err)
	}

	println(string(decodedData))
}
