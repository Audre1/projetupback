package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"projetupback/models"
)

// AjouterArticle permet à l'admin d'ajouter un article de blog
func AjouterArticle(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var article models.BlogArticle
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO blogarticles (titre, contenu, DatePublication, AuteurID, Categorie, Statut) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, article.Titre, article.Contenu, time.Now().Format("2006-01-02"))
	if err != nil {
		http.Error(w, "Erreur lors de l'ajout", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Article ajouté avec succès"))
}

// ModifierArticle permet de modifier un article de blog
func ModifierArticle(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	var article models.BlogArticle
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Données invalides", http.StatusBadRequest)
		return
	}

	query := "UPDATE blogarticle SET titre = ?, contenu = ? WHERE id = ?"
	_, err := db.Exec(query, article.Titre, article.Contenu, id)
	if err != nil {
		http.Error(w, "Erreur lors de la modification", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Article modifié avec succès"))
}

// SupprimerArticle permet de supprimer un article
func SupprimerArticle(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID manquant", http.StatusBadRequest)
		return
	}

	query := "DELETE FROM blogarticle WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, "Erreur lors de la suppression", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Article supprimé avec succès"))
}

// ListeArticles retourne la liste des articles
func ListeArticles(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, titree, contenu, date FROM blog")
	if err != nil {
		http.Error(w, "Erreur lors de la récupération des articles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var articles []models.BlogArticle
	for rows.Next() {
		var article models.BlogArticle
		if err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Date); err != nil {
			http.Error(w, "Erreur lors du traitement des données", http.StatusInternalServerError)
			return
		}
		articles = append(articles, article)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}
