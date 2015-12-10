package checksum

import (
	"bufio"
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
)

// Writer wraps a writer and
// computes data checksum
type Writer struct {
	Hash  hash.Hash
	proxy io.Writer
}

// NewWriter creates a ChecksumWriter
func NewWriter(proxy io.Writer) *Writer {
	h := sha1.New()
	return &Writer{h, proxy}
}

func (cw *Writer) Write(p []byte) (n int, err error) {
	cw.Hash.Write(p)
	return cw.proxy.Write(p)
}

// PipeSum writes checksum of current hash to provided Writer
func (cw *Writer) PipeSum(w io.Writer) {
	buf := bufio.NewWriter(w)
	fmt.Fprintf(buf, "% x", cw.Hash.Sum(nil))
	buf.Flush()
}
