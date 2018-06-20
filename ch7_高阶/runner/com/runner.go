package com

import (
	"time"
	"os"
	"errors"
	"os/signal"
)

/**
	runner包用于展示如何使用通道来监视程序的执行时间，
	如果程序运行时间太长，也可以 用runner包来终止程序。
	当开发需要调度后台处理任务的程序的时候，这种模式会很有用。
	这 个程序可能会作为 cron 作业执行，或者在基于定时任务的云环境（如 iron.io）里执行。
 */

type Runner struct {
	tasks     []func(int)
	complete  chan error
	timeout   <-chan time.Time
	interrupt chan os.Signal
}

func New(tm time.Duration) *Runner {
	return &Runner{
		complete:  make(chan error),
		timeout:   time.After(tm),
		interrupt: make(chan os.Signal, 1),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

var ErrTimeOut = errors.New("执行者执行超时")
var ErrInterrrupt = errors.New("执行者被中断")

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.isInterrupt() {
			return ErrInterrrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) isInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	default:
		return false
	}
}

func (r *Runner) Start() error {
	signal.Notify(r.interrupt,os.Interrupt)
	go func() {
		r.complete<-r.run()
	}()

	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeOut
	}
}