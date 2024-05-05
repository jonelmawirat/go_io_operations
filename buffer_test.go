package buffer

import (
    "testing"
    "io"
)




func Test_NoData(t *testing.T){
    test_data := []byte{}
    buff := Buffer{}
    _,err := buff.Write(test_data)
    if err == nil {
        t.Errorf("Got Nil, Want: error(No Data)")
    }
}

func Test_Short_Container(t *testing.T) {
    test_data := []byte("123456789")
    buff := Buffer{}
    buff.Write(test_data)

    container := make([]byte,8)
    n,err := buff.Read(container)

    for i := 0; i < len(container); i++ {
        if container[i] != test_data[i] {
            t.Errorf("Want %v, Got: %v\n",string(test_data[i]),string(container[i]))
        }
    }

    if n != len(container) {
        t.Errorf("Want 8, Got: %v\n",n)
    }

    if err != nil {
        if err != io.EOF {
            t.Errorf("Want io.EOF, Got: %v\n",err)
        }
    }

    another_container := make([]byte,1)
    nn, err := buff.Read(another_container)
    if nn != 1 {
        t.Errorf("Want: 1, Got: %v\f",nn)
    }

    if err != nil {
        if err != io.EOF {
            t.Errorf("Want io.EOF, Got: %v\n",err)
        }
    }

    if another_container[0] != test_data[len(test_data) - 1] {
        t.Errorf("Want: %v, Got: %v\n",string(another_container[len(another_container) - 1]), string(test_data[len(test_data) - 1]))
    }


}
