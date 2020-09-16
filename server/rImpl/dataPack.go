/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-10-30 10:21
 */
package rImpl

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"runx/server/rIterface"
)

type DataPack struct{}

func NewDataPack() *DataPack {
	d := &DataPack{}
	return d
}

//获取包的长度
func (d *DataPack) GetHeadLen() uint64 {
	//DataLen + DataID
	return 16
}

//封包方法
func (d *DataPack) PackData(m rIterface.IMessage) ([]byte, error) {

	buf := bytes.NewBuffer([]byte{})
	if err := binary.Write(buf, binary.LittleEndian, m.GetMessageLen()); err != nil {

		return nil, errors.New("DataPack err : pack len ")
	}
	if err := binary.Write(buf, binary.LittleEndian, m.GetMessageId()); err != nil {

		return nil, errors.New("DataPack err : pack messageId")
	}
	if err := binary.Write(buf, binary.LittleEndian, m.GetData()); err != nil {

		return nil, errors.New("DataPack err : pack data")
	}
	return buf.Bytes(), nil
}

//拆包方法
func (d *DataPack) UnPackData(bytesData []byte) (m rIterface.IMessage, err error) {

	buf := bytes.NewReader(bytesData)
	message := &Message{}
	if err := binary.Read(buf, binary.LittleEndian, &message.MessageLength); err != nil {
		return nil, errors.New(fmt.Sprintf("UnDataPack err : read messageLen === %s\n", err))
	}

	if err := binary.Read(buf, binary.LittleEndian, &message.MessageId); err != nil {
		return nil, errors.New(fmt.Sprintf("UnDataPack err : read messageId === %s\n", err))
	}

	return message, nil
}
