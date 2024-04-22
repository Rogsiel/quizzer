package db

import "context"

type CreateOTPTxParams struct {
	Email   string `json:"email"`
	OtpCode string `json:"otp_code"`
	OtpType string `json:"otp_type"`
}

/*type CreateOTPTxResult struct {
	OtpID	int64	`json:"id"`
}*/

func (store *Store) CreateOTPTx(ctx context.Context, arg CreateOTPTxParams) (/*CreateOTPTxResult, */error) {
	//var Otp CreateOTPTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		_, err := q.CreateOTP(ctx, CreateOTPParams{
			Email: arg.Email,
			OtpCode: arg.OtpCode,
			OtpType: arg.OtpType,
		})
		if err != nil {
			return err
		}
		//Otp.OtpID = newOtp.ID
		return nil
	})
	return /*Otp, */err
}
