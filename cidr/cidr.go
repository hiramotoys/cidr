package cidr

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"math/bits"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CidrBlock struct {
	ipAddress        uint32
	subnetMask       uint32
	networkAddress   uint32
	broadcastAddress uint32
	ipAddressRange   [][4]uint8
}

func NewCidr(cidr string) (*CidrBlock, error) {
	cidrBlock := CidrBlock{}
	ipAddressStr, subnetMaskStr, err := parseCmmandLineInput(cidr)
	if err != nil {
		logrus.Errorf("%s\n", err.Error())
		os.Exit(1)
	}
	ipAddress := convertIpv4StrInto32bitInteger(ipAddressStr)
	subnetMask := subnetMask32bitInteger(subnetMaskStr)
	cidrBlock.ipAddress = ipAddress
	cidrBlock.subnetMask = subnetMask
	cidrBlock.setNetworkAddress()
	cidrBlock.setBroadcastAddress()
	cidrBlock.ipAddressRange = getIpAdressRange(cidrBlock.networkAddress, cidrBlock.broadcastAddress)
	return &cidrBlock, err
}

func (cb *CidrBlock) Print() int {
	fmt.Printf("CIDR: %s\n", cb.GetCidr())
	ipAddressArray := convert32bitIntegerIntoIpv4Array(cb.ipAddress)
	fmt.Printf("IP address: %d.%d.%d.%d\n", ipAddressArray[0], ipAddressArray[1], ipAddressArray[2], ipAddressArray[3])
	subnetMaskArray := convert32bitIntegerIntoIpv4Array(cb.subnetMask)
	fmt.Printf("Subnet mask: %d.%d.%d.%d\n", subnetMaskArray[0], subnetMaskArray[1], subnetMaskArray[2], subnetMaskArray[3])
	networkAddressArray := convert32bitIntegerIntoIpv4Array(cb.networkAddress)
	fmt.Printf("Network address: %d.%d.%d.%d\n", networkAddressArray[0], networkAddressArray[1], networkAddressArray[2], networkAddressArray[3])
	broadcastAddressArray := convert32bitIntegerIntoIpv4Array(cb.broadcastAddress)
	fmt.Printf("Broadcast address: %d.%d.%d.%d\n", broadcastAddressArray[0], broadcastAddressArray[1], broadcastAddressArray[2], broadcastAddressArray[3])
	fmt.Printf("IP address range:")
	for i := 0; i < len(cb.ipAddressRange); i++ {
		fmt.Printf(" %d.%d.%d.%d", cb.ipAddressRange[i][0], cb.ipAddressRange[i][1], cb.ipAddressRange[i][2], cb.ipAddressRange[i][3])
	}
	fmt.Printf("\n")
	return 0
}

func parseCmmandLineInput(cidr string) (string, string, error) {
	re, _ := regexp.MatchString(`^[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\/[1-9]{1,2}$`, cidr)
	if !re {
		return "", "", fmt.Errorf("Invalid input format -> %s", cidr)
	}
	ret := strings.Split(cidr, "/")
	v, err := strconv.Atoi(ret[1])
	if err == nil && (v < 1 || v > 32) {
		err = fmt.Errorf("Invalid cidr value -> /%d", v)
	}
	return ret[0], ret[1], err
}

func (cb *CidrBlock) GetCidr() string {
	b := IpV4Str(cb.networkAddress)
	b += "/" + strconv.Itoa(bits.OnesCount32(cb.subnetMask))
	return b
}

func (cb *CidrBlock) GetNetworkAddress() string {
	return IpV4Str(cb.networkAddress)
}

func (cb *CidrBlock) GetBroadcastAddress() string {
	return IpV4Str(cb.broadcastAddress)
}

func IpV4Str(ipaddress uint32) string {
	ipV4Array := convert32bitIntegerIntoIpv4Array(ipaddress)
	b := ""
	length := 4
	for i := 0; i < length; i++ {
		b += strconv.Itoa(int(ipV4Array[i]))
		if i < length-1 {
			b += "."
		}
	}
	return b
}

func (cb *CidrBlock) setNetworkAddress() {
	cb.networkAddress = cb.ipAddress & cb.subnetMask
}

func (cb *CidrBlock) setBroadcastAddress() {
	cb.broadcastAddress = cb.ipAddress | (^cb.subnetMask)
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
	ipv4Array := convertIpv4StrIntoIpv4Array(ipv4Str)
	return convertIpv4ArrayInto32bitInteger(ipv4Array)
}

func convertIpv4StrIntoIpv4Array(ipv4 string) [4]uint8 {
	ret := strings.Split(ipv4, ".")
	var ipv4Array [4]uint8
	for i := 0; i < 4; i++ {
		v, _ := strconv.ParseUint(ret[i], 10, 8)
		ipv4Array[i] = uint8(v)
	}
	return ipv4Array
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
		shiftNum := uint32((4 - i - 1) * 8)
		result[i] = uint8(itg & uint32(255<<shiftNum) >> shiftNum)
	}
	return result
}

func subnetMask32bitInteger(subnetMask string) uint32 {
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
	return convertIpv4ArrayInto32bitInteger(ipv4Array)
}
