package config

type Config struct {
	// HTTP Server
	HttpServerPort     int
	MaxMultipartMemory int

	// IPFS
	IpfsPort string
	IpfsHost string
	IpfsPin  bool
}

var DefaultConfig = Config{
	HttpServerPort:     8000,
	MaxMultipartMemory: 8 << 20, // 8 MiB
	IpfsPort:           "5001",
	IpfsHost:           "localhost",
	IpfsPin:            true,
}
