// Code generated by entc, DO NOT EDIT.

package ent

import (
	"IMDB/ent/movie"
	"IMDB/ent/predicate"
	"IMDB/ent/review"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetText sets the "text" field.
func (ru *ReviewUpdate) SetText(s string) *ReviewUpdate {
	ru.mutation.SetText(s)
	return ru
}

// SetRank sets the "rank" field.
func (ru *ReviewUpdate) SetRank(f float32) *ReviewUpdate {
	ru.mutation.ResetRank()
	ru.mutation.SetRank(f)
	return ru
}

// AddRank adds f to the "rank" field.
func (ru *ReviewUpdate) AddRank(f float32) *ReviewUpdate {
	ru.mutation.AddRank(f)
	return ru
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (ru *ReviewUpdate) AddMovieIDs(ids ...int) *ReviewUpdate {
	ru.mutation.AddMovieIDs(ids...)
	return ru
}

// AddMovies adds the "movies" edges to the Movie entity.
func (ru *ReviewUpdate) AddMovies(m ...*Movie) *ReviewUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ru.AddMovieIDs(ids...)
}

// AddUserIDs adds the "user" edge to the Review entity by IDs.
func (ru *ReviewUpdate) AddUserIDs(ids ...int) *ReviewUpdate {
	ru.mutation.AddUserIDs(ids...)
	return ru
}

// AddUser adds the "user" edges to the Review entity.
func (ru *ReviewUpdate) AddUser(r ...*Review) *ReviewUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.AddUserIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (ru *ReviewUpdate) ClearMovies() *ReviewUpdate {
	ru.mutation.ClearMovies()
	return ru
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (ru *ReviewUpdate) RemoveMovieIDs(ids ...int) *ReviewUpdate {
	ru.mutation.RemoveMovieIDs(ids...)
	return ru
}

// RemoveMovies removes "movies" edges to Movie entities.
func (ru *ReviewUpdate) RemoveMovies(m ...*Movie) *ReviewUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ru.RemoveMovieIDs(ids...)
}

// ClearUser clears all "user" edges to the Review entity.
func (ru *ReviewUpdate) ClearUser() *ReviewUpdate {
	ru.mutation.ClearUser()
	return ru
}

// RemoveUserIDs removes the "user" edge to Review entities by IDs.
func (ru *ReviewUpdate) RemoveUserIDs(ids ...int) *ReviewUpdate {
	ru.mutation.RemoveUserIDs(ids...)
	return ru
}

// RemoveUser removes "user" edges to Review entities.
func (ru *ReviewUpdate) RemoveUser(r ...*Review) *ReviewUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ru.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldText,
		})
	}
	if value, ok := ru.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: review.FieldRank,
		})
	}
	if value, ok := ru.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: review.FieldRank,
		})
	}
	if ru.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.MoviesTable,
			Columns: review.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !ru.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.MoviesTable,
			Columns: review.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.MoviesTable,
			Columns: review.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.UserTable,
			Columns: review.UserPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedUserIDs(); len(nodes) > 0 && !ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.UserTable,
			Columns: review.UserPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.UserTable,
			Columns: review.UserPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetText sets the "text" field.
func (ruo *ReviewUpdateOne) SetText(s string) *ReviewUpdateOne {
	ruo.mutation.SetText(s)
	return ruo
}

// SetRank sets the "rank" field.
func (ruo *ReviewUpdateOne) SetRank(f float32) *ReviewUpdateOne {
	ruo.mutation.ResetRank()
	ruo.mutation.SetRank(f)
	return ruo
}

// AddRank adds f to the "rank" field.
func (ruo *ReviewUpdateOne) AddRank(f float32) *ReviewUpdateOne {
	ruo.mutation.AddRank(f)
	return ruo
}

// AddMovieIDs adds the "movies" edge to the Movie entity by IDs.
func (ruo *ReviewUpdateOne) AddMovieIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.AddMovieIDs(ids...)
	return ruo
}

// AddMovies adds the "movies" edges to the Movie entity.
func (ruo *ReviewUpdateOne) AddMovies(m ...*Movie) *ReviewUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ruo.AddMovieIDs(ids...)
}

// AddUserIDs adds the "user" edge to the Review entity by IDs.
func (ruo *ReviewUpdateOne) AddUserIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.AddUserIDs(ids...)
	return ruo
}

// AddUser adds the "user" edges to the Review entity.
func (ruo *ReviewUpdateOne) AddUser(r ...*Review) *ReviewUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.AddUserIDs(ids...)
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// ClearMovies clears all "movies" edges to the Movie entity.
func (ruo *ReviewUpdateOne) ClearMovies() *ReviewUpdateOne {
	ruo.mutation.ClearMovies()
	return ruo
}

// RemoveMovieIDs removes the "movies" edge to Movie entities by IDs.
func (ruo *ReviewUpdateOne) RemoveMovieIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.RemoveMovieIDs(ids...)
	return ruo
}

// RemoveMovies removes "movies" edges to Movie entities.
func (ruo *ReviewUpdateOne) RemoveMovies(m ...*Movie) *ReviewUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return ruo.RemoveMovieIDs(ids...)
}

// ClearUser clears all "user" edges to the Review entity.
func (ruo *ReviewUpdateOne) ClearUser() *ReviewUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// RemoveUserIDs removes the "user" edge to Review entities by IDs.
func (ruo *ReviewUpdateOne) RemoveUserIDs(ids ...int) *ReviewUpdateOne {
	ruo.mutation.RemoveUserIDs(ids...)
	return ruo
}

// RemoveUser removes "user" edges to Review entities.
func (ruo *ReviewUpdateOne) RemoveUser(r ...*Review) *ReviewUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return ruo.RemoveUserIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: review.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Review.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldText,
		})
	}
	if value, ok := ruo.mutation.Rank(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: review.FieldRank,
		})
	}
	if value, ok := ruo.mutation.AddedRank(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat32,
			Value:  value,
			Column: review.FieldRank,
		})
	}
	if ruo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.MoviesTable,
			Columns: review.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedMoviesIDs(); len(nodes) > 0 && !ruo.mutation.MoviesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.MoviesTable,
			Columns: review.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.MoviesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.MoviesTable,
			Columns: review.MoviesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: movie.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.UserTable,
			Columns: review.UserPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedUserIDs(); len(nodes) > 0 && !ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.UserTable,
			Columns: review.UserPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   review.UserTable,
			Columns: review.UserPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: review.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
