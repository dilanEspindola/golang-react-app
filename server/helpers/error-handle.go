package helpers

func ErrorHandle(err error) {
	if err != nil {
		panic(err)
	}
}
