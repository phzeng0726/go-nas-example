package main

import (
	"fmt"
	"net"
	"os"

	"github.com/hirochachacha/go-smb2"
	_ "github.com/joho/godotenv/autoload"
)

var (
	nasAddr     string = fmt.Sprintf("%s:445", os.Getenv("SERVER_NAME"))
	nasUser     string = os.Getenv("USER")
	nasPassword string = os.Getenv("PASSWORD")
)

func main() {
	conn, err := net.Dial("tcp", nasAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	d := &smb2.Dialer{
		Initiator: &smb2.NTLMInitiator{
			User:     nasUser,
			Password: nasPassword,
		},
	}

	s, err := d.Dial(conn)
	if err != nil {
		panic(err)
	}
	defer s.Logoff()

	names, err := s.ListSharenames()
	if err != nil {
		panic(err)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
