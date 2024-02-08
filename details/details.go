package details

import (
	"os"
	"net"
	"log"
)

func GetHostName() (string, error) {
	hostName, err := os.Hostname()
	return hostName, err
}

func GetIP() (net.IP , error){
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP, err
}