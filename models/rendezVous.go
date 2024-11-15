package models

import "time"

type RendezVous struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	DateHeure time.Time `json:"date_heure" gorm:"not null"`
	Statut    string    `json:"statut" gorm:"not null"`
	PatientID int       `json:"patient_id" gorm:"not null"` // Clé étrangère vers Patient
	MedecinID int       `json:"medecin_id" gorm:"not null"` // Clé étrangère vers Medecin
}

func (RendezVous) TableName() string {
	return "rendezvous"
}
