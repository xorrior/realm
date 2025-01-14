package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.26

import (
	"context"
	"fmt"

	"github.com/kcarretto/realm/tavern/internal/ent"
)

// Files is the resolver for the files field.
func (r *queryResolver) Files(ctx context.Context, where *ent.FileWhereInput) ([]*ent.File, error) {
	if where != nil {
		query, err := where.Filter(r.client.File.Query())
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return r.client.File.Query().All(ctx)
}

// Quests is the resolver for the quests field.
func (r *queryResolver) Quests(ctx context.Context, where *ent.QuestWhereInput) ([]*ent.Quest, error) {
	if where != nil {
		query, err := where.Filter(r.client.Quest.Query())
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return r.client.Quest.Query().All(ctx)
}

// Beacons is the resolver for the beacons field.
func (r *queryResolver) Beacons(ctx context.Context, where *ent.BeaconWhereInput) ([]*ent.Beacon, error) {
	if where != nil {
		query, err := where.Filter(r.client.Beacon.Query())
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return r.client.Beacon.Query().All(ctx)
}

// Tags is the resolver for the tags field.
func (r *queryResolver) Tags(ctx context.Context, where *ent.TagWhereInput) ([]*ent.Tag, error) {
	if where != nil {
		query, err := where.Filter(r.client.Tag.Query())
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return r.client.Tag.Query().All(ctx)
}

// Tomes is the resolver for the tomes field.
func (r *queryResolver) Tomes(ctx context.Context, where *ent.TomeWhereInput) ([]*ent.Tome, error) {
	if where != nil {
		query, err := where.Filter(r.client.Tome.Query())
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return r.client.Tome.Query().All(ctx)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, where *ent.UserWhereInput) ([]*ent.User, error) {
	if where != nil {
		query, err := where.Filter(r.client.User.Query())
		if err != nil {
			return nil, fmt.Errorf("failed to apply filter: %w", err)
		}
		return query.All(ctx)
	}
	return r.client.User.Query().All(ctx)
}
