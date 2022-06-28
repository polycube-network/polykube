package polycube

type configuration struct {
	intK8sLbrpName   string
	rName            string
	extK8sLbrpName   string
	k8sDispName      string
	iklToRPortName   string
	rToIklPortName   string
	rToVxlanPortName string
	rToEklPortName   string
	eklToRPortName   string
	eklToKPortName   string
	kToEklPortName   string
	kToIntPortName   string
}
