package configparser

import (
	"fmt"
	"io"
	"os"

	"github.com/lukirs95/strongswan-user-api/pkg/strongswanuser"
)

const userFormat string = "%s %s %s %s\n"

func serializeOneUser(username *string, password *string) string {
	return fmt.Sprintf("%s : EAP \"%s\"\n", *username, *password)
}

func Serialize(file *os.File, userList *strongswanuser.List) error {
	file.Seek(0, io.SeekStart)
	err := goToParseBegin(file, "#beginparsing")
	if err != nil {
		panic(err)
	}
	offset, _ := file.Seek(0, io.SeekCurrent)
	file.Seek(0, io.SeekStart)
	file.Truncate(offset)
	file.Seek(offset, io.SeekStart)
	serialized := ""
	for _, user := range userList.Users() {
		serialized = fmt.Sprintf("%s%s", serialized, serializeOneUser(&user.Username, &user.Password))
	}
	_, err = file.Write([]byte(serialized))
	return err
}

func trimStringFirstLast(s string) string {
	return s[1 : len(s)-1]
}

func goToParseBegin(file *os.File, identifier string) error {
	for {
		line := ""
		_, err := fmt.Fscanln(file, &line)
		if err != nil {
			switch err {
			case io.EOF:
				return fmt.Errorf("parsing error: no \"%s\" found", identifier)
			default:
				continue
			}
		}
		if line == identifier {
			return nil
		}
	}
}

// deserialize ipsec.secrets
func Deserialize(file *os.File) (*strongswanuser.List, error) {
	userList := strongswanuser.NewList()
	if file == nil {
		return userList, fmt.Errorf("Deserialize Error: File is nil-pointer")
	}
	err := goToParseBegin(file, "#beginparsing")
	if err != nil {
		return userList, err
	}

	for {
		username := ""
		password := ""
		a := ""
		_, err := fmt.Fscanf(file, userFormat, &username, &a, &a, &password)
		if err != nil {
			switch err {
			case io.EOF:
				return userList, nil
			default:
				return userList, err
			}
		}
		password = trimStringFirstLast(password)
		userList.Append(username, password)
	}
}
