package repository



type SalaEstado int16

const (
	SalaAvailable SalaEstado = iota
	SalaUnAvailable
	SalaReserved 
)

type NotificationType int64

const (
	//"0"
	NotificationMessageGroup NotificationType = iota
	//"1"
	NotificationMessageComplejo 
	NotificationSalaCreation
	NotificationSalaReservationConflict
	NotificationSalaHasBeenReserved
	NotificationBilling
)