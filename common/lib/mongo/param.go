package mongo

import "flag"

func Flags() (host string, port int) {

	hostPtr := flag.String("host", "localhost", "Target MongoDB Host")
	portPtr := flag.Int("port", 27017, "Target MongoDB Port")

	flag.Parse()

	return *hostPtr, *portPtr
}

func FlagsNoParse() (hostPtr *string, portPtr *int) {
	hostPtr = flag.String("host", "localhost", "Target MongoDB Host")
	portPtr = flag.Int("port", 27017, "Target MongoDB Port")
	return
}
