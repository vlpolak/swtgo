# This project is created for Switch to Go program

#Package hasher contains functions for creating password hash codes and testing them

The example of using

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


