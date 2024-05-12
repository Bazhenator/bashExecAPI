//go:build unit
// +build unit

package service

import (
	"context"
	mocks_repository "github.com/Bazhenator/bashExecAPI/internal/repository/mock"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestServices_DeleteAllRows(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_repository.NewMockIDataBaseRepository(ctrl)
	s := Services{dbRepo: mockRepo}

	mockRepo.EXPECT().DeleteAllRows(gomock.Any()).Return(nil)

	err := s.DeleteAllRows(context.Background())
	assert.NoError(t, err)
}

func TestServices_DeleteRow(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_repository.NewMockIDataBaseRepository(ctrl)
	s := Services{dbRepo: mockRepo}

	mockRepo.EXPECT().DeleteRow(gomock.Any(), gomock.Any()).Return(nil)

	err := s.DeleteRow(context.Background(), 1)
	assert.NoError(t, err)
}
