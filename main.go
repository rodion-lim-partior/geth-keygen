// Copyright 2016 The go-ethereum Authors
// This file is part of go-ethereum.
//
// go-ethereum is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// go-ethereum is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with go-ethereum. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// accountCreate creates a new account into the keystore defined by the CLI flags.
func accountCreate() {
	keydir := "."
	scryptN := keystore.StandardScryptN
	scryptP := keystore.StandardScryptP

	password := ""

	account, err := keystore.StoreKey(keydir, password, scryptN, scryptP)

	if err != nil {
		log.Fatalf("Failed to create account: %v", err)
	}
	fmt.Printf("\nYour new key was generated\n\n")
	fmt.Printf("Public address of the key:   %s\n", account.Address.Hex())
	fmt.Printf("Path of the secret key file: ./%s\n", account.URL.Path)
}

func main() {
	accountCreate()
}
