package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func main() {
	// Generate a new key pair
	priv, err := sm2.GenerateKey(rand.Reader)
	if err != nil {
		panic(err)
	}

	// ==========【1、私钥导出PEM】==========
	// x509序列化SM2私钥（ASN1编码）
	privDer, err := x509.MarshalSm2PrivateKey(priv, nil)
	if err != nil {
		panic(err)
	}
	// PEM封装
	privBlock := &pem.Block{
		Type:  "SM2 PRIVATE KEY",
		Bytes: privDer,
	}
	privPem := pem.EncodeToMemory(privBlock)
	fmt.Println("===导出私钥PEM===\n", string(privPem))

	// ==========【1、私钥导出极简模式】==========
	hexProvKey := PrivToHexD(priv)
	fmt.Println("===导出极简 Private Key===\n", hexProvKey)
	pubKey2Str := PublicToHexD(&priv.PublicKey)
	fmt.Println("===导出极简 public Key===\n", pubKey2Str)

	// Plain text
	msg := []byte("Hello, SM2 encryption!")

	// ==========【2. 通过上面的导出内容获得私钥】==========

	priv2, err := HexDToPriv(hexProvKey)
	if err != nil {
		panic(err)
	}

	// Encrypt using the public key
	cipher, err := sm2.EncryptAsn1(&priv2.PublicKey, msg, rand.Reader)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Ciphertext (hex): %x\n", cipher)

	// Decrypt using the private key
	plain, err := sm2.DecryptAsn1(priv2, cipher)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Decrypted message: %s\n", plain)

	//  加解密2
	fmt.Println("===加解密2-========\n")

	// Plain text
	msg2 := []byte("Hello, SM2 encryption!222222")
	pubKey2, err := XYHexToPub(pubKey2Str)
	if err != nil {
		panic(err)
	}
	enMsg2, err := sm2.Encrypt(pubKey2, msg2, rand.Reader, sm2.C1C3C2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ciphertext (hex): %x\n", enMsg2)

	deMsg2, err := sm2.Decrypt(priv2, enMsg2, sm2.C1C3C2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Decrypted message: %s\n", string(deMsg2))

}

// 1. SM2私钥 → 64位HEX(D值)配置字符串（最短！推荐配置用）
func PrivToHexD(priv *sm2.PrivateKey) string {
	return hex.EncodeToString(priv.D.Bytes())
}

// 1. SM2私钥 → 64位HEX(D值)配置字符串（最短！推荐配置用）
func PublicToHexD(pub *sm2.PublicKey) string {

	xBytes := pub.X.Bytes()
	yBytes := pub.Y.Bytes()
	return hex.EncodeToString(append(xBytes, yBytes...))
}

// HEX字符串还原公钥
func XYHexToPub(hexStr string) (*sm2.PublicKey, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	pub := new(sm2.PublicKey)
	pub.Curve = sm2.P256Sm2()
	pub.X = new(big.Int).SetBytes(buf[:32])
	pub.Y = new(big.Int).SetBytes(buf[32:])
	return pub, nil
}

// 2. 64位HEX(D) → 还原*sm2.PrivateKey
func HexDToPriv(hexStr string) (*sm2.PrivateKey, error) {
	buf, err := hex.DecodeString(hexStr)
	if err != nil {
		return nil, err
	}
	priv := new(sm2.PrivateKey)
	priv.D = new(big.Int).SetBytes(buf)
	// 自动从D反算公钥XY（SM2算法特性，不用存公钥）
	priv.PublicKey.Curve = sm2.P256Sm2()
	priv.PublicKey.X, priv.PublicKey.Y = sm2.P256Sm2().ScalarBaseMult(buf)

	return priv, nil
}

func pemTextToPriv(privConfStr string) {

	//================= 二、从配置字符串反向加载私钥 =================
	// base64解码
	derBuf, _ := base64.StdEncoding.DecodeString(privConfStr)
	loadPriv, _ := x509.ParseSm2PrivateKey(derBuf)
	fmt.Printf("\n从配置串加载私钥成功：%t\n", loadPriv != nil)

	// 公钥加载
	// pubBuf, _ := base64.StdEncoding.DecodeString(pubConfStr)
	// loadPub, _ := x509.ParseSm2PublicKey(pubBuf)
	// fmt.Printf("从配置串加载公钥成功：%t\n", loadPub != nil)

}
