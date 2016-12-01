package service

import (
	"bytes"
	"fmt"
	"runtime"
	"strconv"
	"sync"

	"stathat.com/c/consistent"

	"github.com/aiyun/gomqtt/proto"
	disruptor "github.com/smartystreets/go-disruptor"
	"github.com/sunface/tools"
	"github.com/uber-go/zap"
)

const (
	CACHE_INSERT = 1000 // 插入
	CACHE_GET    = 2000 // 获取
	CACHE_SELECT = 3000 // 查询
	CACHE_DELETE = 4000 // 删除
)

var gQueue map[string]*Controller
var gQueueHash *consistent.Consistent

func InitQueue(queueNum int) {
	gQueue = make(map[string]*Controller)
	gRetChan = make(chan *CacheRet, 100)
	gQueueHash = consistent.New()
	if queueNum == 0 {
		queueNum = runtime.NumCPU()
	}

	Logger.Info("InitQueue", zap.Int("queueNum", queueNum))

	for index := 0; index < queueNum; index++ {
		controller := NewController()
		controller.Init(65536, 65535, 1)
		controller.Start()
		gQueue[fmt.Sprintf("%d", index)] = controller
		gQueueHash.Add(fmt.Sprintf("%d", index))
	}
}

func GetQueue(acc []byte) (*Controller, error) {
	index, err := gQueueHash.Get(tools.Bytes2String(acc))
	if err != nil {
		return nil, err
	}
	Logger.Info("Hash", zap.String("Acc", tools.Bytes2String(acc)), zap.String("index", index))
	c, ok := gQueue[index]
	if !ok {
		return nil, fmt.Errorf("unfind index %s in map ", index)
	}

	return c, nil
}

func CloseQueue() {
	for index, controller := range gQueue {
		controller.Start()
		delete(gQueue, index)
	}
}

type CacheTask struct {
	MsgTy   int
	RetChan chan *CacheRet
	Acc     []byte
	Topic   []byte
	Msg     *proto.PubTextMsg // []byte
	MsgIDs  [][]byte
}

type CacheRet struct {
	MsgTy  int
	Acc    []byte
	Topic  []byte
	MsgIDs *TopicIDMap
	Data   []byte
}

type BufPool struct {
	Data CacheTask
}

type Controller struct {
	controller   disruptor.Disruptor
	ring         []CacheTask
	bufferMask   int64
	reservations int64
	msgIDManger  *MsgIdManger // 消息msgid缓存
	msgCache     *MsgCache    // 消息缓存
}

func NewController() *Controller {
	c := &Controller{
		msgIDManger: NewMsgIdManger(),
		msgCache:    NewMsgCache(),
	}
	return c
}

func (c *Controller) Init(bufferSize int64, bufferMask int64, reservations int64) {
	c.controller = disruptor.Configure(bufferSize).
		WithConsumerGroup(Writer{queue: c, bufPool: &sync.Pool{}}).Build()
	c.ring = make([]CacheTask, bufferSize)
	c.bufferMask = bufferMask
	c.reservations = reservations
}

func (c *Controller) Start() {
	c.controller.Start()
}

func (c *Controller) Close() error {
	c.controller.Stop()
	return nil
}

func (c *Controller) Publish(data CacheTask) {
	sequence := disruptor.InitialSequenceValue
	writer := c.controller.Writer()

	sequence = writer.Reserve(c.reservations)
	for lower := sequence - c.reservations + 1; lower <= sequence; lower++ {
		c.ring[lower&c.bufferMask] = data
	}
	// 提交
	writer.Commit(sequence-c.reservations+1, sequence)
}

// ///////////

type Writer struct {
	queue   *Controller
	bufPool *sync.Pool
}

func (this Writer) Consume(lower, upper int64) {
	for lower <= upper {
		var bufPool *BufPool
		buf := this.bufPool.Get()
		if buf == nil {
			bufPool = &BufPool{}
		} else {
			bufPool = buf.(*BufPool)
		}
		bufPool.Data = this.queue.ring[lower&this.queue.bufferMask]
		// 消费
		switch bufPool.Data.MsgTy {
		case CACHE_INSERT:
			// 插入数据
			this.queue.msgCache.Insert(bufPool.Data.Msg.Mid, bufPool.Data.Msg.Msg)
			this.queue.msgIDManger.InsertTextMsgID(bufPool.Data.Msg)
			Logger.Info("Consume", zap.String("ToAcc", tools.Bytes2String(bufPool.Data.Msg.ToAcc)), zap.String("Topic", tools.Bytes2String(bufPool.Data.Msg.Ttp)), zap.String("Msg", tools.Bytes2String(bufPool.Data.Msg.Msg)))
			break
		case CACHE_GET:
			Logger.Info("CACHE_GET", zap.String("Acc", tools.Bytes2String(bufPool.Data.Acc)), zap.String("Topic", tools.Bytes2String(bufPool.Data.Topic)))
			ret := &CacheRet{}
			data, ok := this.queue.msgCache.Get(bufPool.Data.MsgIDs[0])
			if ok {
				ret.Data = data
			}
			bufPool.Data.RetChan <- ret
			break
		case CACHE_SELECT:
			Logger.Info("Select", zap.String("Acc", tools.Bytes2String(bufPool.Data.Acc)), zap.String("Topic", tools.Bytes2String(bufPool.Data.Topic)))
			// 查询返回msgid信息
			tm := this.queue.msgIDManger.GetMsgIDs(bufPool.Data.Acc, bufPool.Data.Topic)
			ret := &CacheRet{
				MsgIDs: tm,
			}
			bufPool.Data.RetChan <- ret
			// gRetChan <- ret
			break
		case CACHE_DELETE:
			this.queue.msgCache.Delete(bufPool.Data.MsgIDs)
			break
		}

		// <-bufPool.Data.RetChan
		// fmt.Println(<-bufPool.Data.RetChan)
		this.bufPool.Put(buf)
		lower++
	}
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
