package pforward

type ForwardCnf struct {
	serverPort int32
	destHost   string
	destPort   int32
}

func Forward(cnf ForwardCnf) {
}
