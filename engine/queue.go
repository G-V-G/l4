package engine

import "sync"

// Queue - main commands stack
type Queue struct {
	sync.Mutex
	Commands []Command
	onReceiveEmptyChan chan bool
	onEmptyReceive bool
}

// Push adds given command to the stacks tail
func (q *Queue) Push(cmd Command) {
	q.Lock()
	defer q.Unlock()
	q.Commands = append(q.Commands, cmd)
	if q.onEmptyReceive {
		q.onEmptyReceive = false
		q.onReceiveEmptyChan <- true
	}
}

// Pull pops command from the stack
func (q *Queue) Pull() Command {
	q.Lock()
	defer q.Unlock()

	if len(q.Commands) == 0 {
		q.onEmptyReceive = true
		q.Unlock()
		<- q.onReceiveEmptyChan
		q.Lock()
	}

	cmd := q.Commands[0]
	q.Commands[0] = nil
	q.Commands = q.Commands[1:]
	return cmd
}
