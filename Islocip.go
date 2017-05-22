package misc

import "net"

func Islocip(ip string, ips *[]net.Addr) bool {

	if *ips == nil {
		*ips,_=net.InterfaceAddrs()

	}

	for _,p:=range *ips {
		ipnet, _ := p.(*net.IPNet)
		if !IsNotEmpty(ip){continue}
		if ip == ipnet.IP.To4().String() {
			return true
		}
	}
	return false
}