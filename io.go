package x

import (
	"errors"
	"io"
	"os"

	"golang.org/x/term"
)

func StdinIsPipe() bool {
	mode := Must(os.Stdin.Stat()).Mode()
	return mode&os.ModeNamedPipe != 0 && mode&os.ModeCharDevice == 0
}

func StdoutIsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// Closing is a shortcut, instead of writing:
//
//	defer func() { C(f.Close()) }()
//
// One can write:
//
//	defer Closing(f)
func Closing(v io.Closer) {
	Check(v.Close())
}

type (
	closeAfterReader         struct{ r io.ReadCloser }
	closeAfterReaderWriterTo closeAfterReader
)

func (c closeAfterReader) Read(b []byte) (int, error) {
	n, err := c.r.Read(b)
	if err == nil {
		return n, nil
	}
	closeErr := c.r.Close()
	if err == io.EOF && closeErr == nil {
		return n, io.EOF
	}
	return n, errors.Join(err, closeErr)
}

func (c closeAfterReaderWriterTo) Read(b []byte) (int, error) {
	return closeAfterReader(c).Read(b)
}

func (c closeAfterReaderWriterTo) WriteTo(w io.Writer) (int64, error) {
	n, err := c.r.(io.WriterTo).WriteTo(w)
	return n, errors.Join(err, c.r.Close())
}

// CloseAfterRead returns a Reader that automatically closes when there is no more data to read or an error has occurred.
// Examples:
//
//	// Prints SHA2 of "file_to_hash" in hexadecimal notation
//	h := sha256.New()
//	C2(io.Copy(h, CloseAfterRead(C2(os.Open("file_to_hash")))))
//	fmt.Println(hex.EncodeToString(h.Sum(nil)))
//
//	// Downloads a file
//	const url = "https://go.dev/dl/go1.20.2.linux-amd64.tar.gz"
//	dst := C2(os.Create(path.Base(url)))
//	defer Closing(dst)
//	C2(io.Copy(dst, CloseAfterRead(C2(http.Get(url)).Body)))
func CloseAfterRead(r io.ReadCloser) io.Reader {
	if _, ok := r.(io.WriterTo); ok {
		return closeAfterReaderWriterTo{r}
	}
	return closeAfterReader{r}
}
