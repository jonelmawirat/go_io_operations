package buffer

import (
    "errors"
    "io"
)

type Buffer struct {
    data []byte
    cursor int
}

func (b *Buffer) Write(p []byte) (int,error) {
    if len(p) <= 0 {
        return 0,errors.New("No Data")
    }

    b.data = append(b.data,p...)

    return len(p),nil
}


func (b *Buffer) Read(p []byte) (int,error) {

    if b.cursor >= len(b.data) {
        return 0,io.EOF
    }

    n := copy(p,b.data[b.cursor:])
    b.cursor += n

    if n < len(p) {
        return n,io.EOF
    }

    return n,nil
}

