#!/bin/bash

ROOT_DIR="../"
CONF_DIR="$ROOT_DIR/conf"
BIN_DIR="$ROOT_DIR/bin"

set -x
set -e

sudo CNI_COMMAND=CHECK \
CNI_CONTAINERID=containerid \
CNI_NETNS=/run/netns/ns1 \
CNI_IFNAME=veth1 \
CNI_PATH=$BIN_DIR \
$BIN_DIR/polykube-cni-plugin < $CONF_DIR/00-polykube-check-test.json
