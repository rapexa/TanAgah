package stringResource

import "github.com/gin-gonic/gin"

type Strings interface {
	TokenJwtIsRequired(c *gin.Context) string
	TokenJwtIsNotValid(c *gin.Context) string
	UserNotFound(c *gin.Context) string
	NoFilesUploaded(c *gin.Context) string
	UserDeleteSuccess(c *gin.Context) string
	RetrieveHistoryError(c *gin.Context) string
	WebSocketUpgradeError(c *gin.Context) string
	OtpIsNotValid(c *gin.Context) string
	BadRequest(c *gin.Context) string
	UnknownError(c *gin.Context) string
	PasswordError(c *gin.Context) string
	NotEnoughCharge(c *gin.Context) string
	PaymentCancelled(c *gin.Context) string
	DoubleSpending(c *gin.Context) string
	PaymentPending(c *gin.Context) string
	PaymentSucceeded(c *gin.Context) string
	OtpDescription(c *gin.Context, otp string) string
	OtpNotValid(c *gin.Context) string
	OtpTryCountIsToLong(c *gin.Context) string
	OopsUsernameOrPassword(c *gin.Context) string
	AccessDenied(c *gin.Context) string
	SessionExpired(c *gin.Context) string
}
