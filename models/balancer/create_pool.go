package balancer

type CreatePool struct {
	LoadBalancer `argument:"composed" URIParam:"LoadBalancerId,DataCenter"`
	Port         int64  `valid:"required"`
	Method       string `oneOf:"roundRobin,leastConnection"`
	Persistence  string `oneOf:"standard,sticky"`
}
