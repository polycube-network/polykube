#!/bin/bash

ROOT_DIR="../"
CONF_DIR="$ROOT_DIR/conf"
BIN_DIR="$ROOT_DIR/plugins"

NETNS=cni-993652f5-c394-f657-b28f-5674de9515c7
CONTAINERID=61adb5d0d4

set -x
sudo ip netns del $NETNS
polycubectl ilb0 del
sudo rm -r /var/lib/cni/networks/mynet

set -e
sudo ip netns add $NETNS
polycubectl lbrp add ilb0 loglevel=TRACE

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
