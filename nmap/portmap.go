package nmap

// These maps are a best guess and are used if service fingerprinting wasnt enabled
var tcpPorts = map[int]string{
	21:   "ftp",
	22:   "ssh",
	23:   "telnet",
	25:   "smtp",
	53:   "dns",
	80:   "web",
	110:  "pop3",
	143:  "imap",
	443:  "web",
	3389: "rdp",
	8080: "web",
	8443: "web",
	9000: "web",
}

// tcpServiceName will return the known service name or "unknown"
func tcpServiceName(p int) string {
	if val, ok := tcpPorts[p]; ok {
		return val
	}
	return "unknown"
}
