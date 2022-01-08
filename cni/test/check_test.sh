#!/bin/bash

ROOT_DIR="../"
CONF_DIR="$ROOT_DIR/conf"
BIN_DIR="$ROOT_DIR/plugins"

NETNS=cni-993652f5-c394-f657-b28f-5674de9515c7
CONTAINERID=61adb5d0d4

set -x
set -e

sudo CNI_COMMAND=CHECK \
CNI_CONTAINERID=$CONTAINERID \
CNI_NETNS=/run/netns/$NETNS \
CNI_IFNAME=eth0 \
CNI_PATH=$BIN_DIR \
$BIN_DIR/polykube-cni-plugin < $CONF_DIR/sample-check-test.json
