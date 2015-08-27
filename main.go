package main

import (
	"fmt"
    "log"

    "github.com/go-ldap/ldap"
)

type userConfig struct {
	port     int
	address  string
	password string
	base_dn  string
}

func NewUser(port int, address string, password string, base_dn string) *userConfig {
	user := new(userConfig)
	user.address = address
	user.port = port
	user.password = password
	user.base_dn = base_dn

	return user
}

func Bind(user *userConfig) error {
	l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", user.address, user.port))
	if err != nil {
		return err
	}
	defer l.Close()

	err = l.Bind(user.base_dn, user.password)
	if err != nil {
	    return err
	}

    return nil
}

func main() {
	user := &userConfig{389,
		"1.2.3.4",
		"EderBaitola",
		"uid=eder,ou=muito,dc=baitola,dc=com"}

    err := Bind(user)
    if err != nil {
        log.Fatal(err)
    }
}
