package ringbuffer

type Data struct {
	ID      string
	Content []byte
}

var (
	IN_OUT_CHANNEL_BUFFER int
)

type Option func(*RingBuffer)

type RingBuffer struct {
	InBuffer  chan interface{}
	OutBuffer chan interface{}
}

func New(option ...Option) *RingBuffer {
	rb := &RingBuffer{}
	for _, o := range option {
		o(rb)
	}
	return rb
}

func WithInAndOutBufferSize(bs int) Option {
	return func(rb *RingBuffer) {
		rb.InBuffer = make(chan interface{}, bs)
		rb.OutBuffer = make(chan interface{}, bs)
	}
}

func (rb *RingBuffer) Start() {
	go func() {
		for {
			select {
			case in := <-rb.InBuffer:
				rb.OutBuffer <- in
			default:
				//
			}

			if len(rb.OutBuffer) == IN_OUT_CHANNEL_BUFFER {
				<-rb.OutBuffer
			}
		}
	}()
}
