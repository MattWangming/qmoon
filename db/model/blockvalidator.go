// Package model contains the types for schema 'public'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
)

// BlockValidator represents a row from 'public.block_validators'.
type BlockValidator struct {
	ID               int64          `json:"id"`                // id
	ChainID          sql.NullString `json:"chain_id"`          // chain_id
	Height           sql.NullInt64  `json:"height"`            // height
	ValidatorAddress sql.NullString `json:"validator_address"` // validator_address
	ValidatorIndex   sql.NullInt64  `json:"validator_index"`   // validator_index
	Type             sql.NullInt64  `json:"type"`              // type
	Round            sql.NullInt64  `json:"round"`             // round
	Signature        sql.NullString `json:"signature"`         // signature
	VotingPower      sql.NullInt64  `json:"voting_power"`      // voting_power
	Accum            sql.NullInt64  `json:"accum"`             // accum
	Time             pq.NullTime    `json:"time"`              // time
	CreatedAt        pq.NullTime    `json:"created_at"`        // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the BlockValidator exists in the database.
func (bv *BlockValidator) Exists() bool {
	return bv._exists
}

// Deleted provides information if the BlockValidator has been deleted from the database.
func (bv *BlockValidator) Deleted() bool {
	return bv._deleted
}

// Insert inserts the BlockValidator to the database.
func (bv *BlockValidator) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if bv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.block_validators (` +
		`chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, bv.ChainID, bv.Height, bv.ValidatorAddress, bv.ValidatorIndex, bv.Type, bv.Round, bv.Signature, bv.VotingPower, bv.Accum, bv.Time, bv.CreatedAt)
	err = db.QueryRow(sqlstr, bv.ChainID, bv.Height, bv.ValidatorAddress, bv.ValidatorIndex, bv.Type, bv.Round, bv.Signature, bv.VotingPower, bv.Accum, bv.Time, bv.CreatedAt).Scan(&bv.ID)
	if err != nil {
		return err
	}

	// set existence
	bv._exists = true

	return nil
}

// Update updates the BlockValidator in the database.
func (bv *BlockValidator) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !bv._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if bv._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.block_validators SET (` +
		`chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) WHERE id = $12`

	// run query
	XOLog(sqlstr, bv.ChainID, bv.Height, bv.ValidatorAddress, bv.ValidatorIndex, bv.Type, bv.Round, bv.Signature, bv.VotingPower, bv.Accum, bv.Time, bv.CreatedAt, bv.ID)
	_, err = db.Exec(sqlstr, bv.ChainID, bv.Height, bv.ValidatorAddress, bv.ValidatorIndex, bv.Type, bv.Round, bv.Signature, bv.VotingPower, bv.Accum, bv.Time, bv.CreatedAt, bv.ID)
	return err
}

// Save saves the BlockValidator to the database.
func (bv *BlockValidator) Save(db XODB) error {
	if bv.Exists() {
		return bv.Update(db)
	}

	return bv.Insert(db)
}

// Upsert performs an upsert for BlockValidator.
//
// NOTE: PostgreSQL 9.5+ only
func (bv *BlockValidator) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if bv._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.block_validators (` +
		`id, chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.chain_id, EXCLUDED.height, EXCLUDED.validator_address, EXCLUDED.validator_index, EXCLUDED.type, EXCLUDED.round, EXCLUDED.signature, EXCLUDED.voting_power, EXCLUDED.accum, EXCLUDED.time, EXCLUDED.created_at` +
		`)`

	// run query
	XOLog(sqlstr, bv.ID, bv.ChainID, bv.Height, bv.ValidatorAddress, bv.ValidatorIndex, bv.Type, bv.Round, bv.Signature, bv.VotingPower, bv.Accum, bv.Time, bv.CreatedAt)
	_, err = db.Exec(sqlstr, bv.ID, bv.ChainID, bv.Height, bv.ValidatorAddress, bv.ValidatorIndex, bv.Type, bv.Round, bv.Signature, bv.VotingPower, bv.Accum, bv.Time, bv.CreatedAt)
	if err != nil {
		return err
	}

	// set existence
	bv._exists = true

	return nil
}

// Delete deletes the BlockValidator from the database.
func (bv *BlockValidator) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !bv._exists {
		return nil
	}

	// if deleted, bail
	if bv._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.block_validators WHERE id = $1`

	// run query
	XOLog(sqlstr, bv.ID)
	_, err = db.Exec(sqlstr, bv.ID)
	if err != nil {
		return err
	}

	// set deleted
	bv._deleted = true

	return nil
}

// BlockValidatorsQuery returns offset-limit rows from 'public.block_validators' filte by filter,
// ordered by "id" in descending order.
func BlockValidatorFilter(db XODB, filter string, offset, limit int64) ([]*BlockValidator, error) {
	sqlstr := `SELECT ` +
		`id, chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at` +
		` FROM public.block_validators `

	if filter != "" {
		sqlstr = sqlstr + " WHERE " + filter
	}

	if limit > 0 {
		sqlstr = sqlstr + fmt.Sprintf(" offset %d limit %d", offset, limit)
	}

	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	var res []*BlockValidator
	for q.Next() {
		bv := BlockValidator{}

		// scan
		err = q.Scan(&bv.ID, &bv.ChainID, &bv.Height, &bv.ValidatorAddress, &bv.ValidatorIndex, &bv.Type, &bv.Round, &bv.Signature, &bv.VotingPower, &bv.Accum, &bv.Time, &bv.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &bv)
	}

	return res, nil
}

// BlockValidatorsByValidatorAddress retrieves a row from 'public.block_validators' as a BlockValidator.
//
// Generated from index 'block_validators_address_idx'.
func BlockValidatorsByValidatorAddress(db XODB, validatorAddress sql.NullString) ([]*BlockValidator, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at ` +
		`FROM public.block_validators ` +
		`WHERE validator_address = $1`

	// run query
	XOLog(sqlstr, validatorAddress)
	q, err := db.Query(sqlstr, validatorAddress)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*BlockValidator{}
	for q.Next() {
		bv := BlockValidator{
			_exists: true,
		}

		// scan
		err = q.Scan(&bv.ID, &bv.ChainID, &bv.Height, &bv.ValidatorAddress, &bv.ValidatorIndex, &bv.Type, &bv.Round, &bv.Signature, &bv.VotingPower, &bv.Accum, &bv.Time, &bv.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &bv)
	}

	return res, nil
}

// BlockValidatorsByChainIDHeight retrieves a row from 'public.block_validators' as a BlockValidator.
//
// Generated from index 'block_validators_chain_id_height_idx'.
func BlockValidatorsByChainIDHeight(db XODB, chainID sql.NullString, height sql.NullInt64) ([]*BlockValidator, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at ` +
		`FROM public.block_validators ` +
		`WHERE chain_id = $1 AND height = $2`

	// run query
	XOLog(sqlstr, chainID, height)
	q, err := db.Query(sqlstr, chainID, height)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*BlockValidator{}
	for q.Next() {
		bv := BlockValidator{
			_exists: true,
		}

		// scan
		err = q.Scan(&bv.ID, &bv.ChainID, &bv.Height, &bv.ValidatorAddress, &bv.ValidatorIndex, &bv.Type, &bv.Round, &bv.Signature, &bv.VotingPower, &bv.Accum, &bv.Time, &bv.CreatedAt)
		if err != nil {
			return nil, err
		}

		res = append(res, &bv)
	}

	return res, nil
}

// BlockValidatorByID retrieves a row from 'public.block_validators' as a BlockValidator.
//
// Generated from index 'block_validators_pkey'.
func BlockValidatorByID(db XODB, id int64) (*BlockValidator, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, chain_id, height, validator_address, validator_index, type, round, signature, voting_power, accum, time, created_at ` +
		`FROM public.block_validators ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	bv := BlockValidator{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&bv.ID, &bv.ChainID, &bv.Height, &bv.ValidatorAddress, &bv.ValidatorIndex, &bv.Type, &bv.Round, &bv.Signature, &bv.VotingPower, &bv.Accum, &bv.Time, &bv.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &bv, nil
}