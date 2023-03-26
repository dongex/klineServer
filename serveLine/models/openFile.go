package models

import "os"

func OpenFileTest() {
	data := []byte("测试内容")
	f, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	if err != nil {
		panic(err)
	}
	_, err1 := f.Write(data)
	if err1 != nil {
		return
	}
}
