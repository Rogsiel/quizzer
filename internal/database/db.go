// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createQuizStmt, err = db.PrepareContext(ctx, createQuiz); err != nil {
		return nil, fmt.Errorf("error preparing query CreateQuiz: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getCorrectAnswersStmt, err = db.PrepareContext(ctx, getCorrectAnswers); err != nil {
		return nil, fmt.Errorf("error preparing query GetCorrectAnswers: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.incrementAnsweredCountStmt, err = db.PrepareContext(ctx, incrementAnsweredCount); err != nil {
		return nil, fmt.Errorf("error preparing query IncrementAnsweredCount: %w", err)
	}
	if q.sendAnswersStmt, err = db.PrepareContext(ctx, sendAnswers); err != nil {
		return nil, fmt.Errorf("error preparing query SendAnswers: %w", err)
	}
	if q.updateScoreStmt, err = db.PrepareContext(ctx, updateScore); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateScore: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createQuizStmt != nil {
		if cerr := q.createQuizStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createQuizStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getCorrectAnswersStmt != nil {
		if cerr := q.getCorrectAnswersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCorrectAnswersStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.incrementAnsweredCountStmt != nil {
		if cerr := q.incrementAnsweredCountStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing incrementAnsweredCountStmt: %w", cerr)
		}
	}
	if q.sendAnswersStmt != nil {
		if cerr := q.sendAnswersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing sendAnswersStmt: %w", cerr)
		}
	}
	if q.updateScoreStmt != nil {
		if cerr := q.updateScoreStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateScoreStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                         DBTX
	tx                         *sql.Tx
	createQuizStmt             *sql.Stmt
	createUserStmt             *sql.Stmt
	deleteUserStmt             *sql.Stmt
	getCorrectAnswersStmt      *sql.Stmt
	getUserStmt                *sql.Stmt
	getUsersStmt               *sql.Stmt
	incrementAnsweredCountStmt *sql.Stmt
	sendAnswersStmt            *sql.Stmt
	updateScoreStmt            *sql.Stmt
	updateUserStmt             *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                         tx,
		tx:                         tx,
		createQuizStmt:             q.createQuizStmt,
		createUserStmt:             q.createUserStmt,
		deleteUserStmt:             q.deleteUserStmt,
		getCorrectAnswersStmt:      q.getCorrectAnswersStmt,
		getUserStmt:                q.getUserStmt,
		getUsersStmt:               q.getUsersStmt,
		incrementAnsweredCountStmt: q.incrementAnsweredCountStmt,
		sendAnswersStmt:            q.sendAnswersStmt,
		updateScoreStmt:            q.updateScoreStmt,
		updateUserStmt:             q.updateUserStmt,
	}
}
