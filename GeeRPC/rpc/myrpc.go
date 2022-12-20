package rpc

var m = make(map[string]interface{})

func RegisterService(serviceName string, inter interface{}) {
	m[serviceName] = inter
}

func FindService(serviceName string) interface{} {
	return m[serviceName]
}
