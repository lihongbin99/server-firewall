package main

import "security-network/common/security"

func main() {
	err := security.GenerateRSAFile(
		"C:\\Users\\11977\\Documents\\bill\\private.pem",
		"C:\\Users\\11977\\Documents\\bill\\public.pem",
	)
	if err != nil {
		panic(err)
	}
}
