// This demo shows how to HMAC sign a JSON payload.
// We can inject this into a header.
// Note: Different languages have different standards of

package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
)

type demo struct {
	Foo string `json:"foo,omitempty"`
}

var secret = "CLIENT_SECRET"

func main() {

	data := `{"foo": "bar"}`
	log.Printf("Secret: %s Data: %s\n", secret, data)

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	log.Printf("String Literal Bytes: %v \n", []byte(data))
	log.Println("String Literal Result: " + sha)

	log.Println("Marshalled Result: " + createSignatureFromJSON(json.Marshal(demo{
		Foo: "bar",
	})))

}

func createSignatureFromJSON(b []byte, err error) string {
	if err != nil {
		return ""
	}
	h := hmac.New(sha256.New, []byte(secret))

	log.Printf("Secret: %s Data: %s\n", secret, string(b))
	log.Printf("Marshalled Bytes: %v \n", b)

	// Write Data to it
	h.Write(b)

	// Get result and encode as hexadecimal string
	return hex.EncodeToString(h.Sum(nil))
}
