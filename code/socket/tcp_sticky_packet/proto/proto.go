package proto

import (
	"bufio"
	"bytes"
	"encoding/binary"
)

// 自定义协议解决 TCP 黏包问题

/*
TCP/IP(Transmission Control Protocol/Internet Protocol) 即传输控制协议/网间协议，
是一种面向连接（连接导向）的、可靠的、基于字节流的传输层（Transport layer）通信协议，
因为是面向连接的协议，数据像水流一样传输，会存在黏包问题。

需要注意TCP 黏包的问题，与编程语言无关，粘包问题并不是语言本身引起的，而是由 TCP 协议的传输机制 决定的。
TCP 是基于字节流的协议，它并不提供消息边界的概念(需要应用层自己去解决消息边界的识别问题)，传输的数据可能会被拆分成多个包发送，
也可能会将多个包合并为一个包来传输。
- 解决 TCP 粘包问题的常见方案:

  - 1.固定长度消息
    发送和接收方约定每条消息的长度是固定的。接收方可以按固定长度读取数据。适合小消息或固定格式消息的传输。

  - 2. 使用特殊分隔符
    在每条消息后面附加一个特殊字符（如 \n 或自定义分隔符），接收方通过分隔符识别消息的边界。这种方式比较灵活，常用于文本数据的传输。

  - 3. 自定义消息头+消息体 (自定义协议)
    通过在每条消息前附加一个消息头，消息头中包含消息体的长度信息。接收方首先读取消息头的长度，然后根据该长度读取完整的消息体。这种方式适合于需要处理不同长度消息的场景，常用于高效的二进制协议中。
*/

// 我们可以自己定义一个协议，比如数据包的前4个字节为包头，里面存储的是发送的数据的长度。

// Encode 将消息编码
func Encode(message string) ([]byte, error) {
	// 读取消息的长度，转换成int32类型（占4个字节）, 此时设定了消息头为4字节
	length := int32(len(message))
	pkg := new(bytes.Buffer) // bytes.Buffer 是 Go 标准库中的一个用于缓冲字节的容器，它提供了高效的写入和读取操作。
	// 写入消息头，消息头4字节
	err := binary.Write(pkg, binary.LittleEndian, length) // 使用 binary.Write 函数以 小端序（binary.LittleEndian）的格式写入数据。小端序意味着低位字节在前，高位字节在后。不同的系统可能会使用不同的字节序，但在自定义协议中，为了保持一致性，通常明确指定字节序。
	if err != nil {
		return nil, err
	}
	// 写入消息实体，消息体另外占空间
	err = binary.Write(pkg, binary.LittleEndian, []byte(message))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(), nil
	/*
		将整个缓冲区 pkg 转换为字节数组并返回。这个字节数组包含了完整的消息，格式为：[消息头（长度）][消息体（内容）]。
		接收方可以根据消息头中的长度信息知道消息体的大小，从而读取正确的字节数，防止粘包问题。
	*/
}

// Decode 将消息解码
func Decode(reader *bufio.Reader) (string, error) {
	// 读取消息的长度
	lengthByte, _ := reader.Peek(4)                              // 使用 Peek(4) 读取前 4 个字节，但不移动读取指针。Peek 方法仅查看数据，方便检查是否有足够的数据处理。
	lengthBuff := bytes.NewBuffer(lengthByte)                    // 这行代码的作用是将一个字节切片 lengthByte 封装为一个 字节缓冲区对象
	var length int32                                             // length 准备存储的值，是从包头解码出来的整个包体的长度？
	err := binary.Read(lengthBuff, binary.LittleEndian, &length) // 这段代码的作用是使用 binary.Read 函数将 lengthBuff 缓冲区中的 字节数据 按照 小端字节序（Little Endian） 解码为 int32 类型的整数，并将其存储在变量 length 中。
	if err != nil {
		return "", err
	}

	// Buffered返回缓冲中现有的可读取的字节数。
	// 使用 Buffered() 函数检查缓冲区中是否有足够的数据。如果缓冲区中的数据不足 length+4 个字节（即消息头加消息体），则说明还没收到完整的数据包，解码不能继续。
	if int32(reader.Buffered()) < length+4 {
		return "", err
	}

	// 读取真正的消息数据
	// 创建一个字节数组 pack，长度为消息头加消息体的总长度（4+length 个字节），用于存放完整的消息。
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack) // 读取整个消息（包括消息头和消息体）到 pack 中。
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil // 跳过前 4 个字节的消息头，返回剩下的消息体部分，即实际的消息内容。
}
