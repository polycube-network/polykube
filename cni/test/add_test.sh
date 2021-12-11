#!/bin/bash

ROOT_DIR="../"
CONF_DIR="$ROOT_DIR/conf"
BIN_DIR="$ROOT_DIR/plugins"

NETNS=cni-993652f5-c394-f657-b28f-5674de9515c7
CONTAINERID=61adb5d0d4

set -x
sudo ip netns del $NETNS
polycubectl ikl0 del
sudo rm -r /var/lib/cni/networks/mynet

set -e
sudo ip netns add $NETNS
polycubectl k8slbrp add ikl0 loglevel=TRACE

#
#set -x
#sudo ip netns del ns1
#sudo ip link del dev gw
#sudo rm -r /var/lib/cni/networks/testnet
#polycubectl lbrp_veth1_container del
#polycubectl br0 del
#
#set -e
#sudo ip netns add ns1
#sudo ip link add gw type dummy
#sudo ip link set dev gw address aa:bb:cc:dd:ee:ff
#sudo ip link set dev gw up
#sudo ip addr add 10.0.1.254/24 dev gw
#polycubectl simplebridge add br0

RESULT=$(sudo CNI_COMMAND=ADD \
CNI_CONTAINERID=$CONTAINERID \
CNI_NETNS=/run/netns/$NETNS \
CNI_IFNAME=eth0 \
CNI_PATH=$BIN_DIR \
$BIN_DIR/polykube-cni-plugin < $CONF_DIR/sample.json)
set +x
if [[ ! -z "$RESULT" ]]; then
	jq --argjson RESULT "$RESULT" '. += {prevResult: $RESULT}' $CONF_DIR/sample.json > $CONF_DIR/sample-check-test.json
fi
