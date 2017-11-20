package file

import (
    ab "github.com/hyperledger/fabric/protos/orderer"
    "github.com/golang/protobuf/proto"
    
    "sync"
    "os"
    "encoding/binary"
    "io"
)

type FileStorer struct {
    storeFilePath string
    storeFile     *os.File
    metux         sync.Mutex
}

func NewFileStore(storePath string) *FileStorer {
    
    return &FileStorer{
        storeFilePath: storePath,
        metux:         sync.Mutex{},
    }
}

func (filestore * FileStorer) StoreMessages(msgs map[uint64]*ab.BftSmartMessage) error {
    
    filestore.metux.Lock()
    
    var err error
    filestore.storeFile, err = filestore.GetOrCreate(filestore.storeFilePath)
    if err != nil {
        return err
    }
    
    for _, msg := range msgs {
        msgBytes, _ := proto.Marshal(msg)
        filestore.append(msgBytes)
    }
    
    filestore.metux.Unlock()

    return nil
}

func (filestore * FileStorer) GetAllMessage(recv chan *ab.BftSmartMessage) error {
    
    filestore.metux.Lock()
    
    var err error
    filestore.storeFile, err = filestore.GetOrCreate(filestore.storeFilePath)
    if err != nil {
        return err
    }
    
    for {
        buf := make([]byte, 8)
        _, err = filestore.storeFile.Read(buf)
        if err == io.EOF {
            close(recv)
            break
        } else {
            continue
        }
    
        byteLen := bytesToInt(buf)
        msgByte := make([]byte, byteLen)
        filestore.storeFile.Read(msgByte)
    
        var msg ab.BftSmartMessage
        proto.Unmarshal(msgByte, &msg)
        
        recv <- &msg
    }
    
    filestore.metux.Lock()
    
    return nil
}


func (filestore * FileStorer) append(msg []byte) {
    
    lenByte := intToBytes(len(msg))
    
    filestore.storeFile.Write(lenByte)
    filestore.storeFile.Write(msg)
}

func (filestore * FileStorer) GetOrCreate(filepath string) (*os.File, error) {
    
    _, err := os.Stat(filepath)
    if os.IsNotExist(err) {
        os.Create(filepath)
    } else {
        return nil, err
    }
    
    f, err := os.Open(filepath)
    if err != nil {
        return nil, err
    }
    
    return f, nil
}

func intToBytes(value int) []byte {
    
    var buf [8]byte
    
    binary.BigEndian.PutUint64(buf[:], uint64(value))
    
    return buf[:]
}

func bytesToInt(value []byte) uint64 {
    return binary.BigEndian.Uint64(value)
}