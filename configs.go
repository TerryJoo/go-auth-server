package main

import (
	"crypto"
	"golang.org/x/crypto/bcrypt"
)

const (
	HashAsm    = crypto.SHA512
	BCryptCost = bcrypt.DefaultCost
	SecretKey  = "SECRET_KEY" // TOOD: Must change this value
)
