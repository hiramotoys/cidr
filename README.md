# cidr

[![Build Status](https://travis-ci.org/hiramotoys/cidr.svg?branch=master)](https://travis-ci.org/hiramotoys/cidr)

This package provides a command line interface to get a network cidr information.

The cidr package needs golang.

## Installation
The easiest way to install cidr is to use `go get` command.
```
go get github.com/hiramotoys/cidr
```

## How to use
You can get a network cidr information with `cidr detail` command.
```
$ cidr detail 192.168.32.16/29
```

```
CIDR: 192.168.32.16/29
IP address: 192.168.32.16
Subnet mask: 255.255.255.248
Network address: 192.168.32.16
Broadcast address: 192.168.32.23
IP address range: 192.168.32.16 192.168.32.17 192.168.32.18 192.168.32.19 192.168.32.20 192.168.32.21 192.168.32.22 192.168.32.23
```
