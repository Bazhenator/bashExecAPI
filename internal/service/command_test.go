package service

import (
	"context"
	"github.com/Bazhenator/bashExecAPI/internal/domain"

	mocks_repository "github.com/Bazhenator/bashExecAPI/internal/repository/mock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestServices_CreateCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_repository.NewMockICommandRepository(ctrl)
	s := Services{commandRepo: mockRepo}

	mockRepo.EXPECT().CreateCommand(gomock.Any(), gomock.Any()).Return("result", "id", nil)

	result, _, err := s.CreateCommand(context.Background(), "command")
	assert.NoError(t, err)
	assert.Equal(t, "result", result)
}

func TestServices_ListCommands(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_repository.NewMockICommandRepository(ctrl)
	s := Services{commandRepo: mockRepo}

	out1 := "result1"
	out2 := "result2"
	commands := []domain.Command{
		{ID: 1, Command: "command1", Result: &out1},
		{ID: 2, Command: "command2", Result: &out2},
	}

	mockRepo.EXPECT().ListCommands(gomock.Any()).Return(commands, nil)

	result, err := s.ListCommands(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, commands, result)
}

func TestServices_GetCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_repository.NewMockICommandRepository(ctrl)
	s := Services{commandRepo: mockRepo}

	out := "result"
	command := &domain.Command{ID: 1, Command: "command", Result: &out}

	mockRepo.EXPECT().GetCommand(gomock.Any(), gomock.Any()).Return(command, nil)

	result, err := s.GetCommand(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, command, result)
}

func TestServices_RunCommand(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks_repository.NewMockICommandRepository(ctrl)
	s := Services{commandRepo: mockRepo}

	mockRepo.EXPECT().RunCommand(gomock.Any(), gomock.Any()).Return("result", nil)

	result, err := s.RunCommand(context.Background(), 1)
	assert.NoError(t, err)
	assert.Equal(t, "result", result)
}
