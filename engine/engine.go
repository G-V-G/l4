package engine

import "sync"

type onFinishFn func (h Handler)

func (fn onFinishFn) Execute (h Handler) {
	fn(h)
}

// Engine - loop handler
type Engine struct {
	sync.Mutex
	onFinishChan chan bool
	stopRequest bool
	cmdStack *Queue
}


// Start runs main loop
func (e *Engine) Start() {
  e.cmdStack = &Queue{onReceiveEmptyChan: make(chan bool)}
  e.onFinishChan = make(chan bool)
  go func() {
    for {
      cmd := e.cmdStack.Pull(&e.stopRequest)
      if cmd == nil {
        break
      }
      cmd.Execute(e)
    }
    e.onFinishChan <- true
  }()
}

// Post adds given command to the stack
func (e *Engine) Post(cmd Command) {
	e.cmdStack.Push(cmd)
}

// AwaitFinish stops pushing commans to the queue
func (e *Engine) AwaitFinish() {
	var finishCommand onFinishFn = func (_ Handler) {
		e.stopRequest = true
	}
	e.Post(finishCommand)
	<-e.onFinishChan
}
