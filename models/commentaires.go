package models

import "time"

type Commentaire struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Contenu         string    `json:"contenu" gorm:"not null"`
	DateCommentaire time.Time `json:"date_commentaire" gorm:"not null"`
	ArticleID       int       `json:"article_id" gorm:"not null"` // Clé étrangère vers Article
	AuteurID        int       `json:"auteur_id" gorm:"not null"`  // Clé étrangère vers Utilisateur (auteur)
}

// TableName permet de définir le nom de la table en base de données
func (Commentaire) TableName() string {
	return "commentaires"
}
