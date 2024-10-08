package x

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"golang.org/x/term"
)

func CopyFile(src, dst string) (int64, error) {
	srcF, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer srcF.Close()

	st, err := srcF.Stat()
	if err != nil {
		return 0, err
	}

	dstF, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer dstF.Close()

	written, err := io.Copy(dstF, srcF)
	if err != nil {
		return written, err
	}

	if err := dstF.Chmod(st.Mode()); err != nil {
		return written, err
	}
	return written, dstF.Close()
}

// CopyDir copies the contents of src to dst
// This function will be deprecated once os.CopyFS works with symlinks:
//
//	os.CopyFS(dst, os.DirFS(src))
func CopyDir(src, dst string) error {
	var links [][2]string
	if err := filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		newPath := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(newPath, 0o755)
		} else if d.Type().IsRegular() {
			_, err = CopyFile(path, newPath)
			return err
		} else if d.Type()&os.ModeSymlink != 0 {
			target, err := os.Readlink(path)
			if err != nil {
				return err
			}
			links = append(links, [2]string{target, newPath})
			return nil
		}
		return fmt.Errorf("unable to copy file: %s", path)
	}); err != nil {
		return err
	}
	for _, link := range links {
		if err := os.Symlink(link[0], link[1]); err != nil && !os.IsExist(err) {
			return err
		}
	}
	return nil
}

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
