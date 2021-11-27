package main

import (
	"bufio"
	"crypto/rand"
	"fmt"
	"image/png"
	"os"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
)

func main() {
	ops := totp.GenerateOpts{
		Issuer:      "cateiru.com",
		AccountName: "Yuto Watanabe",
		Period:      30,
		SecretSize:  20,
		Secret:      []byte("hogehoge"),
		Digits:      otp.DigitsSix,
		Algorithm:   otp.AlgorithmSHA1,
		Rand:        rand.Reader,
	}

	key, err := totp.Generate(ops)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("Token: %s", key.String())
	logrus.Infof("Secret: %x", key.Secret())

	image, err := key.Image(255, 255)
	if err != nil {
		logrus.Fatal(err)
	}

	f, err := os.Create("qr.png")
	if err != nil {
		logrus.Fatal(err)
	}
	if err := png.Encode(f, image); err != nil {
		logrus.Fatal(err)
	}

	fmt.Print("Passcode?: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	passcode := scanner.Text()

	logrus.Infof("Validate passcode. result: %v", totp.Validate(passcode, key.Secret()))
}
