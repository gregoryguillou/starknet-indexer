// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/cartridge-gg/starknet-indexer/ent/block"
	"github.com/cartridge-gg/starknet-indexer/ent/transaction"
	"github.com/cartridge-gg/starknet-indexer/ent/transactionreceipt"
)

// BlockCreate is the builder for creating a Block entity.
type BlockCreate struct {
	config
	mutation *BlockMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetBlockHash sets the "block_hash" field.
func (bc *BlockCreate) SetBlockHash(s string) *BlockCreate {
	bc.mutation.SetBlockHash(s)
	return bc
}

// SetParentBlockHash sets the "parent_block_hash" field.
func (bc *BlockCreate) SetParentBlockHash(s string) *BlockCreate {
	bc.mutation.SetParentBlockHash(s)
	return bc
}

// SetBlockNumber sets the "block_number" field.
func (bc *BlockCreate) SetBlockNumber(u uint64) *BlockCreate {
	bc.mutation.SetBlockNumber(u)
	return bc
}

// SetStateRoot sets the "state_root" field.
func (bc *BlockCreate) SetStateRoot(s string) *BlockCreate {
	bc.mutation.SetStateRoot(s)
	return bc
}

// SetStatus sets the "status" field.
func (bc *BlockCreate) SetStatus(b block.Status) *BlockCreate {
	bc.mutation.SetStatus(b)
	return bc
}

// SetTimestamp sets the "timestamp" field.
func (bc *BlockCreate) SetTimestamp(t time.Time) *BlockCreate {
	bc.mutation.SetTimestamp(t)
	return bc
}

// SetID sets the "id" field.
func (bc *BlockCreate) SetID(s string) *BlockCreate {
	bc.mutation.SetID(s)
	return bc
}

// AddTransactionIDs adds the "transactions" edge to the Transaction entity by IDs.
func (bc *BlockCreate) AddTransactionIDs(ids ...string) *BlockCreate {
	bc.mutation.AddTransactionIDs(ids...)
	return bc
}

// AddTransactions adds the "transactions" edges to the Transaction entity.
func (bc *BlockCreate) AddTransactions(t ...*Transaction) *BlockCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTransactionIDs(ids...)
}

// AddTransactionReceiptIDs adds the "transaction_receipts" edge to the TransactionReceipt entity by IDs.
func (bc *BlockCreate) AddTransactionReceiptIDs(ids ...string) *BlockCreate {
	bc.mutation.AddTransactionReceiptIDs(ids...)
	return bc
}

// AddTransactionReceipts adds the "transaction_receipts" edges to the TransactionReceipt entity.
func (bc *BlockCreate) AddTransactionReceipts(t ...*TransactionReceipt) *BlockCreate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bc.AddTransactionReceiptIDs(ids...)
}

// Mutation returns the BlockMutation object of the builder.
func (bc *BlockCreate) Mutation() *BlockMutation {
	return bc.mutation
}

// Save creates the Block in the database.
func (bc *BlockCreate) Save(ctx context.Context) (*Block, error) {
	var (
		err  error
		node *Block
	)
	if len(bc.hooks) == 0 {
		if err = bc.check(); err != nil {
			return nil, err
		}
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BlockMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = bc.check(); err != nil {
				return nil, err
			}
			bc.mutation = mutation
			if node, err = bc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			if bc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlockCreate) SaveX(ctx context.Context) *Block {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BlockCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BlockCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BlockCreate) check() error {
	if _, ok := bc.mutation.BlockHash(); !ok {
		return &ValidationError{Name: "block_hash", err: errors.New(`ent: missing required field "Block.block_hash"`)}
	}
	if _, ok := bc.mutation.ParentBlockHash(); !ok {
		return &ValidationError{Name: "parent_block_hash", err: errors.New(`ent: missing required field "Block.parent_block_hash"`)}
	}
	if _, ok := bc.mutation.BlockNumber(); !ok {
		return &ValidationError{Name: "block_number", err: errors.New(`ent: missing required field "Block.block_number"`)}
	}
	if _, ok := bc.mutation.StateRoot(); !ok {
		return &ValidationError{Name: "state_root", err: errors.New(`ent: missing required field "Block.state_root"`)}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Block.status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := block.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Block.status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "Block.timestamp"`)}
	}
	return nil
}

func (bc *BlockCreate) sqlSave(ctx context.Context) (*Block, error) {
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Block.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (bc *BlockCreate) createSpec() (*Block, *sqlgraph.CreateSpec) {
	var (
		_node = &Block{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: block.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: block.FieldID,
			},
		}
	)
	_spec.OnConflict = bc.conflict
	if id, ok := bc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := bc.mutation.BlockHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldBlockHash,
		})
		_node.BlockHash = value
	}
	if value, ok := bc.mutation.ParentBlockHash(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldParentBlockHash,
		})
		_node.ParentBlockHash = value
	}
	if value, ok := bc.mutation.BlockNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint64,
			Value:  value,
			Column: block.FieldBlockNumber,
		})
		_node.BlockNumber = value
	}
	if value, ok := bc.mutation.StateRoot(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: block.FieldStateRoot,
		})
		_node.StateRoot = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: block.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := bc.mutation.Timestamp(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: block.FieldTimestamp,
		})
		_node.Timestamp = value
	}
	if nodes := bc.mutation.TransactionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionsTable,
			Columns: []string{block.TransactionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: transaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.TransactionReceiptsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   block.TransactionReceiptsTable,
			Columns: []string{block.TransactionReceiptsColumn},
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
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Block.Create().
//		SetBlockHash(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlockUpsert) {
//			SetBlockHash(v+v).
//		}).
//		Exec(ctx)
//
func (bc *BlockCreate) OnConflict(opts ...sql.ConflictOption) *BlockUpsertOne {
	bc.conflict = opts
	return &BlockUpsertOne{
		create: bc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Block.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (bc *BlockCreate) OnConflictColumns(columns ...string) *BlockUpsertOne {
	bc.conflict = append(bc.conflict, sql.ConflictColumns(columns...))
	return &BlockUpsertOne{
		create: bc,
	}
}

type (
	// BlockUpsertOne is the builder for "upsert"-ing
	//  one Block node.
	BlockUpsertOne struct {
		create *BlockCreate
	}

	// BlockUpsert is the "OnConflict" setter.
	BlockUpsert struct {
		*sql.UpdateSet
	}
)

// SetBlockHash sets the "block_hash" field.
func (u *BlockUpsert) SetBlockHash(v string) *BlockUpsert {
	u.Set(block.FieldBlockHash, v)
	return u
}

// UpdateBlockHash sets the "block_hash" field to the value that was provided on create.
func (u *BlockUpsert) UpdateBlockHash() *BlockUpsert {
	u.SetExcluded(block.FieldBlockHash)
	return u
}

// SetParentBlockHash sets the "parent_block_hash" field.
func (u *BlockUpsert) SetParentBlockHash(v string) *BlockUpsert {
	u.Set(block.FieldParentBlockHash, v)
	return u
}

// UpdateParentBlockHash sets the "parent_block_hash" field to the value that was provided on create.
func (u *BlockUpsert) UpdateParentBlockHash() *BlockUpsert {
	u.SetExcluded(block.FieldParentBlockHash)
	return u
}

// SetBlockNumber sets the "block_number" field.
func (u *BlockUpsert) SetBlockNumber(v uint64) *BlockUpsert {
	u.Set(block.FieldBlockNumber, v)
	return u
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *BlockUpsert) UpdateBlockNumber() *BlockUpsert {
	u.SetExcluded(block.FieldBlockNumber)
	return u
}

// AddBlockNumber adds v to the "block_number" field.
func (u *BlockUpsert) AddBlockNumber(v uint64) *BlockUpsert {
	u.Add(block.FieldBlockNumber, v)
	return u
}

// SetStateRoot sets the "state_root" field.
func (u *BlockUpsert) SetStateRoot(v string) *BlockUpsert {
	u.Set(block.FieldStateRoot, v)
	return u
}

// UpdateStateRoot sets the "state_root" field to the value that was provided on create.
func (u *BlockUpsert) UpdateStateRoot() *BlockUpsert {
	u.SetExcluded(block.FieldStateRoot)
	return u
}

// SetStatus sets the "status" field.
func (u *BlockUpsert) SetStatus(v block.Status) *BlockUpsert {
	u.Set(block.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *BlockUpsert) UpdateStatus() *BlockUpsert {
	u.SetExcluded(block.FieldStatus)
	return u
}

// SetTimestamp sets the "timestamp" field.
func (u *BlockUpsert) SetTimestamp(v time.Time) *BlockUpsert {
	u.Set(block.FieldTimestamp, v)
	return u
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *BlockUpsert) UpdateTimestamp() *BlockUpsert {
	u.SetExcluded(block.FieldTimestamp)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Block.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(block.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *BlockUpsertOne) UpdateNewValues() *BlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(block.FieldID)
		}
		if _, exists := u.create.mutation.Timestamp(); exists {
			s.SetIgnore(block.FieldTimestamp)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Block.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *BlockUpsertOne) Ignore() *BlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlockUpsertOne) DoNothing() *BlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlockCreate.OnConflict
// documentation for more info.
func (u *BlockUpsertOne) Update(set func(*BlockUpsert)) *BlockUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlockUpsert{UpdateSet: update})
	}))
	return u
}

// SetBlockHash sets the "block_hash" field.
func (u *BlockUpsertOne) SetBlockHash(v string) *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.SetBlockHash(v)
	})
}

// UpdateBlockHash sets the "block_hash" field to the value that was provided on create.
func (u *BlockUpsertOne) UpdateBlockHash() *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateBlockHash()
	})
}

// SetParentBlockHash sets the "parent_block_hash" field.
func (u *BlockUpsertOne) SetParentBlockHash(v string) *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.SetParentBlockHash(v)
	})
}

// UpdateParentBlockHash sets the "parent_block_hash" field to the value that was provided on create.
func (u *BlockUpsertOne) UpdateParentBlockHash() *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateParentBlockHash()
	})
}

// SetBlockNumber sets the "block_number" field.
func (u *BlockUpsertOne) SetBlockNumber(v uint64) *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.SetBlockNumber(v)
	})
}

// AddBlockNumber adds v to the "block_number" field.
func (u *BlockUpsertOne) AddBlockNumber(v uint64) *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.AddBlockNumber(v)
	})
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *BlockUpsertOne) UpdateBlockNumber() *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateBlockNumber()
	})
}

// SetStateRoot sets the "state_root" field.
func (u *BlockUpsertOne) SetStateRoot(v string) *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.SetStateRoot(v)
	})
}

// UpdateStateRoot sets the "state_root" field to the value that was provided on create.
func (u *BlockUpsertOne) UpdateStateRoot() *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateStateRoot()
	})
}

// SetStatus sets the "status" field.
func (u *BlockUpsertOne) SetStatus(v block.Status) *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *BlockUpsertOne) UpdateStatus() *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateStatus()
	})
}

// SetTimestamp sets the "timestamp" field.
func (u *BlockUpsertOne) SetTimestamp(v time.Time) *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.SetTimestamp(v)
	})
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *BlockUpsertOne) UpdateTimestamp() *BlockUpsertOne {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateTimestamp()
	})
}

// Exec executes the query.
func (u *BlockUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlockCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlockUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BlockUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: BlockUpsertOne.ID is not supported by MySQL driver. Use BlockUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BlockUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BlockCreateBulk is the builder for creating many Block entities in bulk.
type BlockCreateBulk struct {
	config
	builders []*BlockCreate
	conflict []sql.ConflictOption
}

// Save creates the Block entities in the database.
func (bcb *BlockCreateBulk) Save(ctx context.Context) ([]*Block, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Block, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BlockMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BlockCreateBulk) SaveX(ctx context.Context) []*Block {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BlockCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BlockCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Block.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BlockUpsert) {
//			SetBlockHash(v+v).
//		}).
//		Exec(ctx)
//
func (bcb *BlockCreateBulk) OnConflict(opts ...sql.ConflictOption) *BlockUpsertBulk {
	bcb.conflict = opts
	return &BlockUpsertBulk{
		create: bcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Block.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (bcb *BlockCreateBulk) OnConflictColumns(columns ...string) *BlockUpsertBulk {
	bcb.conflict = append(bcb.conflict, sql.ConflictColumns(columns...))
	return &BlockUpsertBulk{
		create: bcb,
	}
}

// BlockUpsertBulk is the builder for "upsert"-ing
// a bulk of Block nodes.
type BlockUpsertBulk struct {
	create *BlockCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Block.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(block.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *BlockUpsertBulk) UpdateNewValues() *BlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(block.FieldID)
				return
			}
			if _, exists := b.mutation.Timestamp(); exists {
				s.SetIgnore(block.FieldTimestamp)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Block.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *BlockUpsertBulk) Ignore() *BlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BlockUpsertBulk) DoNothing() *BlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BlockCreateBulk.OnConflict
// documentation for more info.
func (u *BlockUpsertBulk) Update(set func(*BlockUpsert)) *BlockUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BlockUpsert{UpdateSet: update})
	}))
	return u
}

// SetBlockHash sets the "block_hash" field.
func (u *BlockUpsertBulk) SetBlockHash(v string) *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.SetBlockHash(v)
	})
}

// UpdateBlockHash sets the "block_hash" field to the value that was provided on create.
func (u *BlockUpsertBulk) UpdateBlockHash() *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateBlockHash()
	})
}

// SetParentBlockHash sets the "parent_block_hash" field.
func (u *BlockUpsertBulk) SetParentBlockHash(v string) *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.SetParentBlockHash(v)
	})
}

// UpdateParentBlockHash sets the "parent_block_hash" field to the value that was provided on create.
func (u *BlockUpsertBulk) UpdateParentBlockHash() *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateParentBlockHash()
	})
}

// SetBlockNumber sets the "block_number" field.
func (u *BlockUpsertBulk) SetBlockNumber(v uint64) *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.SetBlockNumber(v)
	})
}

// AddBlockNumber adds v to the "block_number" field.
func (u *BlockUpsertBulk) AddBlockNumber(v uint64) *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.AddBlockNumber(v)
	})
}

// UpdateBlockNumber sets the "block_number" field to the value that was provided on create.
func (u *BlockUpsertBulk) UpdateBlockNumber() *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateBlockNumber()
	})
}

// SetStateRoot sets the "state_root" field.
func (u *BlockUpsertBulk) SetStateRoot(v string) *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.SetStateRoot(v)
	})
}

// UpdateStateRoot sets the "state_root" field to the value that was provided on create.
func (u *BlockUpsertBulk) UpdateStateRoot() *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateStateRoot()
	})
}

// SetStatus sets the "status" field.
func (u *BlockUpsertBulk) SetStatus(v block.Status) *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *BlockUpsertBulk) UpdateStatus() *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateStatus()
	})
}

// SetTimestamp sets the "timestamp" field.
func (u *BlockUpsertBulk) SetTimestamp(v time.Time) *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.SetTimestamp(v)
	})
}

// UpdateTimestamp sets the "timestamp" field to the value that was provided on create.
func (u *BlockUpsertBulk) UpdateTimestamp() *BlockUpsertBulk {
	return u.Update(func(s *BlockUpsert) {
		s.UpdateTimestamp()
	})
}

// Exec executes the query.
func (u *BlockUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the BlockCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for BlockCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BlockUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
