package otp

type OTPManager interface {
    NewEmailVerificationOTP(email string) OTP
    NewPasswordChangeOTP(email string) OTP
    VerifyOTP(otpType string) error
}
