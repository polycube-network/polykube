#!/bin/bash

ROOT_DIR="../"
CONF_DIR="$ROOT_DIR/conf"
BIN_DIR="$ROOT_DIR/bin"

set -x
sudo ip netns del ns1
sudo ip link del dev gw
sudo rm -r /var/lib/cni/networks/mynet
polycubectl lbrp-containeri del
polycubectl br0 del

set -e
sudo ip netns add ns1
sudo ip link add gw type dummy
sudo ip link set dev gw address aa:bb:cc:dd:ee:ff
sudo ip link set dev gw up
sudo ip addr add 10.0.1.254/24 dev gw
polycubectl simplebridge add br0

sudo CNI_COMMAND=ADD \
CNI_CONTAINERID=containerid \
CNI_NETNS=/run/netns/ns1 \
CNI_IFNAME=veth1 \
CNI_PATH=$BIN_DIR \
NETCONF=$CONF_DIR/00-polykube-net.json \
./test_chaining
