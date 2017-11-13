package bftsmart

import (
    "net"
    "sync"
    "fmt"
    "encoding/binary"
    "io"
    
    cb "github.com/hyperledger/fabric/protos/common"
    ab "github.com/hyperledger/fabric/protos/orderer"
    "github.com/hyperledger/fabric/protos/utils"
    "github.com/golang/protobuf/proto"
)

type channel struct {
    channelId    string
    chain        chain
}

type connection struct {
    sendProxy    net.Conn
    recvProxy    net.Conn
    sendPool     []net.Conn
    mutex        []*sync.Mutex
    channels     map[string]channel
}


func newConnection() (*connection, error) {
    
    conn := connection{}
    err := conn.New()
    if err != nil {
        return nil, err
    }
    
    return &conn, nil
}

//deal with the error
func (c *connection) New() error {
    logger.Debugf("start to connect to BFTProxy")
    conn, err := net.Dial("unix", "/tmp/bft.sock")
    
    if err != nil {
        logger.Debugf("Could not connect to proxy by unix socket!")
        return err
    } else {
        logger.Debugf("Connected to proxy by unix socket!")
    }
    
    c.sendProxy = conn
    
    addr := fmt.Sprintf("localhost:%d", recvport)
    conn, err = net.Dial("tcp", addr)
    
    if err != nil {
        logger.Debugf("Could not connect to proxy!")
        return err
    } else {
        logger.Debugf("Connected to proxy %s!", addr)
    }
    
    c.recvProxy = conn
    
    _, err = c.sendUint32(uint32(poolsize))
    
    if err != nil {
        logger.Debugf("Error while sending pool size:", err)
        return err
    }
    
    c.sendPool = make([]net.Conn, poolsize)
    c.mutex = make([]*sync.Mutex, poolsize)
    
    //create connection pool
    logger.Debugf("the sendpool size is %d,and the mutex size is %d", poolsize)
    for i := uint(0); i < poolsize; i++ {
        
        conn, err := net.Dial("unix", "/tmp/bft.sock")
        
        if err != nil {
            logger.Panicf("Could not create connection %v: %d\n", i, err)
            //return
        } else {
            logger.Debugf("Created connection: %v\n", i)
            c.sendPool[i] = conn
            c.mutex[i] = &sync.Mutex{}
        }
    }
    
    c.channels = make(map[string]channel)
    
    go c.connLoop() //JCS: my own loop
    return nil
}

func (ch *connection) sendLength(length int, conn net.Conn) (int, error) {
    
    var buf [8]byte
    
    binary.BigEndian.PutUint64(buf[:], uint64(length))
    
    return conn.Write(buf[:])
}

func (ch *connection) sendUint64(length uint64) (int, error) {
    
    var buf [8]byte
    
    binary.BigEndian.PutUint64(buf[:], uint64(length))
    
    return ch.sendProxy.Write(buf[:])
}

func (ch *connection) sendUint32(length uint32) (int, error) {
    
    var buf [4]byte
    
    binary.BigEndian.PutUint32(buf[:], uint32(length))
    
    return ch.sendProxy.Write(buf[:])
}

func (ch *connection) sendEnvToBFTProxy(env *ab.BftSmartMessage, index uint) (int, error) {
    
    ch.mutex[index].Lock()
    bytes, err := utils.Marshal(env)
    
    if err != nil {
        return -1, err
    }
    
    status, err := ch.sendLength(len(bytes), ch.sendPool[index])
    
    if err != nil {
        return status, err
    }
    
    i, err := ch.sendPool[index].Write(bytes)
    
    ch.mutex[index].Unlock()
    return i, err
}

func (ch *connection) recvLength() (int64, error) {
    
    var size int64
    err := binary.Read(ch.recvProxy, binary.BigEndian, &size)
    return size, err
}

func (ch *connection) recvBytes() ([]byte, error) {
    
    size, err := ch.recvLength()
    
    if err != nil {
        return nil, err
    }
    
    buf := make([]byte, size)
    
    _, err = io.ReadFull(ch.recvProxy, buf)
    
    if err != nil {
        return nil, err
    }
    
    return buf, nil
}

func (ch *connection) recvEnvFromBFTProxy() (*cb.Envelope, error) {
    
    size, err := ch.recvLength()
    
    if err != nil {
        return nil, err
    }
    
    buf := make([]byte, size)
    
    _, err = io.ReadFull(ch.recvProxy, buf)
    
    if err != nil {
        return nil, err
    }
    
    env, err := utils.UnmarshalEnvelope(buf)
    
    if err != nil {
        return nil, err
    }
    
    return env, nil
}


func (ch *connection) connLoop() {
    
    for {
        
        //TODO deal with the err
        bytes, err := ch.recvBytes()
        if err != nil {
            logger.Debugf("[recv] Error while receiving block from BFT proxy: %v\n", err)
            //TODO recvProxy maybe down, try to reconnect to bftproxy
            continue
        }
        
        var bftMessage ab.BftSmartProxy
        err = proto.Unmarshal(bytes, &bftMessage)
        if err != nil {
            logger.Debugf("[recv] Unmarshaled BftMessage failed:", err.Error())
        }
        
        channel, ok := ch.channels[bftMessage.BfgMsg.ChannelId]
        if !ok {
            logger.Warningf("The Channel %s is not existed", bftMessage.BfgMsg.ChannelId)
            continue
        }
        
        channel.chain.recvChan <- &bftMessage
    }
}



