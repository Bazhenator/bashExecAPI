package service

import "context"

func (s *Services) DeleteAllRows(ctx context.Context) error {
	return s.dbRepo.DeleteAllRows(ctx)
}

func (s *Services) DeleteRow(ctx context.Context, id int) error {
	return s.dbRepo.DeleteRow(ctx, id)
}
