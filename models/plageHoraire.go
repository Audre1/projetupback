package models

import "time"

type PlageHoraire struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Date       time.Time `json:"date" gorm:"not null"`
	HeureDebut time.Time `json:"heure_debut" gorm:"not null"`
	HeureFin   time.Time `json:"heure_fin" gorm:"not null"`
	MedecinID  int       `json:"medecin_id" gorm:"not null"` // Clé étrangère vers Medecin
}

// TableName permet de définir le nom de la table en base de données
func (PlageHoraire) TableName() string {
	return "plages_horaires"
}
