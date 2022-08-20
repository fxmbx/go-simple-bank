package utils

const (
	CAD = "CAD"
	NGN = "NGN"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case CAD, NGN:
		return true
	default:
		return false
	}
}
