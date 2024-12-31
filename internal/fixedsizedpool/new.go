package fixedsizedpool

// New creates a new instance of Pool with the specified size.
//
// The fixedsizedpool is implemented using a buffered channel to allow non-blocking send operations until the buffer is full,
// and it blocks retrieval operations (`Get`) until an item is available in the fixedsizedpool.
//
// The generic parameter T denotes the type of objects stored in the fixedsizedpool.
//
// # Params
//
//   - size uint32: The capacity of the fixedsizedpool, determining how many objects can be held simultaneously.
//
// # Returns
//
//   - *Pool[T]: A pointer to the created Pool instance.
func New[T any](size uint32) *Pool[T] {

	// Create a buffered channel with buffered amount equal
	// to the value of the size variable. This will allow
	// non-blocking send operations on the channel; We do this
	// as using a non-buffered channel will block on sending in
	// a value until there is something waiting to read a value
	// from it.
	return &Pool[T]{
		size:           size,
		avail:          uint32(0),
		objectPoolChan: make(chan T, size),
	}
}
