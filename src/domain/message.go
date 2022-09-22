package domain

import "time"

type Message struct {
	ID          uint      `json:"id"`
	SenderId    uint      `json:"sender_id"`
	RecipientId uint      `json:"Recipient_id"`
	Type        string    `json:"type"`
	Body        string    `json:"body"`
	SentAt      time.Time `json:"sent_at"`
	SeenAt      time.Time `json:"seen_at"`
	PulledAt    time.Time `json:"pulled_at"`
}
