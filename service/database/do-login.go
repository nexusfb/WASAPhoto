package database

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

func (db *appdbimpl) DoLogin(username string) (string, error) {
	// crea profilo con username
	// dubbio 1: si puo creare profilo con solo quello?? in teoria no perchè userID è not null
	// risposta 1? -> generatore user ID
	// dubbio 2: generare id prima o dopo insert? direi prima perche credo non si possa fare dopo ma se poi lui ci da username che abbiamo gia come si fa??
	// plan:
	// 1- check se esiste
	// 2- se esiste errore -> user already exists
	// 2- se non esiste creare id -> provare a creare profilo

	// 1) check se user già esiste
	// a) usa metodo per prendere profilo
	p, err := db.GetUserProfile(username)
	// b) se non ci sono errori -> allora esiste
	if err == nil {
		// siamo nel caso che l'utente gia esiste, facciamo login con questo utente
		return p.UserID, nil
	} else if err != ErrUserProfileDoesNotExists {
		// se errore diverso da non esiste (che invece è quello che cerchiamo)
		return "00000000", fmt.Errorf("error encountered while checking if profile exists: %w", err)
	}

	//posso evitare un altro if -> la funzione GetProfile ha solo 3 tipi di err return -> nil, err, ErrrUserDoesNotExists
	// se sono qui significa che user non esiste
	// posso quindi creare id
	rawUid := ksuid.New()
	uid := rawUid.String()
	_, err = db.c.Exec(`INSERT INTO users (userid, username) VALUES (?,?)`,
		uid, username)
	if err != nil {
		return "00000000", fmt.Errorf("error when creating new user profile: %w", err)
	}
	return uid, nil
}
