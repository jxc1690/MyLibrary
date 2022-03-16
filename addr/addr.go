package addr

import "net"

func UrltoIp(url string) ([]net.IP, error) {
	return net.LookupIP(url)
}

func UrltoUdpaddr(url string) (*net.UDPAddr, error) {
	ip, err := UrltoIp(url)
	if err != nil {
		return nil, err
	}
	return IptoUdpaddr(ip[0], 0), nil
}

func IptoUdpaddr(ip net.IP, port int) *net.UDPAddr {
	return &net.UDPAddr{IP: ip, Port: port}
}
func IptoTcpaddr(ip net.IP, port int) *net.TCPAddr {
	return &net.TCPAddr{IP: ip, Port: port}
}

func StrtoTcpaddr(ip string) (*net.TCPAddr, error) {
	return net.ResolveTCPAddr("tcp", ip)
}

func StrtoUcpaddr(ip string) (*net.UDPAddr, error) {
	return net.ResolveUDPAddr("udp", ip)
}
