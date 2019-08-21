package dana

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
)

func generateSignature(req Request) (sig string, err error) {
	signer, err := loadPrivateKey()
	if err != nil {
		err = fmt.Errorf("signer is damaged: %v", err)
		return
	}
	plan, err := json.Marshal(req)
	if err != nil {
		err = fmt.Errorf("failed to marshal request: %v", err)
		return
	}
	signed, err := signer.Sign(plan)
	if err != nil {
		err = fmt.Errorf("could not sign request: %v", err)
	}
	sig = base64.StdEncoding.EncodeToString(signed)
	return
}

// loadPublicKey loads an parses a PEM encoded private key file.
func loadPublicKey() (Unsigner, error) {
	return parsePublicKey([]byte(`-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAnaKVGRbin4Wh4KN35OPh
ytJBjYTz7QZKSZjmHfiHxFmulfT87rta+IvGJ0rCBgg+1EtKk1hX8G5gPGJs1htJ
5jHa3/jCk9l+luzjnuT9UVlwJahvzmFw+IoDoM7hIPjsLtnIe04SgYo0tZBpEmkQ
vUGhmHPqYnUGSSMIpDLJDvbyr8gtwluja1SbRphgDCoYVXq+uUJ5HzPS049aaxTS
nfXh/qXuDoB9EzCrgppLDS2ubmk21+dr7WaO/3RFjnwx5ouv6w+iC1XOJKar3CTk
X6JV1OSST1C9sbPGzMHZ8AGB51BM0mok7davD/5irUk+f0C25OgzkwtxAt80dkDo
/QIDAQAB
-----END PUBLIC KEY-----`))
}

// parsePublicKey parses a PEM encoded private key.
func parsePublicKey(pemBytes []byte) (Unsigner, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "PUBLIC KEY":
		rsa, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	return newUnsignerFromKey(rawkey)
}

// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPrivateKey() (Signer, error) {
	return parsePrivateKey([]byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEA4Xa9IgTB/XECPQPFn55sos5apgJIrs3yinqrvq7Yz81rotrI
FU7yWHGoRZDGctobkByMxgVYPrpaoD3m2siduIKwZvLGIMqjqJZH94MSTQPz4HvJ
TCvWYH8X2Mr/SjPHRuG6sV8slMrAubKiLvMIY1iUmK6lm7QzbjMpwetBt0keHV9m
6DuyElZ+VubNhhiq+YgDCm9x7RA69zxaDn0FatSI90cjFva2199St3XIOO2m1D6T
YXQHFDp0mi7uKBPk1QMRlOSaFPwoQLCgR6S5AOLCf7c1icFhpdzBVZfNCeorGO87
jt6F6wdOIf3Qw+tS8cD7GdtNkYpXlKyMZDmFBQIDAQABAoIBAQDSnZ4eej1RaNED
eFOTQWYQTB4n9/g3u1F4BRIMxb7pl+aVsjS4mxAOiNX9bjnDrnWTEOh8Tx/ZMTJZ
gJn4BPI5G3R4Jw/oMAPkB8bgs0NUSH34VQp9KnJPoEoBjdgfU/EzCMjiypqVwDV7
AyE4JHXQgD4HxxYTNqnYpti6Ou8fq9+Z+Z2UnSHNT/KPTXy8x4esNLLjLkhfyEvH
ci7yS9zPgIR8dNSjN9DgDiP25NJnaasbmdPFrWGq2RDSQB/iGeF9BwVQnA0Rrm3L
gniy5ck9sAwE9V1RHn7wzRzqgoOi/y8vBo6kNiLQbiJ4NkA6BCNRqb2Y481+wFW9
krGJ6ojBAoGBAPx15Mys8u/BOe4YR+BIwd6618O0YqXXcFjAf8yLWqGGrh5ZhprL
ooRp0nHZKgIYi3gM5gBY3q3ahpZM3mmRp7XLn/kLUAl8/gH/ckPdQBYR4+nckAt4
psjz+He2yOLQtUVomgYRqgCDwbnshcIjF0rsaFBb8Z/LDzgEQf31wlGxAoGBAOSf
84B+OEcZzY/CLVJVCO9hKfoMSU5wxYpgdd2ojldZYhvDscAJ4IWTqMekCgrlSPps
nmNoUeCfrO/L38tzpTwRG7tRtodlN9HWaWSGu9jD4Mtk7fna/sFcA7vzRSCaSFYr
A0wqrbg8hQWsovZg0LJ818b8WJ9mo9mJJZfqvMmVAoGAchIi64hJzKMWNcqqz2CQ
mL77tuXOnPiXPsLb++QLc2iTTAtPkqnoKMRX3jEtlfzZJ8lMM2P0WcGfhlH9PN4f
VTlIcxL1exQLAYQWcSuLRW2X+Zc7TrOZSLtTzWO9qR58iXwV/CAHHGxsIEIKvgl1
ANmI2KZyqqzpl3n1B3b/8kECgYEAjge4dJj1SWaER8zv7vgY5u3L8CV13R89+WtX
je7LDwTeaDU41f7M1u8WqYAFJSOwJLMZ6tPCGfPYsdhMHc6oVhqiycpHxPFirh+v
td3mERIPMwxuswwTu/f38el9hnWSfsWV9NmGtrASS6YmTz1yKZBZVKssO0q+e8Qx
8KDJy8kCgYBcLjkZs4TKVqwGes8YlwwWvGCl7z477cQHkU5jmNJREl0pcMl7ZimX
m658eBoq5KOzGECD0kfsD59b8LdYL6xRDwQczv4B2araIYJN495Jz23FBnKuI3Eg
d59SdOyF8W/z7k8LSN/t+eKBjZPf3ifCrye20Q7UFV+z4e9IXIWetg==
-----END RSA PRIVATE KEY-----`))
}

// parsePrivateKey parses a PEM encoded private key.
func parsePrivateKey(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
	return newSignerFromKey(rawkey)
}

// A Signer is can create signatures that verify against a public key.
type Signer interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Sign(data []byte) ([]byte, error)
}

// A Signer is can create signatures that verify against a public key.
type Unsigner interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Unsign(data []byte, sig []byte) error
}

func newSignerFromKey(k interface{}) (Signer, error) {
	var sshKey Signer
	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

func newUnsignerFromKey(k interface{}) (Unsigner, error) {
	var sshKey Unsigner
	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

type rsaPublicKey struct {
	*rsa.PublicKey
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

// Sign signs data with rsa-sha256
func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha256.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, d)
}

// Unsign verifies the message using a rsa-sha256 signature
func (r *rsaPublicKey) Unsign(message []byte, sig []byte) error {
	h := sha256.New()
	h.Write(message)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, d, sig)
}
