package main

import "github.com/putitaT/thaiid/thaiid"

func main() {
	println(thaiid.ValidateThaiId("1234567890121"))
	println(thaiid.ValidateThaiId("1234567892121"))
}
