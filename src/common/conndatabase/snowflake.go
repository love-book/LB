package conndatabase

import(
	"errors"
	"fmt"
	"sync"
	"time"
)

const (
	twepoch = int64(1417937700000)
	DistrictIdBits = uint64(5)
	NodeIdBits = uint64(5)
	maxNodeId = -1 ^ (-1 << NodeIdBits)
	maxDistrictId = -1 ^ (-1 << DistrictIdBits)
	sequenceBits = uint64(12)
	nodeIdShift = sequenceBits
	DistrictIdshift = sequenceBits + NodeIdBits
	timestampLeftShift = sequenceBits + NodeIdBits + DistrictIdBits
	sequenceMask = -1 ^ (-1 << sequenceBits)
	maxNextIdsNum = 50
)

type IdWorker struct {
	sequence int64
	lastTimestamp int64
	nodeId int64
	twepoch int64
	districtId int64
	mutex sync.Mutex
}

func NewIdWorker(NodeId int64,districtId int64) (* IdWorker, error) {
	idworker := &IdWorker{}

	if NodeId > maxNodeId || NodeId < 0 {
		fmt.Errorf("Nodeid  id can't be greater than %d or less than 0",maxNodeId)
		return nil, errors.New(fmt.Sprintf("Nodeid id :%d error",NodeId))
	}

	if districtId > maxDistrictId || districtId < 0 {
		fmt.Sprintf("District id  can't be greater than %d or less than 0",maxDistrictId)
		return nil, errors.New(fmt.Sprintf("Nodeid id :%d error",districtId))
	}

	idworker.nodeId = NodeId
	idworker.districtId = districtId
	idworker.lastTimestamp = -1
	idworker.sequence = 0
	idworker.twepoch = twepoch
	idworker.mutex = sync.Mutex{}
	fmt.Sprintf("worker starting. timestamp left shift %d, District id bits %d, worker id bits %d, sequence bits %d, workerid %d", timestampLeftShift, DistrictIdBits, NodeIdBits, sequenceBits, NodeId)
	return idworker,nil
}

func timeGen() int64  {
	return time.Now().UnixNano()
}

func tilNexMillis(lastTimestamp int64) int64  {
	timestamp:= timeGen()
	for timestamp <= lastTimestamp {
		timestamp = timeGen()
	}

	return timestamp
}

func (id * IdWorker) NextId ()(int64,error)  {
	defer id.mutex.Unlock()

	id.mutex.Lock()
	/*ids,_  := id.nextid()
	for ids  < 0 {
		ids,_ = id.nextid()
	}*/
	return id.nextid()
}

func (id * IdWorker) NextIds(num int) ([]int64,error)   {
	defer  id.mutex.Unlock()

	if num > maxNextIdsNum {
		fmt.Sprintf("NextIds num can't be greater than %d or less than 0", maxNextIdsNum)
		return nil, errors.New(fmt.Sprintf("NextIds num: %d error", num))
	}
	ids := make ([]int64, num)
	id.mutex.Lock()
	for i := 0 ; i< num ; i++{
		ids[i] , _ = id.nextid()
	}
	return  ids,nil
}

func (id * IdWorker) nextid () (int64,error){
	timestamp := timeGen()
	if timestamp < id.lastTimestamp {
		return  0 , errors.New(fmt.Sprintf("Clock moved backwards.  Refusing to generate id for %d milliseconds", id.lastTimestamp-timestamp))
	}

	if id.lastTimestamp == timestamp {
		id.sequence = (id.sequence + 1) & sequenceMask
		if id.sequence == 0 {
			timestamp = tilNexMillis(id.lastTimestamp)
		}
	}else{
		id.sequence = 0
	}

	id.lastTimestamp = timestamp
	return ((timestamp - id .twepoch) << timestampLeftShift) | (id.districtId << DistrictIdshift) | (id.nodeId << nodeIdShift) | id.sequence,nil
}













