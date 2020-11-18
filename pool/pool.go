package pool


// New create a goroutine pool
func New(size int) Pool {
	if size < 0 {
		panic("go goroutine pool size should not less than 0")
	}
	if size == 0 {
		size = 1
	}
	return pool{
		buffer: make(chan int, size),
	}
}

//  Pool goroutine pool interface
type Pool interface {
	Run(func())
}

type pool struct {
	locker sync.Locker
	buffer chan int
}

// Run get a goroutine from pool, if zhe pool is exhausted
// it will wait a free a goroutine to run this function
func (p pool) Run(f func()) {
	p.buffer <- 1
	go func() {
		defer func() {
			<-p.buffer
		}()
		f()
	}()
}
