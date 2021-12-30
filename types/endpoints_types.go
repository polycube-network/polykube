package types

type Backend struct {
	Ip   string
	Port int32
}
type BackendsSet map[Backend]struct{}

func (fs BackendsSet) Contains(b Backend) bool {
	_, ok := fs[b]
	return ok
}

type ServiceToBackends map[string]BackendsSet

func (stb ServiceToBackends) Add(s string, b Backend) {
	if stb[s] == nil {
		stb[s] = make(BackendsSet)
	}
	stb[s][b] = struct{}{}
}

func (stb ServiceToBackends) AddBackends(s string, bs BackendsSet) {
	stb[s] = bs
}

func (stb ServiceToBackends) CountDistinctServices() int {
	return len(stb)
}
func (stb ServiceToBackends) GetBackendsSet(s string) BackendsSet {
	return stb[s]
}
func (stb ServiceToBackends) checkService(s string) bool {
	_, ok := stb[s]
	return ok
}
func (stb ServiceToBackends) Contains(s string, b Backend) bool {
	if _, ok := stb[s]; !ok {
		return false
	}
	_, ok := stb[s][b]
	return ok
}

type EndpointsDetail struct {
	EndpointsId                string
	ClusterIPServiceToBackends ServiceToBackends
	NodePortServiceToBackends  ServiceToBackends
}
