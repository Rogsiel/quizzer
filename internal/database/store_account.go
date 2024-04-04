package db

import (
	"context"
)


type SignupTxParams struct{
	Name	string	`json:"name"`
	Email	string	`json:"email"`
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
		})
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

type GetUserTxParams struct{
	ID	int64	`json:"id"`
}

func (store *Store) GetUserTx(ctx context.Context, arg GetUserTxParams) (UserTxResult, error) {
	var user UserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		user.User, err = q.GetUser(ctx, arg.ID)
		if err != nil {
			return err
		}
		return nil
	})
	return user, err
}

type GetUserByNameTxParams struct {
	Name	string	`json:"name"`
}

type GetUserByNameTxResult struct {
	IDs		[]int64		`json:"ids"`
	Names	[]string	`json:"names"`
}

func (store *Store) GetUserByNameTx(ctx context.Context, arg GetUserByNameTxParams) (GetUserByNameTxResult, error) {
	var SearchResult GetUserByNameTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		UserList, err := q.GetUserUsername(ctx, arg.Name)
		if err != nil {
			return err
		}
		for i := 0; i < len(UserList); i++ {
			SearchResult.IDs = append(SearchResult.IDs, UserList[i].ID)
			SearchResult.Names = append(SearchResult.Names, UserList[i].Name)
		}
		return nil	
	})
	return SearchResult, err
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

type UpdateUserTxParams struct {
	ID		int64	`json:"id"`
	NewName	string	`json:"name"`
}

func (store *Store) UpdateUserTx(ctx context.Context, arg UpdateUserTxParams) (UserTxResult, error) {
	var UpdatedUser UserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		UpdatedUser.User, err = q.UpdateUser(ctx, UpdateUserParams{
			ID: arg.ID,
			Name: arg.NewName,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return UpdatedUser, err
}

type DeleteUserTxParams int64

func (store *Store) DeleteUserTx(ctx context.Context, arg DeleteUserTxParams) (error) {
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.DeleteUser(ctx, int64(arg))
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
