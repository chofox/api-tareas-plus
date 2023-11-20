package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func generateRandomKey(length int) (string, error) {
	key := make([]byte, length)
	_, err := rand.Read(key)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func key() {
	// Generar una clave secreta de 32 bytes
	key, err := generateRandomKey(32)
	if err != nil {
		fmt.Println("Error al generar la clave secreta:", err)
		return
	}
	fmt.Println("Clave secreta generada:", key)
}
