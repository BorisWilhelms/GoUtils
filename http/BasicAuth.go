package BasicAuth

import (
	"encoding/base64"
	"fmt"
)

func Create(userName string, password string) string {
	data := []byte(fmt.Sprintf("%s:%s", userName, password))
	return base64.StdEncoding.EncodeToString(data)
}
