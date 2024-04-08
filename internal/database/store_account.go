package db

import (
	"context"
)


type SignupTxParams struct{
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	HashedPassword	string	`json:"hashed_password"`
}

type UserTxResult struct{
	User	User	`json:"user"`
}

func (store *Store) SignupTx(ctx context.Context, arg SignupTxParams) (UserTxResult, error) {
	var user UserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		user.User, err = q.CreateUser(ctx, CreateUserParams{
			Name: arg.Name,
			Email: arg.Email,
			HashedPassword: arg.HashedPassword,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

type GetUserTxParams struct{
	Name	string	`json:"name"`
}

func (store *Store) GetUserTx(ctx context.Context, arg GetUserTxParams) (UserTxResult, error) {
	var user UserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		user.User, err = q.GetUser(ctx, arg.Name)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

type GetUsersTxParams struct {
	Limit	int32	`json:"limit"`
	Offset	int32	`json:"offset"`
}

type GetUsersTxResult []User

func (store *Store) GetUsersTx(ctx context.Context, arg GetUsersTxParams) (GetUsersTxResult, error) {
	var Users GetUsersTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		Users, err = q.GetUsers(ctx, GetUsersParams{
			Limit: arg.Limit,
			Offset: arg.Offset,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return Users, err
}

type UpdatePasswordTxParams struct {
	Name				string	`json:"name"`
	NewHashedPassword	string	`json:"password"`
}

func (store *Store) UpdatePasswordTx(ctx context.Context, arg UpdatePasswordTxParams) (UserTxResult, error) {
	var UpdatedUser UserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		UpdatedUser.User, err = q.UpdatePassword(ctx, UpdatePasswordParams{
			Name: arg.Name,
			HashedPassword: arg.NewHashedPassword,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return UpdatedUser, err
}

type DeleteUserTxParams string

func (store *Store) DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) (error) {
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteUser(ctx, string(arg))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
