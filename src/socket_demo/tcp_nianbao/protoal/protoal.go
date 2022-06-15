package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

func Encode(message string) ([]byte, error) {
	//读取消息长度转换为int32(4字节)
	var length = int32(len(message))
	var pkg = new(bytes.Buffer)
	//将消息长度作为消息头写入缓冲区(小端)
	err := binary.Write(pkg, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	//将消息转化为字节写入
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}
	return pkg.Bytes(), nil
}

func Decode(reader *bufio.Reader) (string, error) {
	//读取消息
	lengthByte, _ := reader.Peek(4)           //读取reader前4个字节[25 0 0 0]
	lengthBuff := bytes.NewBuffer(lengthByte) //创建缓冲区
	var length int32
	err := binary.Read(lengthBuff, binary.LittleEndian, &length) //读取缓冲区长度length=25
	if err != nil {
		return "", err
	}
	//Buffered返回缓冲区中现有的可读字节数
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}
	//读取消息
	pack := make([]byte, int(4+length)) //创建固定长度切片
	_, err = reader.Read(pack)
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}
