// Code generated by "esc -modtime 12345 -prefix openapi/ -pkg openapi -ignore .go -o openapi/assets.go ."; DO NOT EDIT.

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/asset-gen.sh": {
		local:   "asset-gen.sh",
		size:    238,
		modtime: 12345,
		compressed: `
H4sIAAAAAAAC/0zKz0rEMBDH8Xue4rfTnBbSsP45LR5EfAHrTUTWdpIO0hlJIgjiu0sr6M5l4PP9dbv4
Khrr7NztMNw/vgwPdzf+4JIVCERB/s8p7o+YzAGAJOzwhDDBC56PaDPrFtYbTZvoB2+QxG2fx9lAmZXL
qYlmpGILvNBvrSPCYlOThUGHi8ura0J4L5zkE+S/pOv28Tuu9pb/gRAkqxVGnw3BzqanWrnVPhuBenKT
KbufAAAA//9BiTev7gAAAA==
`,
	},

	"/index.html": {
		local:   "openapi/index.html",
		size:    636,
		modtime: 12345,
		compressed: `
H4sIAAAAAAAC/0ySQW/bMAyF7/kVjC+9RJaHDtiQyd6wpceuQ9DLblUk2lYrS55IpzC2/ffBUdLlRr4n
fXwgqNa7h2+PP3/cQc+Db1ZqLcTq+8Pj3Rb2U4CnQb8gaCJk0WEQvyZM8xO4FuY4QTbDDKbXoUMCjsC9
I2idx/VKiGalMhZA9ajtUgAoduyxub/dfYU97qJRMivZHZD1QkyEXBcTt+JjIa+9oAesi6PD1zEmLsDE
wBi4Ll6d5b62eHQGxanZgAuOnfaCjPZYvyvOIO/CC/QJ27romUfaStnGwFR2MXYe9eioNHGQhuhzqwfn
5/p+8TElzdvbqtq8r6rNh6r6s4+HyPFaKiChrwvi2SP1iHwZelJyDXCIdobf5wZg0KlzYQvVpzdp1Na6
0F1pfzNHvoGUvKxVLbzznIQ2GqARjZiSr2/iiEGPThJrdkYuRjkP/qZR8vT0Es8kNzJQMv+XYmwon8mi
d8dUBmQZxiF/+uI1I7E8TMF6pCyWxDpY7WPA8pmKZsl6ouawOaOS+Sj+BQAA//8by2IcfAIAAA==
`,
	},

	"/spec.yml": {
		local:   "openapi/spec.yml",
		size:    11361,
		modtime: 12345,
		compressed: `
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
KKHvaavhGnrjU8KKjv8GAAD//1caENhhLAAA
`,
	},

	"/": {
		isDir: true,
		local: "openapi",
	},
}
