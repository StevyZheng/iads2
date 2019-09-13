package hardware

import (
	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/vishvananda/netlink"
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

func (e *NetInfo) NetInit() (err error) {
	_, err = netlink.LinkList()
	return
}
