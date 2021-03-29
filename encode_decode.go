package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
	"log"
	"loso/config"
	"loso/database"
	"strings"
)

//var CNX = database.Connection()
var Conn, _ = database.NewCon(config.Hostmgo, "ln-smt")

type Trainer struct {
	Name string
	Age  int
	City string
}

const password1 = "loso"
const password = "kimera"
const pw1 = "631492b1c37b17aea26fc1a91617f9d27c086b9b22eb2b2cd17c190b2008adf6.1e2a93d51782daf46e811b9e94326d0102b4452a2716540376922380eb49cca4"
const pw2 = "366d7c45c902fc551e3194d551fcddd079173ae845b8438c0b5b04b901af845f.94c315e50ac96f7fe95c8df02e08a7a37c6a89d0a44e00a1077832f2d795246a"
const pw3t = "24a45cd5a2b51338927950a4d4c21f9dcedf49368573b8901b9d8e5adb2aeca3.84f8a20da29b16122f2af267119541373b5b3e204817ff30f7cf832ecd558c8c"
const loso = "c7beeb003131766e750efd83f86e44c19e9a41ec732eeda36eabf37a94a7b079.74ba1a0501ca2784cc51912e3077709310ace416adb9da5bac25196f601bcbeb"

const lnteam = "12523cefb41e299f761ade103ca2a4e8b708c69afc5dc6931b576926fc160632.fcbd654c6a92ae5950c1bc605ec42ab9d3961b68ea7504d4809e3d48c40ddd38"

func main() {

	fmt.Println(hashPassword(password))

	//comparePasswordss(pw1,password)
	fmt.Println(comparePasswordss(lnteam, "lnteam"))

}

func hashPassword(password string) (string, error) {
	// example for making salt - https://play.golang.org/p/_Aw6WeWC42I
	salt := make([]byte, 32)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	// using recommended cost parameters from - https://godoc.org/golang.org/x/crypto/scrypt
	shash, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}

	// return hex-encoded string with salt appended to password
	hashedPW := fmt.Sprintf("%s.%s", hex.EncodeToString(shash), hex.EncodeToString(salt))

	return hashedPW, nil
}

func comparePasswordss(storedPassword string, suppliedPassword string) (bool, error) {
	pwsalt := strings.Split(storedPassword, ".")

	// check supplied password salted with hash
	salt, err := hex.DecodeString(pwsalt[1])

	if err != nil {
		return false, fmt.Errorf("Unable to verify user password")
	}

	shash, err := scrypt.Key([]byte(suppliedPassword), salt, 32768, 8, 1, 32)

	pddws := hex.EncodeToString(shash) == pwsalt[0]

	log.Println("storedPassword:", storedPassword)
	log.Println("Status:", pddws)
	log.Println("salt:", salt)
	log.Println("shash:", shash)
	log.Println("pwsalt:", pwsalt[0])
	log.Println("suppliedPassword:", suppliedPassword)
	log.Println("Hex:", hex.EncodeToString(shash))
	log.Println("-------------------------------------:")

	return hex.EncodeToString(shash) == pwsalt[0], nil
}
