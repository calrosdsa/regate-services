package repository

import "context"

type BillingRepository interface {
	SendNotificationUserBilling(ctx context.Context, d []byte)
}

type BillingUseCase interface {
	SendNotificationUserBilling(ctx context.Context, d []byte)
}

// type BillingPayload struct {
// 	ProfileId int `json:"id"`
// 	Message   string `json:"message"`
// }
