# Polykube operator
Polykube is an eBPF-based disaggragated network provider for Kubernetes. This repository contains the operator and the
CNI plugin code. In order to install this provider in a fresh new cluster, you must:
- clone this repository
- navigate inside the downloaded folder
- replace the `REPLACE_WITH_API_SERVER_IP` and `REPLACE_WITH_API_SERVER_PORT` placeholders in
  `manifests/config_map.yaml` with the real IP address and the port of the node on
  which the API Server is exposed
- optionally, customize the content of `manifests/config_map.yaml` with your preferences
- install the operator by issuing `kubectl apply -f manifests`