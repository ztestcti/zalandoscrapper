package aws

import "fmt"

type MockService struct{}

func NewMockService() *MockService {
	return &MockService{}
}

func (m *MockService) ListBuckets() []string {
	fmt.Println("Fetching buckets from AWS...")
	return []string{"bucket1", "bucket2", "bucket3"}
}
