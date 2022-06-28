package polycube

type configuration struct {
	cubesLogLevel    string
	intLbrpName      string
	rName            string
	extLbrpName      string
	k8sDispName      string
	ilbToRPortName   string
	rToIlbPortName   string
	rToVxlanPortName string
	rToHostPortName  string
	rToElbPortName   string
	elbToRPortName   string
	elbToKPortName   string
	kToElbPortName   string
	kToIntPortName   string
}
