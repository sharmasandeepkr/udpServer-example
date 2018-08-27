package main
import (
	"time"
	"log"
	"net"
	"fmt"
	"os"
)


func main(){
	if len(os.Args)!= 2 {
		fmt.Println("Give port to run UDP server")
		os.Exit(3)
	}
	service:= os.Args[1]
	udpAddr, errCon:= net.ResolveUDPAddr("udp4", service)
	if errCon != nil {
		panic(errCon)
	}
	
	udpCon, errCon:= net.ListenUDP("udp", udpAddr)
		if errCon!= nil {
			log.Println(errCon)
			os.Exit(2)
		}
		for {
			handleUdp(udpCon)
		}
}

func handleUdp(udpCon *net.UDPConn){
	buf:= make([]byte, 4096)
	n, addr, err:= udpCon.ReadFromUDP(buf[0:])
	if err!= nil {
		return
	}
	fmt.Println(string(buf[0:n]))
	daytime:= time.Now().String()
	udpCon.WriteToUDP([]byte(daytime), addr)

}
