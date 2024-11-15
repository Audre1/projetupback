package models

type Utilisateur struct {
	ID         int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nom        string `json:"nom" gorm:"not null"`
	Prenom     string `json:"prenom" gorm:"not null"`
	Email      string `json:"email" gorm:"unique;not null"`
	MotDePasse string `json:"mot_de_passe" gorm:"not null"`
	Role       string `json:"role" gorm:"not null"`
}

// TableName permet de définir le nom de la table en base de données
func (Utilisateur) TableName() string {
	return "utilisateurs"
}
