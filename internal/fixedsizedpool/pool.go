package fixedsizedpool

import (
	"fmt"
	"sync/atomic"
)

// Pool represents a generic fixedsizedpool of objects of type T.
//
// The fixedsizedpool uses a buffered channel to manage memory efficiently, allowing for reuse of pre-existing objects
// without requiring frequent allocations and de-allocations. The type T defines the type of objects stored
// within the fixedsizedpool.
type Pool[T any] struct {

	// The total capacity of the fixedsizedpool (maximum number of objects it can hold).
	size uint32

	// The number of currently avail objects in the fixedsizedpool.
	avail uint32

	// The underlying buffered channel used to manage the fixedsizedpool's objects.
	objectPoolChan chan T
}

// Get retrieves an object of type T from the fixedsizedpool.
//
// This method blocks if the buffered channel is empty, waiting until an object is available.
// It reduces the availability count atomically to maintain thread-safety.
//
// # Returns
//
//   - T: An available object of type T retrieved from the fixedsizedpool.
func (p *Pool[T]) Get() T {

	// Block here until there is something is an object of type T
	// in the buffer.
	obj := <-p.objectPoolChan

	// Atomically decrease the value of p.avail by 1 to ensure
	// thread-safety
	atomic.AddUint32(&p.avail, ^uint32(0))

	return obj
}

// Put adds an object of type T back into the fixedsizedpool.
//
// This method blocks if the buffered channel is full, waiting until there is space for the object.
// It also ensures thread-safety by atomically increasing the availability count.
//
// # Params
//
//   - t T: The object to be added back into the fixedsizedpool.
//
// # Returns
//
//   - error: Returns an error if the fixedsizedpool has reached its maximum capacity.
func (p *Pool[T]) Put(t T) error {

	// Fail if sending t into the buffered channel will block;
	// It will block if there the amount of T in the buffered channel
	// has reached the buffer size.
	if p.avail == p.size {
		return fmt.Errorf("fixedsizedpool is full. Cannot add more objects")
	}

	// Send the object into the buffered channel.
	p.objectPoolChan <- t

	// Atomically increase the value of p.avail by 1 to ensure
	// thread-safety
	atomic.AddUint32(&p.avail, 1)

	return nil
}

// Avail returns the number of currently available objects in the fixedsizedpool.
//
// This method uses atomic operations to ensure thread-safety when reading the availability count.
//
// # Returns
//
//   - uint32: The number of available objects currently present in the fixedsizedpool.
func (p *Pool[T]) Avail() uint32 {

	// Atomically read the value from avail to ensure
	// thread-safety.
	return atomic.LoadUint32(&p.avail)
}

// Teardown performs cleanup operations for the fixed-sized pool.
//
// This method ensures that any resources held by the pool, such as objects in the channel,
// are released, and the channel is safely closed. It prepares the pool for proper disposal
// to prevent memory leaks or unexpected behavior.
//
// This method should be called when the pool is no longer needed.
//
// # Returns
//
//   - error: Returns an error if the teardown process encounters issues, such as a failure
//     to release resources or close the channel properly.
func (p *Pool[T]) Teardown() error {

	for atomic.LoadUint32(&p.avail) > 0 {
		_ = p.Get()
	}
	close(p.objectPoolChan)

	return nil
}
