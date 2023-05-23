package helper

func HandleIfPanicError(err error) {
	if err != nil {
		// fmt.Println(err)
		panic(err)
	}
}
