package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {

	b64, _ := hexToBase64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")
	fmt.Println(b64)

	xor, _ := xorString("1c0111001f010100061a024b53535009181c", "686974207468652062756c6c277320657965")
	fmt.Println(xor)

	bestResult, _ := singleByteXorBruteForce("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	fmt.Println(bestResult)
}

func hexToBase64(hexString string) (string, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return "Failed to decode String with Error - ", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

func xorString(a, b string) (string, error) {
	if len(a) != len(b) {
		return "Strings are not of equal length", nil
	}
	bytesOfA, err := hex.DecodeString(a)
	if err != nil {
		return "Failed to decode String A with Error - ", err
	}

	bytesOfB, err := hex.DecodeString(b)
	if err != nil {
		return "Failed to decode String B with Error - ", err
	}

	result := make([]byte, len(bytesOfA))
	for i := 0; i < len(bytesOfA); i++ {
		result[i] = bytesOfA[i] ^ bytesOfB[i]
	}
	return hex.EncodeToString(result), nil
}

func singleByteXorBruteForce(input string) (string, error) {
	bytes, err := hex.DecodeString(input)
	if err != nil {
		return "Failed to decode String with Error - ", err
	}
	for k := 0; k < 256; k++ {
		result := make([]byte, len(bytes))
		for i := 0; i < len(bytes); i++ {
			result[i] = bytes[i] ^ byte(k)
		}
		fmt.Println(string(result))
	}
	return "", nil
}
