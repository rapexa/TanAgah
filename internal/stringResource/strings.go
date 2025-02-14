package stringResource

import "github.com/gin-gonic/gin"

type ImplementationStrings struct {
	StringsImpl Strings
}

func (pr *ImplementationStrings) DoubleSpending(c *gin.Context) string {
	defaultStr := "This transaction has already been processed."
	switch GetLng(c) {
	case "fa":
		return " این تراکنش قبلاً پردازش شده است."
	}
	return defaultStr
}

func (pr *ImplementationStrings) PasswordError(c *gin.Context) string {
	defaultStr := "Password must be more than 8 characters and requires a number, lowercase letter, uppercase letter and a symbol."
	switch GetLng(c) {
	case "fa":
		return "رمز عبور باید بیش از ۸ کاراکتر باشد و نیازمند یک عدد، یک حرف کوچک، یک حرف بزرگ و یک نماد است."
	}
	return defaultStr
}

func (pr *ImplementationStrings) PaymentCancelled(c *gin.Context) string {
	defaultStr := "Payment cancelled by the user"
	switch GetLng(c) {
	case "fa":
		return "عملیات پرداخت توسط کاربر لفو شده است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) PaymentPending(c *gin.Context) string {
	defaultStr := "The payment is still pending and has not been completed yet"
	switch GetLng(c) {
	case "fa":
		return "پرداخت هنوز در انتظار است و هنوز تکمیل نشده است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) PaymentSucceeded(c *gin.Context) string {
	defaultStr := "Payment succeeded: The transaction has been completed successfully."
	switch GetLng(c) {
	case "fa":
		return "پرداخت با موفقیت انجام شد."
	}
	return defaultStr
}

func (pr *ImplementationStrings) UserNotFound(c *gin.Context) string {
	defaultStr := "User not found"
	switch GetLng(c) {
	case "fa":
		return " کاربر پیدا نشد"
	}
	return defaultStr
}

func (pr *ImplementationStrings) UserDeleteSuccess(c *gin.Context) string {
	defaultStr := "User deleted successfully"
	switch GetLng(c) {
	case "fa":
		return "حساب کاربر با موفقیت پاک شد"
	}
	return defaultStr
}

func (pr *ImplementationStrings) UserIdIsRequired(c *gin.Context) string {
	defaultStr := "User id is required"
	switch GetLng(c) {
	case "fa":
		return "شناسه کاربر ضروری است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) OtpIsNotValid(c *gin.Context) string {
	defaultStr := "Code is not valid"
	switch GetLng(c) {
	case "fa":
		return "کد ورود اشتباه است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) UnknownError(c *gin.Context) string {
	defaultStr := "Unknown error please try again"
	switch GetLng(c) {
	case "fa":
		return "مشکل ناشناخته لطفا دوباره تلاش کنید"
	}
	return defaultStr
}

func (pr *ImplementationStrings) OtpDescription(c *gin.Context, otp string) string {
	defaultStr := "XVPN OTP Code: " + otp
	switch GetLng(c) {
	case "fa":
		return " کد ورود به برنامه XVPN " + otp
	}
	return defaultStr
}

func (pr *ImplementationStrings) BadRequest(c *gin.Context) string {
	defaultStr := "Bad request the request is not valid"
	switch GetLng(c) {
	case "fa":
		return "درخواست نامعتبر است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) OopsUsernameOrPassword(c *gin.Context) string {
	defaultStr := "Oops! Username or password is not valid"
	switch GetLng(c) {
	case "fa":
		return "نام کاربری یا کلمه عبور اشتباه است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) AccessDenied(c *gin.Context) string {
	defaultStr := "Access denied you are not authorized to access this resource"
	switch GetLng(c) {
	case "fa":
		return "دسرسی به این منابع غیر مجاز است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) OtpNotValid(c *gin.Context) string {
	defaultStr := "Otp code not valid"
	switch GetLng(c) {
	case "fa":
		return "رمز ورود اشتباه است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) OtpTryCountIsToLong(c *gin.Context) string {
	defaultStr := "You have tried more than allowed. Please wait until the login code expires"
	switch GetLng(c) {
	case "fa":
		return "بیش از حد مجاز تلاش کردید لطفا تا منقضی شدن کد ورود صبر کنید"
	}
	return defaultStr
}

func (pr *ImplementationStrings) SessionExpired(c *gin.Context) string {
	defaultStr := "Session is expired"
	switch GetLng(c) {
	case "fa":
		return " نشست شما منقضی شده است"
	}
	return defaultStr
}

func (pr *ImplementationStrings) NotEnoughCharge(c *gin.Context) string {
	defaultStr := "Your current balance is insufficient to complete this action"
	switch GetLng(c) {
	case "fa":
		return " موجودی فعلی شما برای انجام این عمل ناکافی است."
	}
	return defaultStr
}
