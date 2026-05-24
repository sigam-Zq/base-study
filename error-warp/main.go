package main

import (
	"fmt"
	"io"
)

type Header struct {
	Key, Value string
}

type Status struct {
	Code   int
	Reason string
}

func WriteResponse(w io.Writer, st Status, headers []Header, body io.Reader) error {
	_, err := fmt.Fprintf(w, "HTTP/1/1 %d %s\r\n", st.Code, st.Reason)
	if err != nil {
		return err
	}
	for _, n := range headers {
		_, err := fmt.Fprintf(w, "%s: %s\r\n", n.Key, n.Value)
		if err != nil {
			return err
		}
	}

	if _, err := fmt.Fprintf(w, "\r\n"); err != nil {
		return err
	}
	_, err = io.Copy(w, body)
	return err
}

//  上面的基础上封装 包装类进行错误放封装

type WarpErrWriter struct {
	io.Writer
	err error
}

func (w *WarpErrWriter) Write(buf []byte) (int, error) {
	if w.err != nil {
		return 0, w.err
	}
	var n int
	n, w.err = w.Write(buf)
	return n, nil
}

func WriteResponseRf(w io.Writer, st Status, headers []Header, body io.Reader) error {
	wr := &WarpErrWriter{
		Writer: w,
	}
	fmt.Fprintf(wr, "HTTP/1/1 %d %s\r\n", st.Code, st.Reason)

	for _, n := range headers {
		fmt.Fprintf(wr, "%s: %s\r\n", n.Key, n.Value)

	}
	// errors.Wrap()
	fmt.Fprintf(wr, "\r\n")
	io.Copy(wr, body)
	return wr.err
}
