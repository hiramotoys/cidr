package cidr

import (
	"testing"
)

func TestNewCidr(t *testing.T) {
	cb, _ := NewCidr("192.168.32.0/28")
	var networkaddress string
	var broadcastaddress string
	networkaddress = cb.GetNetworkAddress()
	broadcastaddress = cb.GetBroadcastAddress()
	if networkaddress != "192.168.32.0" {
		t.Errorf("Failed network address unmatched: %s\n", networkaddress)
	}
	if broadcastaddress != "192.168.32.15" {
		t.Errorf("Failed broadcast address unmatched: %s\n", broadcastaddress)
	}
}
