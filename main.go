package main

import (
	"encoding/base64"
	"encoding/json"
	"errors"
)

type Header struct {
	Algorithm string
}

func hs(a string) string {
	return "snjc"
}

const algorithmFunc map[string]func(string) {
	"HS256": hs
}
// header.payload.signature

func encodeHeader(options Header) (string, error) {
	var byteHeader, err = json.Marshal(options)

	if err != nil {
		return "", errors.New("unable to encode header")
	}

	var encodedHeader string = base64.StdEncoding.EncodeToString(byteHeader)
	return encodedHeader, nil
}

func sign(payload []byte, secret string, options Header) (string, error) {
	// var payloadMap map[string]interface{}
	// json.Unmarshal(payload, &payloadMap)
	// fmt.Printf("%#v\n", payloadMap)

	if options.Algorithm == "" {
		options.Algorithm = "HS256"
	}

	encodedPayload := base64.StdEncoding.EncodeToString(payload)
	encodedHeader, err := encodeHeader(options)

	if err != nil {
		return "", errors.New("unable to generate token")
	}

	var token string = encodedHeader + "." + encodedPayload + "."

	return token, nil
}

func main() {
	var secret string = "abcdef"
	payload := []byte(`{"course": "testcourse","price": 100}`)

	sign(payload, secret, Header{})
}
