package cidr

import (
	"testing"
)

func TestNewCidrCase001(t *testing.T) {
	cb, _ := NewCidr("192.168.32.0/28")
	checkNetworkAddress("Case001", cb.GetNetworkAddress(), "192.168.32.0", t)
	checkBroadcastAddress("Case001", cb.GetBroadcastAddress(), "192.168.32.15", t)
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
