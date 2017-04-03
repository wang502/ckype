package encryption

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
)

// encrypt ckype message using RSA
func rsaEncrypt(message, dir string) ([]byte, error) {
	var res []byte

	pemBytes, err := readPemFile(dir)
	if err != nil {
		return res, err
	}

	pub, err := x509.ParsePKIXPublicKey(pemBytes)
	if err != nil {
		return res, err
	}

	secretMessage := []byte(message)
	label := []byte("orders")

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader
	return rsa.EncryptOAEP(sha256.New(), rng, pub.(*rsa.PublicKey), secretMessage, label)
}

// decrypt ckype message
func rsaDecrpt(ciphertext []byte, dir string) (string, error) {
	label := []byte("orders")

	pemBytes, err := readPemFile(dir)
	if err != nil {
		return "", err
	}

	private, err := x509.ParsePKCS1PrivateKey(pemBytes)
	if err != nil {
		return "", err
	}

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, private /*.(*rsa.PrivateKey)*/, ciphertext, label)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}

// calculates the RSA signature of of hashed message
func sign(message, dir string) ([]byte, error) {
	var signature []byte

	pemBytes, err := readPemFile(dir)
	if err != nil {
		return signature, err
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(pemBytes)
	if err != nil {
		return signature, err
	}

	hashed := sha256.Sum256([]byte(message))
	rng := rand.Reader
	return rsa.SignPKCS1v15(rng, privateKey, crypto.SHA256, hashed[:])
}

// verify the RSA signature of hashed message
func verify(message [32]byte, signature []byte, dir string) error {
	pemBytes, err := readPemFile(dir)
	if err != nil {
		return err
	}

	publicKey, err := x509.ParsePKIXPublicKey(pemBytes)
	if err != nil {
		return err
	}

	return rsa.VerifyPKCS1v15(publicKey.(*rsa.PublicKey), crypto.SHA256, message[:], signature)
}
