package health

import (
	"context"
	"fiber-onion-boiler-plate/internal/entity"
	"fiber-onion-boiler-plate/internal/health/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCheck(t *testing.T) {
	tests := []struct {
		name                  string
		mockDatabaseError     error
		expectedDatabaseState entity.HealthState
		expectedRedisState    entity.HealthState
		expectedHealthy       bool
	}{
		{
			name:                  "database healthy",
			mockDatabaseError:     nil,
			expectedDatabaseState: entity.HealthStateOK,
			expectedRedisState:    entity.HealthStateOK,
			expectedHealthy:       true,
		},
		{
			name:                  "database unhealthy",
			mockDatabaseError:     assert.AnError,
			expectedDatabaseState: entity.HealthStateFail,
			expectedRedisState:    entity.HealthStateOK,
			expectedHealthy:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := mocks.NewRepository(t)
			mockRepo.On("CheckDatabase", mock.Anything).Return(tt.mockDatabaseError).Once()

			svc := &Service{repo: mockRepo}
			healthComponent, isHealthy := svc.Check(context.Background())

			assert.Equal(t, tt.expectedDatabaseState, healthComponent.Database)
			assert.Equal(t, tt.expectedHealthy, isHealthy)

			mockRepo.AssertExpectations(t)
		})
	}
}
