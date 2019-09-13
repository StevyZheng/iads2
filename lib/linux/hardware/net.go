package hardware

import (
	"fmt"
	"github.com/emirpasic/gods/lists/arraylist"
	"log"
	"net"
)

type NetInfo struct {
	NetInterface struct {
		Ipaddr     string
		Netmask    string
		Gateway    string
		Mac        string
		Dns        string
		LinkStatus bool
	}
	Interfaces arraylist.List
}

func (e *NetInfo) Init() error {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		fmt.Println(addr)
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}
	return err
}
