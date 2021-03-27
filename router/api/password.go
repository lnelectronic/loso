// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 27/3/2564 18:51
// ---------------------------------------------------------------------------
package api

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/scrypt"
)

func hashPassword(Passwd string) (string, error) {

	// example for making salt - https://play.golang.org/p/_Aw6WeWC42I
	salt := make([]byte, 32)
	//var _, err = bcrypt.GenerateFromPassword([]byte(Passwd), 32)
	var _, err = rand.Read(salt)
	if err != nil {
		return "", err
	}

	// using recommended cost parameters from - https://godoc.org/golang.org/x/crypto/scrypt
	hash, err := scrypt.Key([]byte(Passwd), salt, 32768, 8, 1, 32)
	if err != nil {
		//return "", err
		return "", err
	}

	// return hex-encoded string with salt appended to password
	encoded := fmt.Sprintf("%s.%s", hex.EncodeToString(hash), hex.EncodeToString(salt))

	return encoded, nil
}
