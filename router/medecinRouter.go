package router

import (
	"database/sql"
	"net/http"

	"projetupback/controller" // Import des contrôleurs
)

// ConfigurerMedecinRoutes configure les routes pour les fonctionnalités du médecin
func ConfigurerMedecinRoutes(db *sql.DB) {
	http.HandleFunc("/medic/rendez-vous", func(w http.ResponseWriter, r *http.Request) {
		controller.ConsulterRendezVous(db, w, r)
	})
}
