package com

import (
	"sync"
	"io"
	"errors"
	"log"
)

type Pool struct {
	m       sync.Mutex
	res     chan io.Closer
	factory func() (io.Closer, error)
	closed  bool
}

var ErrPoolClosed = errors.New("资源池已经关闭")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size的值太小了")
	}
	return &Pool{
		factory: fn,
		res:     make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := p.res:
		log.Println("Acquire:共享资源")
		if !ok {
			return nil,ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire:新生成资源")
		return p.factory()
	}
}

func (p *Pool) Close() {
	p.m.Lock()
	defer p.m.Unlock()

	if p.closed {
		return
	}

	p.closed= true
	close(p.res)
	for r := range p.res {
		r.Close()
	}
}

func (p *Pool) Release(r io.Closer) {
	p.m.Lock()
	defer p.m.Unlock()
	if p.closed {
		r.Close()
		return
	}
	select {
	case p.res <- r:
		log.Println("资源释放到池子里了")
	default:
		log.Println("资源池满了，释放这个资源吧")
		r.Close()
	}
}