// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"shopular/ent/cart"
	"shopular/ent/product"
	"shopular/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CartCreate is the builder for creating a Cart entity.
type CartCreate struct {
	config
	mutation *CartMutation
	hooks    []Hook
}

// SetQuantity sets the "quantity" field.
func (cc *CartCreate) SetQuantity(i int) *CartCreate {
	cc.mutation.SetQuantity(i)
	return cc
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (cc *CartCreate) AddUserIDs(ids ...int) *CartCreate {
	cc.mutation.AddUserIDs(ids...)
	return cc
}

// AddUser adds the "user" edges to the User entity.
func (cc *CartCreate) AddUser(u ...*User) *CartCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cc.AddUserIDs(ids...)
}

// AddProductIDs adds the "product" edge to the Product entity by IDs.
func (cc *CartCreate) AddProductIDs(ids ...int) *CartCreate {
	cc.mutation.AddProductIDs(ids...)
	return cc
}

// AddProduct adds the "product" edges to the Product entity.
func (cc *CartCreate) AddProduct(p ...*Product) *CartCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddProductIDs(ids...)
}

// Mutation returns the CartMutation object of the builder.
func (cc *CartCreate) Mutation() *CartMutation {
	return cc.mutation
}

// Save creates the Cart in the database.
func (cc *CartCreate) Save(ctx context.Context) (*Cart, error) {
	var (
		err  error
		node *Cart
	)
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Cart)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from CartMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CartCreate) SaveX(ctx context.Context) *Cart {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CartCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CartCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CartCreate) check() error {
	if _, ok := cc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "Cart.quantity"`)}
	}
	if len(cc.mutation.ProductIDs()) == 0 {
		return &ValidationError{Name: "product", err: errors.New(`ent: missing required edge "Cart.product"`)}
	}
	return nil
}

func (cc *CartCreate) sqlSave(ctx context.Context) (*Cart, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (cc *CartCreate) createSpec() (*Cart, *sqlgraph.CreateSpec) {
	var (
		_node = &Cart{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: cart.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: cart.FieldID,
			},
		}
	)
	if value, ok := cc.mutation.Quantity(); ok {
		_spec.SetField(cart.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if nodes := cc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cart.UserTable,
			Columns: cart.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   cart.ProductTable,
			Columns: cart.ProductPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CartCreateBulk is the builder for creating many Cart entities in bulk.
type CartCreateBulk struct {
	config
	builders []*CartCreate
}

// Save creates the Cart entities in the database.
func (ccb *CartCreateBulk) Save(ctx context.Context) ([]*Cart, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cart, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CartMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CartCreateBulk) SaveX(ctx context.Context) []*Cart {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CartCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CartCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
