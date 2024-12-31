package pool

// Pool represents a generic interface for an object fixedsizedpool of type T.
//
// The Pool interface provides methods for retrieving, returning, and checking
// the availability of objects in the fixedsizedpool. This allows the reuse of pre-allocated
// objects, reducing the overhead of repeated object creation and garbage collection.
//
// Type Parameters:
//   - T: The type of objects that the fixedsizedpool will manage.
type Pool[T any] interface {

	// Get retrieves an object of type T from the fixedsizedpool.
	//
	// This method should block if no objects are available in the fixedsizedpool until an object
	// becomes available.
	//
	// # Returns
	//
	//   - T: An object of type T retrieved from the fixedsizedpool.
	Get() T

	// Put adds an object of type T back into the fixedsizedpool.
	//
	// This method should block if the fixedsizedpool is already full, only allowing additions
	// when space in the fixedsizedpool becomes available.
	//
	// # Params
	//
	//   - t T: The object to be returned to the fixedsizedpool.
	//
	// # Returns
	//
	//   - error: Returns an error if the object could not be added to the fixedsizedpool,
	//            typically because the fixedsizedpool is full.
	Put(t T) error

	// Avail returns the number of currently available objects in the fixedsizedpool.
	//
	// This method should provide a thread-safe way to query the fixedsizedpool for the count
	// of reusable objects currently held.
	//
	// # Returns
	//
	//   - uint32: The number of available objects in the fixedsizedpool.
	Avail() uint32

	// Teardown cleans up any resources or performs any necessary finalization
	// for the pool.
	//
	// This method should release any internal resources held by the pool,
	// such as open channels, and prepare it for safe disposal.
	//
	// This is typically called when the pool is no longer needed.
	//
	// # Returns
	//
	//   - error: Returns an error if the teardown process fails or encounters
	//            any issues during cleanup.
	Teardown() error
}
