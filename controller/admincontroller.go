package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"projetupback/models" // Import des modèles
)

// AjouterMedicin permet à l'administrateur d'ajouter un médecin
func AjouterMedecin(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Décodage de la requête
	var medecin models.Medecin
	if err := json.NewDecoder(r.Body).Decode(&medecin); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	// Insertion dans la base de données
	query := "INSERT INTO medecins (speciality, consultation_hours) VALUES (?, ?)"
	_, err := db.Exec(query, medecin.Speciality, medecin.ConsultationHours)
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Médecin ajouté avec succès"))
}

// SupprimerMedicin permet à l'admin de supprimer un médecin
func SupprimerMedecin(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM medecins WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Médecin supprimé avec succès"))
}
