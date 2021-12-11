#!/bin/bash

function error_message {
  set +x
  echo
  echo "Error during uninstallation"
}

function success_message {
  set +x
  echo
  echo "Uninstallation completed successfully"
}
trap error_message ERR

set -x
set -e
set -u

HOST_PREFIX=${HOST_PREFIX:-/host}
CNI_DIR=${CNI_DIR:-${HOST_PREFIX}/opt/cni}
CNI_CONF_FILE_PATH=${CNI_CONF_FILE_PATH:-${HOST_PREFIX}/etc/cni/net.d/00-polykube.json}

echo "Uninstalling Polykube CNI plugin..."
rm -rf "${CNI_DIR}/bin/polykube-cni-plugin" || (echo "Failed to remove Polykube CNI plugin" && exit 1)
echo "Polykube CNI plugin uninstalled"

echo "Removing Polykube CNI configuration..."
rm -rf "${CNI_CONF_FILE_PATH}" || (echo "Failed to remove Polykube CNI configuration" && exit 1)
echo "Polykube CNI configuration removed"