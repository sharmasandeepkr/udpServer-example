package main
import (
	"fmt"
	"net"
	"os"
)

func main(){
	service:= ":1200"
	udpAddr, err:= net.ResolveUDPAddr("udp4", service)
	if err!= nil {
		os.Exit(1)
	}
	udpCon, err:= net.DialUDP("udp", nil, udpAddr)
	if err!= nil {
		os.Exit(2)
	}
	_, err= udpCon.Write([]byte("bye"))
	if err!= nil {
		os.Exit(3)
	}
	buf:= make([]byte, 4096)
	n, err:= udpCon.Read(buf[0:])
	if err!= nil {
		os.Exit(4)
	}
	fmt.Println(string(buf[0:n]))
}