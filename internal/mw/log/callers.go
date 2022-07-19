package log

func getCallers(err error) []string {
	c, ok := err.(interface{ Callers() []string })
	if !ok {
		return []string{}
	}
	return c.Callers()
}
