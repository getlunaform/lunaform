package database

import (
	"io"
	"os"
)

type mockCall struct {
	Name      string
	Arguments []interface{}
}

type mockFile struct {
	Calls []mockCall
}

func (f mockFile) Close() error {
	f.Calls = append(f.Calls, mockCall{
		Name: "Close",
	})
	return nil
}

func (f mockFile) Read(b []byte) (int, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Read",
		Arguments: []interface{}{b},
	})
	r := "{ \"collections 0 name\": \"hello\"}"
	copy(b[:], []byte(r))
	return len(r), io.EOF
}

func (f mockFile) Seek(a int64, b int) (int64, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Seek",
		Arguments: []interface{}{a, b},
	})
	return 0, nil
}

func (f mockFile) Stat() (os.FileInfo, error) {
	f.Calls = append(f.Calls, mockCall{
		Name: "Stat",
	})
	return nil, nil
}

func (f mockFile) Write(b []byte) (int, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Write",
		Arguments: []interface{}{b},
	})
	return 0, nil
}

func (f mockFile) WriteAt(b []byte, i int64) (int, error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "WriteAt",
		Arguments: []interface{}{b},
	})
	return 0, nil
}

func (f mockFile) Truncate(i int64) (err error) {
	f.Calls = append(f.Calls, mockCall{
		Name:      "Write",
		Arguments: []interface{}{i},
	})
	return nil
}
