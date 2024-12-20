package dto

import "time"

type NotificationResponse struct {
	NotificationId string    `json:"notification_id"`
	ReservationId  string    `json:"reservation_id"`
	Content        string    `json:"content"`
	Time           time.Time `json:"time"`
}
