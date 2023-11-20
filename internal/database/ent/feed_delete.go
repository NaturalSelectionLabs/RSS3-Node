// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/naturalselectionlabs/rss3-node/internal/database/ent/feed"
	"github.com/naturalselectionlabs/rss3-node/internal/database/ent/predicate"
)

// FeedDelete is the builder for deleting a Feed entity.
type FeedDelete struct {
	config
	hooks    []Hook
	mutation *FeedMutation
}

// Where appends a list predicates to the FeedDelete builder.
func (fd *FeedDelete) Where(ps ...predicate.Feed) *FeedDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FeedDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FeedDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FeedDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(feed.Table, sqlgraph.NewFieldSpec(feed.FieldID, field.TypeString))
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FeedDeleteOne is the builder for deleting a single Feed entity.
type FeedDeleteOne struct {
	fd *FeedDelete
}

// Where appends a list predicates to the FeedDelete builder.
func (fdo *FeedDeleteOne) Where(ps ...predicate.Feed) *FeedDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FeedDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{feed.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FeedDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
