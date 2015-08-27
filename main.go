package main

import (
	"fmt"
	"log"

	"github.com/go-ldap/ldap"
)

type userConfig struct {
	username string
	password string
}

type ldapConfig struct {
	base_dn string
	port    int
	address string
}

func (l *ldapConfig) Bind(user *userConfig) (*ldap.Conn, error) {
	conn, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", l.address, l.port))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.Bind(l.base_dn, user.password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func main() {

	user := &userConfig{"eder", "EderBaitola"}
	l := &ldapConfig{"base_dn", 389, "127.0.0.1"}

	_, err := l.Bind(user)
	if err != nil {
		log.Fatal(err)
	}
}
