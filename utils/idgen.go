package utils

import gonanoid "github.com/matoous/go-nanoid/v2"

// alphabet uses URL-safe characters, easy to read and share.
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// NewID generates a short, human-friendly unique ID (12 chars).
// Example: "aB3xZ9mK2pLq"
func NewID() string {
	id, err := gonanoid.Generate(alphabet, 12)
	if err != nil {
		// fallback: should never happen
		panic("idgen: " + err.Error())
	}
	return id
}
