#!/bin/bash

function error_message {
  set +x
  echo
  echo "Error during installation"
}

function success_message {
  set +x
  echo
  echo "Installation completed successfully"
}
trap error_message ERR

set -x
set -e
set -u

HOST_PREFIX=${HOST_PREFIX:-/host}
CNI_DIR=${CNI_DIR:-${HOST_PREFIX}/opt/cni}

# Install the CNI loopback plugin if not already installed
if [ ! -f "${CNI_DIR}"/bin/loopback ]; then
	echo "Installing loopback plugin..."
	cp /cni-plugins/loopback "${CNI_DIR}/bin/" || (echo "Failed to copy loopback plugin" && exit 1)
	echo "loopback plugin installed"
fi
# Install the host-local IPAM plugin if not already installed
if [ ! -f "${CNI_DIR}"/bin/host-local ]; then
	echo "Installing host-local IPAM plugin..."
	cp /cni-plugins/host-local "${CNI_DIR}/bin/" || (echo "Failed to copy host-local IPAM plugin" && exit 2)
  echo "host-local IPAM plugin installed"
fi
# Always replace the "old" polykube plugin
echo "Installing Polykube CNI plugin..."
cp /cni-plugins/polykube-cni-plugin "${CNI_DIR}/bin/" || (echo "Failed to copy Polykube CNI plugin" && exit 3)
echo "Polykube CNI plugin installed"