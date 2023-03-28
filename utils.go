package main

func Unwrap(res any, err error) any {
	if err != nil {
		panic(err)
	}
	return res
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}