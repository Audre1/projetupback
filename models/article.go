package models

import "time"

type Article struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Titre           string    `json:"titre" gorm:"not null"`
	Contenu         string    `json:"contenu" gorm:"not null"`
	DatePublication time.Time `json:"date_publication" gorm:"not null"`
	AuteurID        int       `json:"auteur_id" gorm:"not null"` // Clé étrangère vers Utilisateur (auteur)
	Categorie       string    `json:"categorie"`
	Statut          string    `json:"statut" gorm:"not null"`
}

// TableName permet de définir le nom de la table en base de données
func (Article) TableName() string {
	return "articles"
}
