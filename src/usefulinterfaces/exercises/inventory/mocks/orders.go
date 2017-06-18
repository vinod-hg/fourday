package mocks

import "usefulinterfaces/exercises/inventory"
import "fmt"
import "sync"

//DomainModel type
type DomainModel map[int]*inventory.Order

//MockOrder funciton
type MockOrder struct {
	mu   sync.Mutex
	omap DomainModel
}

var _ inventory.OrderStorage = (*MockOrder)(nil)

// NewMockOrder function
func NewMockOrder() *MockOrder {
	return &MockOrder{omap: make(map[int]*inventory.Order)}
}

// Get function
func (os *MockOrder) Get(id int) (*inventory.Order, error) {
	defer os.mu.Unlock()
	os.mu.Lock()
	if i, ok := os.omap[id]; ok {
		return i, nil
	}
	return nil, fmt.Errorf("not found")
}

// Create function
func (os *MockOrder) Create(o inventory.Order) (*inventory.Order, error) {
	defer os.mu.Unlock()
	os.mu.Lock()
	os.omap[o.ID] = &o
	return nil, nil
}
