/*
@Time : 2020/2/17 14:48
@Author : Minus4
*/
package snowflake

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type Worker struct {
	mu             sync.Mutex
	epoch          time.Time
	lastTime       int64
	backupLastTime int64
	workerId       int64
	step           int64
	workerMax      int64
	workerMask     int64
	stepMask       int64
	timeShift      uint8
	workerShift    uint8
}

var (
	Epoch         int64 = 1288834974657
	NodeBits      uint8 = 10
	StepBits      uint8 = 12
	MaxBackwardMs int64 = time.Millisecond.Milliseconds() * 5
)
var ErrClockBackward = errors.New("system clock move backward")

func NewWorker(workerId int64) (*Worker, error) {

	n := &Worker{}
	n.workerId = workerId
	n.workerMax = -1 ^ (-1<<NodeBits)>>1
	n.workerMask = n.workerMax << StepBits
	n.stepMask = -1 ^ (-1 << StepBits)
	n.timeShift = NodeBits + StepBits
	n.workerShift = StepBits
	n.mu = sync.Mutex{}
	if n.workerId < 0 || n.workerId > n.workerMax {
		return nil, errors.New("worker number must be between 0 and " + strconv.FormatInt(n.workerMax, 10))
	}

	var curTime = time.Now()
	// add time.Duration to curTime to make sure we use the monotonic clock if available
	n.epoch = curTime.Add(time.Unix(Epoch/1e3, (Epoch%1e3)*1e6).Sub(curTime))
	return n, nil
}

func (w *Worker) Generate() (int64, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	curTimeStamp := time.Since(w.epoch).Milliseconds()

	if w.lastTime > curTimeStamp {
		if w.lastTime-curTimeStamp > MaxBackwardMs {
			time.Sleep(time.Duration(w.lastTime - curTimeStamp))
		} else {
			var interval int64 = 512
			if w.workerId >= interval {
				w.workerId -= interval
			} else {
				w.workerId += interval
			}
			tempLastTime := w.backupLastTime
			w.backupLastTime = w.lastTime
			w.lastTime = tempLastTime
			if w.lastTime > curTimeStamp {
				return 0, ErrClockBackward
			}
		}
	}

	if w.lastTime == curTimeStamp {
		w.step = (w.step + 1) & w.stepMask

		if w.step == 0 {
			for curTimeStamp <= w.lastTime {
				// wait for next ms
				curTimeStamp = time.Since(w.epoch).Milliseconds()
			}
		}
	} else {
		// optimization for hash
		w.step = 127 & curTimeStamp
	}

	w.lastTime = curTimeStamp

	res := curTimeStamp<<w.timeShift | w.workerId<<w.workerShift | w.step
	return res, nil
}

func (w Worker) RedisScoreMapping(snowflakeId int64) int64 {
	timePart := snowflakeId >> w.timeShift
	stepPart := (snowflakeId & w.stepMask) >> 1
	return timePart<<(StepBits-1) | stepPart
}
