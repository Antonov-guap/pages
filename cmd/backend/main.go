package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/samber/lo"
)

func main() {
	app := fiber.New()

	priv := lo.Must(rsa.GenerateKey(rand.Reader, 2048))

	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Your Organization"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 365),
		KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageServerAuth,
		},
		BasicConstraintsValid: true,
	}

	certDER := lo.Must(x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv))
	tlsCert := tls.Certificate{
		Certificate: [][]byte{certDER},
		PrivateKey:  priv,
	}

	app.Get("/time", func(c *fiber.Ctx) error {
		return c.SendString(time.Now().Format(time.RFC3339))
	})

	lo.Must0(app.ListenTLSWithCertificate("0.0.0.0:8080", tlsCert))
}
