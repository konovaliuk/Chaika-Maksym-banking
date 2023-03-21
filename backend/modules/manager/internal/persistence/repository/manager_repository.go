package repository

import (
	"context"

	manager_events "github.com/fabl3ss/banking_system/api/manager/events"
	manager_model "github.com/fabl3ss/banking_system/modules/manager/internal/model"
	"github.com/fabl3ss/banking_system/pkg/producer"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ManagerRepository struct {
	eventProducer *producer.Producer
}

func NewManagerRepository(eventProducer *producer.Producer) *ManagerRepository {
	return &ManagerRepository{
		eventProducer: eventProducer,
	}
}

func (r *ManagerRepository) Create(ctx context.Context, manager *manager_model.Manager) error {
	event := &manager_events.ManagerEvent{
		ManagerId: manager.Id.String(),
		Event: &manager_events.ManagerEvent_ManagerCreated{
			ManagerCreated: &manager_events.ManagerCreated{
				FullName:  manager.FullName,
				Role:      mapManagerRoleToDto(manager.Role),
				CreatedAt: timestamppb.New(manager.JoinedAt),
			},
		},
	}

	return r.eventProducer.Publish(ctx, event, "manager-events", manager.Id.String())
}

func mapManagerRoleToDto(role manager_model.ManagerRole) manager_events.ManagerRole {
	switch role {
	case manager_model.ManagerRoleAdmin:
		return manager_events.ManagerRole_MANAGER_ROLE_ADMIN
	case manager_model.ManagerRoleModerator:
		return manager_events.ManagerRole_MANAGER_ROLE_MODERATOR
	}

	return manager_events.ManagerRole_MANAGER_ROLE_UNDEFINED
}
