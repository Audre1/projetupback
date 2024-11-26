package controller

import (
	"database/sql"
	"net/http"
)

// ConsulterRendezVous permet au médecin de consulter ses rendez-vous
func ConsulterRendezVous(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	// Exemple basique pour récupérer des rendez-vous liés à un médecin
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID médecins manquant", http.StatusBadRequest)
		return
	}

	query := "SELECT * FROM rendezvous WHERE medecins_id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des rendez-vous", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Traitement des données à renvoyer
	w.Write([]byte("Liste des rendez-vous du médecin"))
}
