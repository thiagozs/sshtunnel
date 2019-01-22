package sshtunnel

import (
	"io/ioutil"

	"golang.org/x/crypto/ssh"
)

// PrivateKeyFile ssh auth method
func PrivateKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}

	return ssh.PublicKeys(key)
}
