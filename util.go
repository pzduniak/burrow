package burrow

import (
	"bufio"
	"io"
)

func CopyData(r io.Reader, w io.Writer) error {
	rb := bufio.NewReader(r)
	wb := bufio.NewWriter(w)

	buffer := make([]byte, 1024)
	for {
		n, err := rb.Read(buffer)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if n == 0 {
			break
		}

		if _, err := wb.Write(buffer[:n]); err != nil {
			return err
		}
	}

	if err := wb.Flush(); err != nil {
		return err
	}

	return nil
}
