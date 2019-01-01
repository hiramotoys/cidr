package cidr

import (
	"fmt"
	"strconv"
	"strings"
)

type CidrBlock struct {
	ipAddress        uint32
	subnetMask       uint32
	networkAddress   uint32
	broadcastAddress uint32
	IpAddressRange   [][4]uint8
}

func NewCidr(cidr string) (*CidrBlock, error) {
	ipAddressStr, subnetMaskStr, err := parseCmmandLineInput(cidr)
	ipAddress := convertIpv4StrInto32bitInteger(ipAddressStr)
	subnetMask, _ := subnetMask32bitInteger(subnetMaskStr)
	networkAddress := getNetworkAddress(ipAddress, subnetMask)
	broadcastAddress := getBroadcastAddress(ipAddress, subnetMask)
	ipAddressRange := getIpAdressRange(networkAddress, broadcastAddress)
	cidrBlock := CidrBlock{ipAddress, subnetMask, networkAddress, broadcastAddress, ipAddressRange}
	return &cidrBlock, err
}

func (cb *CidrBlock) Print() int {
	ipAddressArray := convert32bitIntegerIntoIpv4Array(cb.ipAddress)
	fmt.Printf("IP address: %d.%d.%d.%d\n", ipAddressArray[0], ipAddressArray[1], ipAddressArray[2], ipAddressArray[3])
	subnetMaskArray := convert32bitIntegerIntoIpv4Array(cb.subnetMask)
	fmt.Printf("Subnet mask: %d.%d.%d.%d\n", subnetMaskArray[0], subnetMaskArray[1], subnetMaskArray[2], subnetMaskArray[3])
	networkAddressArray := convert32bitIntegerIntoIpv4Array(cb.networkAddress)
	fmt.Printf("Network address: %d.%d.%d.%d\n", networkAddressArray[0], networkAddressArray[1], networkAddressArray[2], networkAddressArray[3])
	broadcastAddressArray := convert32bitIntegerIntoIpv4Array(cb.broadcastAddress)
	fmt.Printf("Broadcast address: %d.%d.%d.%d\n", broadcastAddressArray[0], broadcastAddressArray[1], broadcastAddressArray[2], broadcastAddressArray[3])
	fmt.Printf("IP address range:")
	for i := 0; i < len(cb.IpAddressRange); i++ {
		fmt.Printf(" %d.%d.%d.%d", cb.IpAddressRange[i][0], cb.IpAddressRange[i][1], cb.IpAddressRange[i][2], cb.IpAddressRange[i][3])
	}
	fmt.Printf("\n")
	return 0
}

func parseCmmandLineInput(cidr string) (string, string, error) {
	ret := strings.Split(cidr, "/")
	return ret[0], ret[1], nil
}

func getNetworkAddress(ipaddress, subnetmask uint32) uint32 {
	return ipaddress & subnetmask
}

func getBroadcastAddress(ipaddress, subnetmask uint32) uint32 {
	return ipaddress | (^subnetmask)
}

func getIpAdressRange(networkAddress, broadcastAddress uint32) [][4]uint8 {
	num := broadcastAddress - networkAddress + 1
	result := make([][4]uint8, num)
	for i := 0; i < int(num); i++ {
		result[i] = convert32bitIntegerIntoIpv4Array(networkAddress + uint32(i))
	}
	return result
}

func convertIpv4StrInto32bitInteger(ipv4Str string) uint32 {
	ipv4Array, _ := convertIpv4StrIntoIpv4Array(ipv4Str)
	return convertIpv4ArrayInto32bitInteger(ipv4Array)
}

func convertIpv4StrIntoIpv4Array(ipv4 string) ([4]uint8, error) {
	ret := strings.Split(ipv4, ".")
	var ipv4Array [4]uint8
	for i := 0; i < 4; i++ {
		v, _ := strconv.ParseUint(ret[i], 10, 8)
		ipv4Array[i] = uint8(v)
	}
	return ipv4Array, nil
}

func convertIpv4ArrayInto32bitInteger(ipv4 [4]uint8) uint32 {
	var result uint32
	result = 0
	for i := 0; i < 4; i++ {
		result = result | (uint32(ipv4[i]) << uint32((4-i-1)*8))
	}
	return result
}

func convert32bitIntegerIntoIpv4Array(itg uint32) [4]uint8 {
	var result [4]uint8
	for i := 0; i < 4; i++ {
		shift_num := uint32((4 - i - 1) * 8)
		result[i] = uint8(itg & uint32(255<<shift_num) >> shift_num)
	}
	return result
}

func subnetMask32bitInteger(subnetMask string) (uint32, error) {
	var ipv4Array [4]uint8
	for i := 0; i < 4; i++ {
		b := ""
		for j := 0; j < 8; j++ {
			k := j + (i * 8)
			v := "0"
			if s, _ := strconv.Atoi(subnetMask); k/s == 0 {
				v = "1"
			}
			b += v
		}
		r, _ := strconv.ParseUint(b, 2, 8)
		ipv4Array[i] = uint8(r)
	}
	return convertIpv4ArrayInto32bitInteger(ipv4Array), nil
}
