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

func (w *Writer) Write(p []byte) (n int, err error) {
	w.Hash.Write(p)
	return w.proxy.Write(p)
}

// PipeSum writes checksum of current hash to provided Writer
func (w *Writer) PipeSum(dest io.Writer) {
	buf := bufio.NewWriter(dest)
	fmt.Fprintf(buf, "% x", w.Hash.Sum(nil))
	buf.Flush()
}
