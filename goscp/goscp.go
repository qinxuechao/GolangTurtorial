package main

import (
	"flag"
	"fmt"
	scp "github.com/bramvdbogaerde/go-scp"
	"github.com/bramvdbogaerde/go-scp/auth"
	"golang.org/x/crypto/ssh"
	"os"
)

func main() {
	// Use SSH key authentication from the auth package
	// we ignore the host key in this example, please change this if you use this library
	clientConfig, _ := auth.PasswordKey("root", "Huawei2020", ssh.InsecureIgnoreHostKey())

	// For other authentication methods see ssh.ClientConfig and ssh.AuthMethod
	server := flag.String("host", "", "eg: 192.168.82.65")


	flag.Parse()
	// Create a new SCP client
	host := server + ":22"
	client := scp.NewClient("192.168.90.74:22", &clientConfig)

	// Connect to the remote server
	err := client.Connect()
	if err != nil {
		fmt.Println("Couldn't establish a connection to the remote server ", err)
		return
	}

	// Open a file
	f, _ := os.Open("/Users/qinxuechao/Desktop/pb_reader.py")

	// Close client connection after the file has been copied
	defer client.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	err = client.CopyFile(f, "/root/config/pb_reader.py", "0655")

	if err != nil {
		fmt.Println("Error while copying file ", err)
	}
	fmt.Println("Transferred!")
}