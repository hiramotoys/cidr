package cidr

import (
	"fmt"
	"strconv"
	"strings"
)

type CidrBlock struct {
	IpAddressBase10        [4]uint8
	SubnetMaskBase10       [4]uint8
	NetworkAddressBase10   [4]uint8
	IpAdressRangeBase10    [][4]uint8
	BroadcastAddressBase10 [4]uint8
}

func NewCidr(cidr string) (*CidrBlock, error) {
	ipAddress, subnetMask, err := parseCidrStr(cidr)
	ipAdressBase10, _ := convertIntoIpv4Array(ipAddress)
	subnetMaskBase10, _ := convertIntoIpv4SubnetMaskArray(subnetMask)
	networkAddressBase10, _ := networkAddress(ipAdressBase10, subnetMaskBase10)
	broadcastAddressBase10, _ := broadcastAddress(ipAdressBase10, subnetMaskBase10)
	ipAddressRangeBase10 := ipAdressRange(networkAddressBase10, broadcastAddressBase10)
	cidrBlock := CidrBlock{ipAdressBase10, subnetMaskBase10, networkAddressBase10, ipAddressRangeBase10, broadcastAddressBase10}
	return &cidrBlock, err
}

func (cb *CidrBlock) Print() int {
	fmt.Printf("IP address: %d.%d.%d.%d\n", cb.IpAddressBase10[0], cb.IpAddressBase10[1], cb.IpAddressBase10[2], cb.IpAddressBase10[3])
	fmt.Printf("Subnet mask: %d.%d.%d.%d\n", cb.SubnetMaskBase10[0], cb.SubnetMaskBase10[1], cb.SubnetMaskBase10[2], cb.SubnetMaskBase10[3])
	fmt.Printf("Network address: %d.%d.%d.%d\n", cb.NetworkAddressBase10[0], cb.NetworkAddressBase10[1], cb.NetworkAddressBase10[2], cb.NetworkAddressBase10[3])
	fmt.Printf("Broadcast address: %d.%d.%d.%d\n", cb.BroadcastAddressBase10[0], cb.BroadcastAddressBase10[1], cb.BroadcastAddressBase10[2], cb.BroadcastAddressBase10[3])
	fmt.Printf("IP address range:")
	for i := 0; i < len(cb.IpAdressRangeBase10); i++ {
		fmt.Printf(" %d.%d.%d.%d", cb.IpAdressRangeBase10[i][0], cb.IpAdressRangeBase10[i][1], cb.IpAdressRangeBase10[i][2], cb.IpAdressRangeBase10[i][3])
	}
	fmt.Printf("\n")
	return 0
}

func parseCidrStr(cidr string) (string, string, error) {
	ret := strings.Split(cidr, "/")
	return ret[0], ret[1], nil
}

func networkAddress(ipv4 [4]uint8, ipv4SubnetMask [4]uint8) ([4]uint8, error) {
	var na [4]uint8
	for i := 0; i < 4; i++ {
		na[i] = ipv4[i] & ipv4SubnetMask[i]
	}
	return na, nil
}

func broadcastAddress(ipv4 [4]uint8, ipv4SubnetMask [4]uint8) ([4]uint8, error) {
	var na [4]uint8
	for i := 0; i < 4; i++ {
		notIpv4SubnetMask := ^ipv4SubnetMask[i]
		na[i] = ipv4[i] | notIpv4SubnetMask
	}
	return na, nil
}

func ipAdressRange(naIpv4, baIpv4 [4]uint8) [][4]uint8 {
	na := convertIpv4ArrayInto32BitInteger(naIpv4)
	ba := convertIpv4ArrayInto32BitInteger(baIpv4)
	num := ba - na + 1
	result := make([][4]uint8, num)
	for i := 0; i < int(num); i++ {
		result[i] = convert32BitIntegerIntoIpv4Array(na + uint32(i))
	}
	return result
}

func convertIpv4ArrayInto32BitInteger(ipv4 [4]uint8) uint32 {
	var result uint32
	result = 0
	for i := 0; i < 4; i++ {
		result = result | (uint32(ipv4[i]) << uint32((4-i-1)*8))
	}
	return result
}

func convert32BitIntegerIntoIpv4Array(itg uint32) [4]uint8 {
	var result [4]uint8
	for i := 0; i < 4; i++ {
		shift_num := uint32((4 - i - 1) * 8)
		result[i] = uint8(itg & uint32(255<<shift_num) >> shift_num)
	}
	return result
}

func convertIntoIpv4Array(ipv4 string) ([4]uint8, error) {
	ret := strings.Split(ipv4, ".")
	var ipv4Array [4]uint8
	for i := 0; i < 4; i++ {
		v, _ := strconv.ParseUint(ret[i], 10, 8)
		ipv4Array[i] = uint8(v)
	}
	return ipv4Array, nil
}

func convertIntoIpv4SubnetMaskArray(subnetMask string) ([4]uint8, error) {
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
	return ipv4Array, nil
}
