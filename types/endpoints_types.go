package types

type Backend struct {
	Ip     string
	Port   int32
	Weight int32
}
type BackendsSet map[Backend]struct{}

func (bs BackendsSet) Contains(b Backend) bool {
	_, ok := bs[b]
	return ok
}
func (bs BackendsSet) Add(b Backend) {
	bs[b] = struct{}{}
}

type ServiceToBackends map[string]BackendsSet

func (stb ServiceToBackends) Add(s string, b Backend) {
	if stb[s] == nil {
		stb[s] = make(BackendsSet)
	}
	stb[s][b] = struct{}{}
}

func (stb ServiceToBackends) GetBackendsSet(s string) BackendsSet {
	return stb[s]
}

type EndpointsDetail struct {
	EndpointsId                string
	ClusterIPServiceToBackends ServiceToBackends
	NodePortServiceToBackends  ServiceToBackends
}
