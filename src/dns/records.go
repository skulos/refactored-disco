package dns

/*
type Record struct {
}*/

// IP4 represents an IP address.
type IP4 struct {
	Byte1 byte
	Byte2 byte
}

// A DNS Record: address mapping record. Also known as a DNS host record, resolves a hostname to its corresponding IPv4 address.
type A struct {
	Host string
	IP   [4]byte
}
