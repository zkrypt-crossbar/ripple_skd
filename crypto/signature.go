package crypto

import (
	"crypto/ed25519"
	"fmt"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

func Sign(privateKey, hash, msg []byte) ([]byte, error) {
	switch len(privateKey) {
	case ed25519.PrivateKeySize:
		return signEd25519(privateKey, msg)
	case btcec.PrivKeyBytesLen:
		return signECDSA(privateKey, hash)
	default:
		return nil, fmt.Errorf("Unknown private key format")
	}
}

func Verify(publicKey, hash, msg, signature []byte) (bool, error) {
	switch publicKey[0] {
	case 0xED:
		return verifyEd25519(publicKey, signature, msg)
	case 0x02, 0x03:
		return verifyECDSA(publicKey, signature, hash)
	default:
		return false, fmt.Errorf("Unknown public key format")
	}
}

func SignEd25519(privateKey, msg []byte) ([]byte, error) {
	return ed25519.Sign(privateKey, msg)[:], nil
}

func verifyEd25519(pubKey, signature, msg []byte) (bool, error) {
	switch {
	case len(pubKey) != ed25519.PublicKeySize+1:
		return false, fmt.Errorf("Wrong public key length: %d", len(pubKey))
	case pubKey[0] != 0xED:
		return false, fmt.Errorf("Wrong public format:")
	case len(signature) != ed25519.SignatureSize:
		return false, fmt.Errorf("Wrong Signature length: %d", len(signature))
	default:
		return ed25519.Verify(pubKey[1:], msg, signature), nil
	}
}

// Returns DER encoded signature from input hash
func SignECDSA(privateKey, hash []byte) ([]byte, error) {
	priv, _ := btcec.PrivKeyFromBytes(privateKey)
	sig := ecdsa.Sign(priv, hash)
	return sig.Serialize(), nil
}

// Verifies a hash using DER encoded signature
func verifyECDSA(pubKey, signature, hash []byte) (bool, error) {
	sig, err := ecdsa.ParseDERSignature(signature)
	if err != nil {
		return false, err
	}
	pk, err := btcec.ParsePubKey(pubKey)
	if err != nil {
		return false, nil
	}
	return sig.Verify(hash, pk), nil
}
