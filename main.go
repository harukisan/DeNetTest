package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"io/ioutil"
	"log"
	"os"
)

//Создать кошелек

func CreateAccount(pass string) {

	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	_, err := key.NewAccount(pass)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	var pass string
	var answerTrue int32 = 'y'
	var answer string

	files, _ := os.ReadDir("./wallet")
	if len(files) == 0 {
		fmt.Println("Кошелька нет, хотите создать?[y/n]")
		fmt.Scan(&answer)
		if answer == string(answerTrue) {
			fmt.Println("Введите пароль от кошелька: ")
			fmt.Scan(&pass)
			CreateAccount(pass)
			fmt.Println("Аккаунт успешно создан!")
			os.Exit(1)
		} else {
			fmt.Println("Пока..")
		}

	} else {

		for _, file := range files {

			fmt.Println("Введите пароль от аккаунта:")
			fmt.Scan(&pass)

			b, err := ioutil.ReadFile("./wallet/" + file.Name())
			if err != nil {
				log.Fatal("Error: Cannot open a wallet..")
			}

			key, err := keystore.DecryptKey(b, pass)
			if err != nil {
				log.Fatal("Wrong Password!")
			}

			pData := crypto.FromECDSA(key.PrivateKey)
			fmt.Println("Private key:", hexutil.Encode(pData))
			pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
			fmt.Println("Public key:", hexutil.Encode(pubData))
			fmt.Println("Address:", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())
		}
	}
}
