package db

import (
	"context"
	"time"
)


type SignupTxParams struct{
	UserName		string	`json:"user_name"`
	Email			string	`json:"email"`
	HashedPassword	string	`json:"hashed_password"`
}

type UserTxResult struct{
	ID					int64		`json:"id"`
	UserName			string		`json:"user_name"`
	Email				string		`json:"email"`
	CreatedAt			time.Time	`json:"created_at"`
	PasswordChangedAt	time.Time	`json:"password_changed_at"`
}

func (store *Store) SignupTx(ctx context.Context, arg SignupTxParams) (UserTxResult, error) {
	var user UserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		result, err := q.CreateUser(ctx, CreateUserParams{
			UserName: arg.UserName,
			Email: arg.Email,
			HashedPassword: arg.HashedPassword,
		})
		if err != nil {
			return err
		}
		user = UserTxResult{
			ID: result.ID,
			UserName: result.UserName,
			Email: result.Email,
			CreatedAt: result.CreatedAt,
			PasswordChangedAt: result.PasswordChangedAt,
		}
		return nil
	})
	return user, err
}

type GetUserTxParams struct{
	UserName	string	`json:"user_name"`
}

func (store *Store) GetUserTx(ctx context.Context, arg GetUserTxParams) (UserTxResult, error) {
	var user UserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		result, err := q.GetUser(ctx, arg.UserName)
		if err != nil {
			return err
		}
		user = UserTxResult{
			ID: result.ID,
			UserName: result.UserName,
			Email: result.Email,
			CreatedAt: result.CreatedAt,
			PasswordChangedAt: result.PasswordChangedAt,
		}
		return nil
	})
	return user, err
}

type GetUsersTxParams struct {
	Limit	int32	`json:"limit"`
	Offset	int32	`json:"offset"`
}

type GetUsersTxResult []GetUsersRow

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
	UserName       string `json:"user_name"`
	HashedPassword string `json:"hashed_password"`
}

func (store *Store) UpdatePasswordTx(ctx context.Context, arg UpdatePasswordTxParams) (error) {
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.UpdatePassword(ctx, UpdatePasswordParams{
			UserName: arg.UserName,
			HashedPassword: arg.HashedPassword,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return err
} 

type VerifyEmailTxParams struct {
	OtpID	int64	`json:"id"`
	Email	string	`json:"email"`
	OtpCode	string	`json:"otp_code"`
}

type VerifyEmailTxResult struct {
	UserName	string	`json:"user_name"`
}
func (store *Store) VerifyEmailTx(ctx context.Context, arg VerifyEmailTxParams) (VerifyEmailTxResult, error) {
	var results VerifyEmailTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		err := q.UpdateOTP(ctx, UpdateOTPParams{
			ID: arg.OtpID,
			OtpCode: arg.OtpCode,
		})

		if err != nil {
			return err
		}
		user, err := q.VerifyEmail(ctx, arg.Email)
		if err != nil {
			return err
		}
		results.UserName = user.UserName
		return nil
	})
	return results, err
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
