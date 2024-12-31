package fixedsizedpool

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the test suite
type PoolTestSuite struct {
	suite.Suite
	pool *Pool[int] // Test instance of the Pool
}

// SetupTest initializes a new instance of Pool
// between each test run.
func (suite *PoolTestSuite) SetupTest() {

	// Create a Fixed Sized Pool with a size of 3 for testing
	suite.pool = New[int](3)
}

// TearDownTest releases/destroys the current suite.pool
// object to allow a fresh instance to be initialised
// between each test.
func (suite *PoolTestSuite) TearDownTest() {

	err := suite.pool.Teardown()
	if err != nil {
		panic(err)
	}
}

// Test for the Get function
func (suite *PoolTestSuite) TestGet() {
	// Add an item to the fixedsizedpool
	suite.pool.objectPoolChan <- 42
	suite.pool.avail = 1

	// Retrieve the item using Get
	result := suite.pool.Get()

	// Check expectations
	assert.Equal(suite.T(), 42, result)
	assert.Equal(suite.T(), uint32(0), suite.pool.Avail(), "Available count should be 0 after retrieval")
}

// Test for the Put function
func (suite *PoolTestSuite) TestPut() {
	// Add an object to the fixedsizedpool using Put
	err := suite.pool.Put(99)

	// Check expectations
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), uint32(1), suite.pool.Avail(), "Available count should increment")
	assert.Equal(suite.T(), 99, suite.pool.Get(), "The object should be present in the channel")
}

// Test for Put function when the fixedsizedpool is full
func (suite *PoolTestSuite) TestPutWhenFull() {
	// Add objects to fill the fixedsizedpool
	_ = suite.pool.Put(1)
	_ = suite.pool.Put(2)
	_ = suite.pool.Put(3)

	// Attempt to add another object when the fixedsizedpool is full
	err := suite.pool.Put(4)

	// Check expectations
	assert.Error(suite.T(), err)
	assert.EqualError(suite.T(), err, "fixedsizedpool is full. Cannot add more objects")
	assert.Equal(suite.T(), uint32(3), suite.pool.Avail(), "Available count should not change when full")
}

// Test for the Avail function
func (suite *PoolTestSuite) TestAvail() {
	// Initially, the fixedsizedpool should be empty
	assert.Equal(suite.T(), uint32(0), suite.pool.Avail())

	// Add an object to the fixedsizedpool
	_ = suite.pool.Put(23)
	assert.Equal(suite.T(), uint32(1), suite.pool.Avail())
}

// Run the test suite
func TestPoolTestSuite(t *testing.T) {
	suite.Run(t, new(PoolTestSuite))
}
