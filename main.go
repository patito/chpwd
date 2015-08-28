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

	err = conn.Bind(l.base_dn, user.password)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (l *ldapConfig) ChangePassword(conn *ldap.Conn, user *userConfig, newpwd string) error {

	passwordModifyRequest := ldap.NewPasswordModifyRequest("", user.password, newpwd)
	_, err := conn.PasswordModify(passwordModifyRequest)

	if err != nil {
		return err
	}

	return nil
}

func main() {

    user := &userConfig{"eder", "EderBaitola"}
    l := &ldapConfig{"base_dn", 389, "127.0.0.1"}

	conn, err := l.Bind(user)
	if err != nil {
		log.Fatal(err)
	}

	err = l.ChangePassword(conn, user, "xuxa")
	if err != nil {
		log.Fatal(err)
	}
	conn.Close()
}
