package procFile



func ReadProc(s string) string {
	switch s {
	case "mem":
		return "/proc/meminfo"
	case "cpu":
		return "/proc/stat"
	case "netconn":
		return "/proc/net/tcp"
	}
	return ""
}
