package ypipeline

import "sync"

/**
------------------------------------------------
Created on 2022-11-07 15:53
@Author: ZhangYundi
@Email: yundi.xxii@outlook.com
------------------------------------------------
**/

type (
	Worker func(in interface{}) (out interface{})
)

type Pipeline struct {
	entry       Worker // 步骤1
	final       Worker // 最后的步骤
	workflows   []Worker
	workingChan chan struct{} // 正在工作的数量
	inWg        *sync.WaitGroup
	inCache     chan interface{} // 参数缓存
	exitFlag    bool
	threadWg    *sync.WaitGroup
}

func NewPipelines(parallelSize int, workers ...Worker) *Pipeline {
	newWorkers := make([]Worker, 0)
	for _, w := range workers {
		if w != nil {
			newWorkers = append(newWorkers, w)
		}
	}
	var entry, final Worker
	if len(newWorkers) > 0 {
		entry = newWorkers[0]
		final = newWorkers[len(newWorkers)-1]
	}
	pipeline := &Pipeline{
		entry:       entry,
		final:       final,
		workflows:   newWorkers,
		workingChan: make(chan struct{}, parallelSize),
		inWg:        &sync.WaitGroup{},
		threadWg:    &sync.WaitGroup{},
		inCache:     make(chan interface{}, parallelSize),
	}
	return pipeline
}

// 开始第一步的任务
func (pl *Pipeline) start(in interface{}) {
	defer func() {
		pl.threadWg.Done()
		<-pl.workingChan
	}()
	pl.threadWg.Add(1)
	for !pl.exitFlag {
		if len(pl.workflows) >= 1 {
			var out interface{}
			out = pl.entry(in)
			for _, w := range pl.workflows[1:] {
				in = out
				out = w(in)
			}
		}

		pl.inWg.Done()
		in = <-pl.inCache
	}
}

func (pl *Pipeline) AddTask(ins ...interface{}) {
	for _, in := range ins {
		pl.inWg.Add(1)
		select {
		case pl.workingChan <- struct{}{}:
			go pl.start(in)
		default:
			pl.inCache <- in
		}
	}
	pl.inWg.Wait() // 等待所有的参数进入 working
	close(pl.inCache)
	pl.exitFlag = true
	pl.threadWg.Wait() // 等待所有的 thread 完毕
	close(pl.workingChan)
}
