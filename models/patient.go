package models

// Patient représente un patient, qui est aussi un utilisateur
type Patient struct {
	Utilisateur        // Composition : Patient hérite des attributs de Utilisateur
	Telephone   string `json:"telephone" gorm:"not null"`
	Assurance   string `json:"assurance"`
}

// TableName permet de définir le nom de la table en base de données
func (Patient) TableName() string {
	return "patients"
}
