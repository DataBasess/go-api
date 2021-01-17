package services

import (
	"bytes"
	"fmt"
	"net"

	"golang.org/x/crypto/ssh"
)

func ExecuteCmd(command, hostname, port, user, passwd string) string {
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(passwd),
		},
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
	}
	conn, connErr := ssh.Dial("tcp", fmt.Sprintf("%s:%s", hostname, port), config)
	if connErr != nil {
		return fmt.Sprintf("Connection Fail")
	}
	session, sessionErr := conn.NewSession()
	if sessionErr != nil {
		return fmt.Sprintf("Session Fail")
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Run(command)

	return fmt.Sprintf("%s -> %s", hostname, stdoutBuf.String())
}
