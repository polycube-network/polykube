package controllers

type Frontend struct {
	Vip   string
	Vport int32
	Proto string
}

type FrontendsSet map[Frontend]struct{}

func (fs FrontendsSet) Contains(f Frontend) bool {
	_, ok := fs[f]
	return ok
}
func (fs FrontendsSet) Add(f Frontend) {
	fs[f] = struct{}{}
}

type ServiceDetail struct {
	NodePortFrontendsSet  FrontendsSet
	ClusterIPFrontendsSet FrontendsSet
}
