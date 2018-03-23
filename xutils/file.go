package xutils

import "io/ioutil"

func File_Read(filename string) ([]byte, error) {

	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return bs, nil
}

func File_Write(filename string, bs []byte) {

	ioutil.WriteFile(filename, bs, 0750)

}
