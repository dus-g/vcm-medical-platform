package utils

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func GenerateOTP() string {
	otp := ""
	for i := 0; i < 6; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(10))
		otp += fmt.Sprintf("%d", num.Int64())
	}
	return otp
}
