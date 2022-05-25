// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/event"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"
)

// TransactionCreate is the builder for creating a Transaction entity.
type TransactionCreate struct {
	config
	mutation *TransactionMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetContractAddress sets the "contract_address" field.
func (tc *TransactionCreate) SetContractAddress(s string) *TransactionCreate {
	tc.mutation.SetContractAddress(s)
	return tc
}

// SetEntryPointSelector sets the "entry_point_selector" field.
func (tc *TransactionCreate) SetEntryPointSelector(s string) *TransactionCreate {
	tc.mutation.SetEntryPointSelector(s)
	return tc
}

// SetNillableEntryPointSelector sets the "entry_point_selector" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableEntryPointSelector(s *string) *TransactionCreate {
	if s != nil {
		tc.SetEntryPointSelector(*s)
	}
	return tc
}

// SetTransactionHash sets the "transaction_hash" field.
func (tc *TransactionCreate) SetTransactionHash(s string) *TransactionCreate {
	tc.mutation.SetTransactionHash(s)
	return tc
}

// SetCalldata sets the "calldata" field.
func (tc *TransactionCreate) SetCalldata(s []string) *TransactionCreate {
	tc.mutation.SetCalldata(s)
	return tc
}

// SetSignature sets the "signature" field.
func (tc *TransactionCreate) SetSignature(s []string) *TransactionCreate {
	tc.mutation.SetSignature(s)
	return tc
}

// SetNonce sets the "nonce" field.
func (tc *TransactionCreate) SetNonce(s string) *TransactionCreate {
	tc.mutation.SetNonce(s)
	return tc
}

// SetNillableNonce sets the "nonce" field if the given value is not nil.
func (tc *TransactionCreate) SetNillableNonce(s *string) *TransactionCreate {
	if s != nil {
		tc.SetNonce(*s)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TransactionCreate) SetID(s string) *TransactionCreate {
	tc.mutation.SetID(s)
	return tc
}

// SetBlockID sets the "block" edge to the Block entity by ID.
func (tc *TransactionCreate) SetBlockID(id string) *TransactionCreate {
	tc.mutation.SetBlockID(id)
	return tc
}

// SetNillableBlockID sets the "block" edge to the Block entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableBlockID(id *string) *TransactionCreate {
	if id != nil {
		tc = tc.SetBlockID(*id)
	}
	return tc
}

// SetBlock sets the "block" edge to the Block entity.
func (tc *TransactionCreate) SetBlock(b *Block) *TransactionCreate {
	return tc.SetBlockID(b.ID)
}

// SetReceiptID sets the "receipt" edge to the TransactionReceipt entity by ID.
func (tc *TransactionCreate) SetReceiptID(id string) *TransactionCreate {
	tc.mutation.SetReceiptID(id)
	return tc
}

// SetNillableReceiptID sets the "receipt" edge to the TransactionReceipt entity by ID if the given value is not nil.
func (tc *TransactionCreate) SetNillableReceiptID(id *string) *TransactionCreate {
	if id != nil {
		tc = tc.SetReceiptID(*id)
	}
	return tc
}

// SetReceipt sets the "receipt" edge to the TransactionReceipt entity.
func (tc *TransactionCreate) SetReceipt(t *TransactionReceipt) *TransactionCreate {
	return tc.SetReceiptID(t.ID)
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (tc *TransactionCreate) AddEventIDs(ids ...string) *TransactionCreate {
	tc.mutation.AddEventIDs(ids...)
	return tc
}

// AddEvents adds the "events" edges to the Event entity.
func (tc *TransactionCreate) AddEvents(e ...*Event) *TransactionCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tc.AddEventIDs(ids...)
}

// Mutation returns the TransactionMutation object of the builder.
func (tc *TransactionCreate) Mutation() *TransactionMutation {
	return tc.mutation
}

// Save creates the Transaction in the database.
func (tc *TransactionCreate) Save(ctx context.Context) (*Transaction, error) {
	var (
		err  error
		node *Transaction
	)
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TransactionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TransactionCreate) SaveX(ctx context.Context) *Transaction {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TransactionCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TransactionCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TransactionCreate) check() error {
	if _, ok := tc.mutation.ContractAddress(); !ok {
		return &ValidationError{Name: "contract_address", err: errors.New(`ent: missing required field "Transaction.contract_address"`)}
	}
	if _, ok := tc.mutation.TransactionHash(); !ok {
		return &ValidationError{Name: "transaction_hash", err: errors.New(`ent: missing required field "Transaction.transaction_hash"`)}
	}
	if _, ok := tc.mutation.Calldata(); !ok {
		return &ValidationError{Name: "calldata", err: errors.New(`ent: missing required field "Transaction.calldata"`)}
	}
	return nil
}

func (tc *TransactionCreate) sqlSave(ctx context.Context) (*Transaction, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Transaction.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (tc *TransactionCreate) createSpec() (*Transaction, *sqlgraph.CreateSpec) {
	var (
		_node = &Transaction{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: transaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: transaction.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.ContractAddress(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldContractAddress,
		})
		_node.ContractAddress = value
	}
	if value, ok := tc.mutation.EntryPointSelector(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldEntryPointSelector,
		})
		_node.EntryPointSelector = value
	}
	if value, ok := tc.mutation.TransactionHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldTransactionHash,
		})
		_node.TransactionHash = value
	}
	if value, ok := tc.mutation.Calldata(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transaction.FieldCalldata,
		})
		_node.Calldata = value
	}
	if value, ok := tc.mutation.Signature(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: transaction.FieldSignature,
		})
		_node.Signature = value
	}
	if value, ok := tc.mutation.Nonce(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: transaction.FieldNonce,
		})
		_node.Nonce = value
	}
	if nodes := tc.mutation.BlockIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   transaction.BlockTable,
			Columns: []string{transaction.BlockColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: block.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.block_transactions = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ReceiptIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   transaction.ReceiptTable,
			Columns: []string{transaction.ReceiptColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transactionreceipt.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   transaction.EventsTable,
			Columns: []string{transaction.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: event.FieldID,
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

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Transaction.Create().
//		SetContractAddress(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TransactionUpsert) {
//			SetContractAddress(v+v).
//		}).
//		Exec(ctx)
//
func (tc *TransactionCreate) OnConflict(opts ...sql.ConflictOption) *TransactionUpsertOne {
	tc.conflict = opts
	return &TransactionUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tc *TransactionCreate) OnConflictColumns(columns ...string) *TransactionUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TransactionUpsertOne{
		create: tc,
	}
}

type (
	// TransactionUpsertOne is the builder for "upsert"-ing
	//  one Transaction node.
	TransactionUpsertOne struct {
		create *TransactionCreate
	}

	// TransactionUpsert is the "OnConflict" setter.
	TransactionUpsert struct {
		*sql.UpdateSet
	}
)

// SetContractAddress sets the "contract_address" field.
func (u *TransactionUpsert) SetContractAddress(v string) *TransactionUpsert {
	u.Set(transaction.FieldContractAddress, v)
	return u
}

// UpdateContractAddress sets the "contract_address" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateContractAddress() *TransactionUpsert {
	u.SetExcluded(transaction.FieldContractAddress)
	return u
}

// SetEntryPointSelector sets the "entry_point_selector" field.
func (u *TransactionUpsert) SetEntryPointSelector(v string) *TransactionUpsert {
	u.Set(transaction.FieldEntryPointSelector, v)
	return u
}

// UpdateEntryPointSelector sets the "entry_point_selector" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateEntryPointSelector() *TransactionUpsert {
	u.SetExcluded(transaction.FieldEntryPointSelector)
	return u
}

// ClearEntryPointSelector clears the value of the "entry_point_selector" field.
func (u *TransactionUpsert) ClearEntryPointSelector() *TransactionUpsert {
	u.SetNull(transaction.FieldEntryPointSelector)
	return u
}

// SetTransactionHash sets the "transaction_hash" field.
func (u *TransactionUpsert) SetTransactionHash(v string) *TransactionUpsert {
	u.Set(transaction.FieldTransactionHash, v)
	return u
}

// UpdateTransactionHash sets the "transaction_hash" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateTransactionHash() *TransactionUpsert {
	u.SetExcluded(transaction.FieldTransactionHash)
	return u
}

// SetCalldata sets the "calldata" field.
func (u *TransactionUpsert) SetCalldata(v []string) *TransactionUpsert {
	u.Set(transaction.FieldCalldata, v)
	return u
}

// UpdateCalldata sets the "calldata" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateCalldata() *TransactionUpsert {
	u.SetExcluded(transaction.FieldCalldata)
	return u
}

// SetSignature sets the "signature" field.
func (u *TransactionUpsert) SetSignature(v []string) *TransactionUpsert {
	u.Set(transaction.FieldSignature, v)
	return u
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateSignature() *TransactionUpsert {
	u.SetExcluded(transaction.FieldSignature)
	return u
}

// ClearSignature clears the value of the "signature" field.
func (u *TransactionUpsert) ClearSignature() *TransactionUpsert {
	u.SetNull(transaction.FieldSignature)
	return u
}

// SetNonce sets the "nonce" field.
func (u *TransactionUpsert) SetNonce(v string) *TransactionUpsert {
	u.Set(transaction.FieldNonce, v)
	return u
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *TransactionUpsert) UpdateNonce() *TransactionUpsert {
	u.SetExcluded(transaction.FieldNonce)
	return u
}

// ClearNonce clears the value of the "nonce" field.
func (u *TransactionUpsert) ClearNonce() *TransactionUpsert {
	u.SetNull(transaction.FieldNonce)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transaction.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TransactionUpsertOne) UpdateNewValues() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(transaction.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Transaction.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *TransactionUpsertOne) Ignore() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TransactionUpsertOne) DoNothing() *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TransactionCreate.OnConflict
// documentation for more info.
func (u *TransactionUpsertOne) Update(set func(*TransactionUpsert)) *TransactionUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TransactionUpsert{UpdateSet: update})
	}))
	return u
}

// SetContractAddress sets the "contract_address" field.
func (u *TransactionUpsertOne) SetContractAddress(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetContractAddress(v)
	})
}

// UpdateContractAddress sets the "contract_address" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateContractAddress() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateContractAddress()
	})
}

// SetEntryPointSelector sets the "entry_point_selector" field.
func (u *TransactionUpsertOne) SetEntryPointSelector(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetEntryPointSelector(v)
	})
}

// UpdateEntryPointSelector sets the "entry_point_selector" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateEntryPointSelector() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateEntryPointSelector()
	})
}

// ClearEntryPointSelector clears the value of the "entry_point_selector" field.
func (u *TransactionUpsertOne) ClearEntryPointSelector() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearEntryPointSelector()
	})
}

// SetTransactionHash sets the "transaction_hash" field.
func (u *TransactionUpsertOne) SetTransactionHash(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTransactionHash(v)
	})
}

// UpdateTransactionHash sets the "transaction_hash" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateTransactionHash() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTransactionHash()
	})
}

// SetCalldata sets the "calldata" field.
func (u *TransactionUpsertOne) SetCalldata(v []string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCalldata(v)
	})
}

// UpdateCalldata sets the "calldata" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateCalldata() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCalldata()
	})
}

// SetSignature sets the "signature" field.
func (u *TransactionUpsertOne) SetSignature(v []string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateSignature() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateSignature()
	})
}

// ClearSignature clears the value of the "signature" field.
func (u *TransactionUpsertOne) ClearSignature() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearSignature()
	})
}

// SetNonce sets the "nonce" field.
func (u *TransactionUpsertOne) SetNonce(v string) *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.SetNonce(v)
	})
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *TransactionUpsertOne) UpdateNonce() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateNonce()
	})
}

// ClearNonce clears the value of the "nonce" field.
func (u *TransactionUpsertOne) ClearNonce() *TransactionUpsertOne {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearNonce()
	})
}

// Exec executes the query.
func (u *TransactionUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TransactionCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TransactionUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TransactionUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TransactionUpsertOne.ID is not supported by MySQL driver. Use TransactionUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TransactionUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TransactionCreateBulk is the builder for creating many Transaction entities in bulk.
type TransactionCreateBulk struct {
	config
	builders []*TransactionCreate
	conflict []sql.ConflictOption
}

// Save creates the Transaction entities in the database.
func (tcb *TransactionCreateBulk) Save(ctx context.Context) ([]*Transaction, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Transaction, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TransactionMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TransactionCreateBulk) SaveX(ctx context.Context) []*Transaction {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TransactionCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TransactionCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Transaction.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TransactionUpsert) {
//			SetContractAddress(v+v).
//		}).
//		Exec(ctx)
//
func (tcb *TransactionCreateBulk) OnConflict(opts ...sql.ConflictOption) *TransactionUpsertBulk {
	tcb.conflict = opts
	return &TransactionUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (tcb *TransactionCreateBulk) OnConflictColumns(columns ...string) *TransactionUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TransactionUpsertBulk{
		create: tcb,
	}
}

// TransactionUpsertBulk is the builder for "upsert"-ing
// a bulk of Transaction nodes.
type TransactionUpsertBulk struct {
	create *TransactionCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(transaction.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *TransactionUpsertBulk) UpdateNewValues() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(transaction.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Transaction.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *TransactionUpsertBulk) Ignore() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TransactionUpsertBulk) DoNothing() *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TransactionCreateBulk.OnConflict
// documentation for more info.
func (u *TransactionUpsertBulk) Update(set func(*TransactionUpsert)) *TransactionUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TransactionUpsert{UpdateSet: update})
	}))
	return u
}

// SetContractAddress sets the "contract_address" field.
func (u *TransactionUpsertBulk) SetContractAddress(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetContractAddress(v)
	})
}

// UpdateContractAddress sets the "contract_address" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateContractAddress() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateContractAddress()
	})
}

// SetEntryPointSelector sets the "entry_point_selector" field.
func (u *TransactionUpsertBulk) SetEntryPointSelector(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetEntryPointSelector(v)
	})
}

// UpdateEntryPointSelector sets the "entry_point_selector" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateEntryPointSelector() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateEntryPointSelector()
	})
}

// ClearEntryPointSelector clears the value of the "entry_point_selector" field.
func (u *TransactionUpsertBulk) ClearEntryPointSelector() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearEntryPointSelector()
	})
}

// SetTransactionHash sets the "transaction_hash" field.
func (u *TransactionUpsertBulk) SetTransactionHash(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetTransactionHash(v)
	})
}

// UpdateTransactionHash sets the "transaction_hash" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateTransactionHash() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateTransactionHash()
	})
}

// SetCalldata sets the "calldata" field.
func (u *TransactionUpsertBulk) SetCalldata(v []string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetCalldata(v)
	})
}

// UpdateCalldata sets the "calldata" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateCalldata() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateCalldata()
	})
}

// SetSignature sets the "signature" field.
func (u *TransactionUpsertBulk) SetSignature(v []string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetSignature(v)
	})
}

// UpdateSignature sets the "signature" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateSignature() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateSignature()
	})
}

// ClearSignature clears the value of the "signature" field.
func (u *TransactionUpsertBulk) ClearSignature() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearSignature()
	})
}

// SetNonce sets the "nonce" field.
func (u *TransactionUpsertBulk) SetNonce(v string) *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.SetNonce(v)
	})
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *TransactionUpsertBulk) UpdateNonce() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.UpdateNonce()
	})
}

// ClearNonce clears the value of the "nonce" field.
func (u *TransactionUpsertBulk) ClearNonce() *TransactionUpsertBulk {
	return u.Update(func(s *TransactionUpsert) {
		s.ClearNonce()
	})
}

// Exec executes the query.
func (u *TransactionUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TransactionCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TransactionCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TransactionUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
