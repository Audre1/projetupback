package models

type Admin struct {
	Utilisateur
}

// TableName permet de définir le nom de la table en base de données
func (Admin) TableName() string {
	return "admins"
}
