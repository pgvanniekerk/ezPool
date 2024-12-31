package pool

import (
	"github.com/pgvanniekerk/ezPool/internal/fixedsizedpool"
)

// NewFixedSizedPool creates a new instance of a Pool with the specified size.
//
// The created fixedsizedpool manages reusable objects of type T to minimize the overhead associated
// with frequent object creation and garbage collection. The fixedsizedpool allows for reusing pre-allocated
// objects, improving performance in workloads that require repeated allocation and deallocation.
//
// Type Parameters:
//   - T: The type of objects that the fixedsizedpool will manage.
//
// # Params
//
//   - size uint32: The capacity of the fixedsizedpool, determining the maximum number of objects
//     that can be held in the fixedsizedpool simultaneously.
//
// # Returns
//
//   - Pool[T]: An instance of the Pool interface for managing objects of type T.
func NewFixedSizedPool[T any](size uint32) Pool[T] {
	return fixedsizedpool.New[T](size)
}
