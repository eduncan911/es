package api

func guard(name string, err error) {
	if err != nil {
		l.Panicf("%s: %s", name, err)
	}
}
