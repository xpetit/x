# X

A collection of functions to write concise code.

```go
sha2 := sha256.New()
src := tar.NewReader(
	CloseAfterRead(C2(gzip.NewReader(
		io.TeeReader(
			CloseAfterRead(C2(http.Get("https://go.dev/dl/go1.20.2.linux-amd64.tar.gz")).Body),
			sha2,
		),
	))),
)
for th, err := src.Next(); err != io.EOF; th, err = src.Next() {
	C(err)
	switch th.Typeflag {
	case tar.TypeDir:
		C(os.Mkdir(th.Name, 0o755))
	case tar.TypeReg:
		dst := C2(os.OpenFile(th.Name, os.O_CREATE|os.O_WRONLY, th.FileInfo().Mode()))
		C2(io.Copy(dst, src))
		C(dst.Close())
	}
}
Assert(hex.EncodeToString(sha2.Sum(nil)) == "4eaea32f59cde4dc635fbc42161031d13e1c780b87097f4b4234cfce671f1768")
```

This snippet downloads an archive, unpacks it and verifies its checksum, all in one go.
Any error causes a panic, cascading errors are stacked and each reader is closed as expected.

---

If you get a warning from `gopls/staticcheck` about dot-import, create a file named `staticcheck.conf` in your project directory (or parents) with:

```
dot_import_whitelist = [ "github.com/xpetit/x/v3" ]
```
