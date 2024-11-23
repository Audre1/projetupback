package router

import (
	"database/sql"
	"net/http"

	"projetupback/controller" // Import des contrôleurs
)

// ConfigurerAdminRoutes configure les routes pour les fonctionnalités de l'admin
func ConfigurerAdminRoutes(db *sql.DB) {
	http.HandleFunc("/admin/ajouter-medecin", func(w http.ResponseWriter, r *http.Request) {
		controller.AjouterMedecin(db, w, r)
	})

	http.HandleFunc("/admin/supprimer-medecin", func(w http.ResponseWriter, r *http.Request) {
		controller.SupprimerMedecin(db, w, r)
	})
}
