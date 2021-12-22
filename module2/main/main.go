package main

import (
	"fmt"
	"github.com/vlpolak/swtgo/hasher"
)

func main() {
	password := "password"
	hash, _ := hasher.HashPassword(password)

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := hasher.CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
