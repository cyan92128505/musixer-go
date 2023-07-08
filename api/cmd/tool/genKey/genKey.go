package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println("無法生成私鑰:", err)
		return
	}

	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	})

	privateKeyFile, err := os.Create("id_rsa_priv.pem")
	if err != nil {
		fmt.Println("無法建立私鑰檔案:", err)
		return
	}
	defer privateKeyFile.Close()

	_, err = privateKeyFile.Write(privateKeyPEM)
	if err != nil {
		fmt.Println("無法寫入私鑰檔案:", err)
		return
	}

	publicKey := privateKey.PublicKey
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		fmt.Println("無法序列化公鑰:", err)
		return
	}

	publicKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicKeyBytes,
	})

	publicKeyFile, err := os.Create("id_rsa_pub.pem")
	if err != nil {
		fmt.Println("無法建立公鑰檔案:", err)
		return
	}
	defer publicKeyFile.Close()

	_, err = publicKeyFile.Write(publicKeyPEM)
	if err != nil {
		fmt.Println("無法寫入公鑰檔案:", err)
		return
	}

	fmt.Println("已成功生成 RSA 公私鑰檔案")

	// 讀取私鑰檔案
	privKeyBytes, err := ioutil.ReadFile("id_rsa_priv.pem")
	if err != nil {
		fmt.Println("讀取私鑰檔案時發生錯誤:", err)
		return
	}

	// 讀取公鑰檔案
	pubKeyBytes, err := ioutil.ReadFile("id_rsa_pub.pem")
	if err != nil {
		fmt.Println("讀取公鑰檔案時發生錯誤:", err)
		return
	}

	// 將私鑰和公鑰轉換為 base64 格式
	privKeyBase64 := base64.StdEncoding.EncodeToString(privKeyBytes)
	pubKeyBase64 := base64.StdEncoding.EncodeToString(pubKeyBytes)

	// 合併私鑰和公鑰為一個文字檔案
	output := fmt.Sprintf("ACCESS_TOKEN_PRIVATE_KEY=%s\nACCESS_TOKEN_PUBLIC_KEY=%s", privKeyBase64, pubKeyBase64)

	// 儲存文字檔案
	err = ioutil.WriteFile("keys.env", []byte(output), 0644)
	if err != nil {
		fmt.Println("儲存文字檔案時發生錯誤:", err)
		return
	}

	fmt.Println("已成功轉換並儲存 base64 文字檔案")
}
