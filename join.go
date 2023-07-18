package paths

func Join(parts []string) string {
	s := ""
	for _, part := range parts {
		s += "/" + part
	}
	return s
}
