package cidr

import (
	"testing"
)

func TestNewCidrCase001(t *testing.T) {
	cb, _ := NewCidr("192.168.32.0/28")
	checkNetworkAddress("Case001", cb.GetNetworkAddress(), "192.168.32.0", t)
	checkBroadcastAddress("Case001", cb.GetBroadcastAddress(), "192.168.32.15", t)
	checkCidr("Case001", cb.GetCidr(), "192.168.32.0/28", t)
}

func TestNewCidrCase002(t *testing.T) {
	cb, _ := NewCidr("192.168.32.0/30")
	checkNetworkAddress("Case002", cb.GetNetworkAddress(), "192.168.32.0", t)
	checkBroadcastAddress("Case002", cb.GetBroadcastAddress(), "192.168.32.3", t)
	checkCidr("Case002", cb.GetCidr(), "192.168.32.0/30", t)
}

func TestNewCidrCase003(t *testing.T) {
	cb, _ := NewCidr("10.100.0.0/8")
	checkNetworkAddress("Case003", cb.GetNetworkAddress(), "10.0.0.0", t)
	checkBroadcastAddress("Case003", cb.GetBroadcastAddress(), "10.255.255.255", t)
	checkCidr("Case003", cb.GetCidr(), "10.0.0.0/8", t)
}

func TestNewCidrCase101(t *testing.T) {
	var input string
	input = "192.168.32.0/a"
	_, err := NewCidr(input)
	if err == nil {
		t.Errorf("Case101 case failed. Invalid input format -> %s", input)
	}
}

func TestNewCidrCase102(t *testing.T) {
	var input string
	input = "1920.168.32.0/10"
	_, err := NewCidr(input)
	if err == nil {
		t.Errorf("Case101 case failed. Invalid input format -> %s", input)
	}
}

func checkNetworkAddress(caseID, actualAddress, okAddress string, t *testing.T) {
	if actualAddress != okAddress {
		t.Errorf("%s case failed. Network address incorrectly. %s is correctly, but actual address is %s .\n", caseID, okAddress, actualAddress)
	}
}

func checkBroadcastAddress(caseID, actualAddress, okAddress string, t *testing.T) {
	if actualAddress != okAddress {
		t.Errorf("%s case failed. Broadcast address incorrectly. %s is correctly, but actual address is %s .\n", caseID, okAddress, actualAddress)
	}
}

func checkCidr(caseID, actualCidr, okCidr string, t *testing.T) {
	if actualCidr != okCidr {
		t.Errorf("%s case failed. CIDR is incorrectly. %s is correctly, but actual address is %s .\n", caseID, okCidr, actualCidr)
	}
}
