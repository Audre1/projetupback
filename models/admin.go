package models

// Import de la bibliothèque GORM pour la gestion des bases de données

// Définition du modèle Admin
type Admin struct {
	ID       uint   `json:"id"`                  // Identifiant unique de l'administrateur (clé primaire)
	Email    string `json:"email" gorm:"unique"` // Email de l'administrateur, doit être unique dans la base de données
	Password string `json:"password"`            // Mot de passe de l'administrateur
}

// TableName permet de définir le nom de la table en base de données
func (Admin) TableName() string {
	return "admins"
}
