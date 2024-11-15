package models

type Medecin struct {
	Utilisateur
	Specialite     string         `json:"specialite" gorm:"not null"`
	Disponibilites []PlageHoraire `json:"disponibilites" gorm:"foreignKey:MedecinID"`
}

// TableName permet de définir le nom de la table en base de données
func (Medecin) TableName() string {
	return "medecins"
}
