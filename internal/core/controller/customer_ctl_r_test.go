package controller

// import (
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/tbtec/tremligeiro/internal/dto"
// 	"github.com/tbtec/tremligeiro/internal/infra/container"
// 	"github.com/tbtec/tremligeiro/internal/infra/database/model"
// 	"github.com/tbtec/tremligeiro/test/repository"
// )

// func TestFindCustomerController_Execute_Success_WithMock(t *testing.T) {

// 	mockCustomer := &model.Customer{
// 		ID:             "customer1",
// 		Name:           "Customer 1",
// 		DocumentNumber: "123456789",
// 		Email:          "teste@teste.com",
// 		CreatedAt:      time.Now(),
// 		UpdatedAt:      time.Now(),
// 	}

// 	customerRepo := &repository.MockCustomerRepo{
// 		FindByDocumentNumberFunc: func(ctx context.Context, id string) (*model.Customer, error) {
// 			return mockCustomer, nil
// 		},
// 	}

// 	container := &container.Container{
// 		CustomerRepository: customerRepo,
// 	}

// 	controller := NewFindCustomerController(container)

// 	customerIdRequest := dto.FindCustomer{
// 		DocumentNumber: "123456789",
// 	}

// 	ctx := context.Background()
// 	result, err := controller.Execute(ctx, customerIdRequest)

// 	assert.NoError(t, err)
// 	assert.Equal(t, "customer1", result.Content.CustomerId)
// 	assert.Equal(t, "Customer 1", result.Content.Name)
// }

// func TestFindCustomerController_Execute_NotFound(t *testing.T) {

// 	customerRepo := &repository.MockCustomerRepo{
// 		FindByDocumentNumberFunc: func(ctx context.Context, id string) (*model.Customer, error) {

// 			return nil, errors.New("not found")
// 		},
// 	}

// 	container := &container.Container{
// 		CustomerRepository: customerRepo,
// 	}

// 	customerIdRequest := dto.FindCustomer{
// 		DocumentNumber: "customer-404",
// 	}

// 	controller := NewFindCustomerController(container)

// 	ctx := context.Background()
// 	_, err := controller.Execute(ctx, customerIdRequest)

// 	assert.Error(t, err)
// }
