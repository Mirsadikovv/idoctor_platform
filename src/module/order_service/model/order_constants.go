package order_model

// This file contains all constants related to Order model
// Usage example:
//   order.Status = order_model.StatusDiagnostic
//   order.PaymentType = order_model.PaymentTypeCash
//   if order_model.IsValidStatus(status) { ... }

// Order Status Constants
const (
	StatusPending    = "PENDING"
	StatusReceived   = "RECEIVED"
	StatusDiagnostic = "DIAGNOSTIC"
	StatusCompleted  = "COMPLETED"
	StatusCancelled  = "CANCELLED"
)

// Payment Type Constants
const (
	PaymentTypeCash   = "CASH"
	PaymentTypeCard   = "CARD"
	PaymentTypeOnline = "ONLINE"
)

// Payment Status Constants
const (
	PaymentStatusPending   = "PENDING"
	PaymentStatusPaid      = "PAID"
	PaymentStatusRefunded  = "REFUNDED"
	PaymentStatusCancelled = "CANCELLED"
)

// ValidOrderStatuses returns all valid order statuses
var ValidOrderStatuses = []string{
	StatusPending,
	StatusReceived,
	StatusDiagnostic,
	StatusCompleted,
	StatusCancelled,
}

// ValidPaymentTypes returns all valid payment types
var ValidPaymentTypes = []string{
	PaymentTypeCash,
	PaymentTypeCard,
	PaymentTypeOnline,
}

// ValidPaymentStatuses returns all valid payment statuses
var ValidPaymentStatuses = []string{
	PaymentStatusPending,
	PaymentStatusPaid,
	PaymentStatusRefunded,
	PaymentStatusCancelled,
}

// IsValidStatus checks if the given status is valid
func IsValidStatus(status string) bool {
	for _, s := range ValidOrderStatuses {
		if s == status {
			return true
		}
	}
	return false
}

// IsValidPaymentType checks if the given payment type is valid
func IsValidPaymentType(paymentType string) bool {
	for _, pt := range ValidPaymentTypes {
		if pt == paymentType {
			return true
		}
	}
	return false
}

// IsValidPaymentStatus checks if the given payment status is valid
func IsValidPaymentStatus(paymentStatus string) bool {
	for _, ps := range ValidPaymentStatuses {
		if ps == paymentStatus {
			return true
		}
	}
	return false
}
