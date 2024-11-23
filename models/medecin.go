package models

// Medicin représente un médecin specialiste travaillant dans la clinique
type Medecin struct {
	Utilisateur
	ID                int    `json:"id"`                 // Identifiant unique du médecin
	Speciality        string `json:"speciality"`         // Spécialité du médecin (par ex., cardiologie)
	ConsultationHours string `json:"consultation_hours"` // Horaires de consultation
}

// TableName permet de définir le nom de la table en base de données
func (Medecin) TableName() string {
	return "medecins"
}
