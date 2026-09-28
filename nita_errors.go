// Couche mince écrite à la main (NON générée) — voir nita_client.go.
package nita

import (
	"fmt"
	"strings"
)

// NitaError est l'erreur métier NITA : produite par Unwrap quand une enveloppe de
// réponse a un status autre que "success" — souvent renvoyée en HTTP 200 avec le
// détail dans l'enveloppe, donc le SDK généré ne renvoie aucune erreur Go de
// lui-même dans ce cas : c'est Unwrap qui la transforme en erreur typée.
type NitaError struct {
	// Status canonique renvoyé par l'enveloppe (ERROR, ALERT, NOT_FOUND, CONFLICT…).
	// "ERROR" si l'enveloppe elle-même est absente/nil.
	Status string
	// Code applicatif (400, 404, 409, 500…).
	Code int
	// Message : celui de l'enveloppe si présent, sinon un message par défaut
	// ("NITA a répondu <Status>").
	Message string
	// Data : charge utile éventuelle accompagnant l'erreur (typiquement
	// map[string]interface{}, mais dépend de l'endpoint).
	Data any
}

// Error implémente l'interface error.
func (e *NitaError) Error() string {
	return e.Message
}

func newNitaError(status string, code int32, message string, data any) *NitaError {
	if status == "" {
		status = "ERROR"
	}
	if message == "" {
		message = fmt.Sprintf("NITA a répondu %s", status)
	}
	return &NitaError{Status: status, Code: int(code), Message: message, Data: data}
}

// nitaStatus est l'interface implicitement satisfaite par CHACUNE des enveloppes
// générées (ApisResponse, ApisResponseV2Double, ApisResponseV2String,
// ApisResponseV2ListVilleResponse, ApisResponseV2IndicatifsResponseV2,
// ApisResponseV2FraisPreviewResponseV2, ApisResponseV2CheckStatusResponseV2,
// ApisResponseV2ModelEditEnvoieRequest, ApisResponseV2FraisPreviewResponseV2…) : seul le type
// de leur champ Data varie d'une enveloppe à l'autre, donc GetStatus/GetCode/
// GetMessage suffisent à trancher succès/échec indépendamment du type générique T
// utilisé par Unwrap.
//
// Chaque méthode générée gère déjà un récepteur nil (o == nil renvoie la valeur
// zéro), donc passer un pointeur nil à Unwrap est sans danger.
type nitaStatus interface {
	GetStatus() string
	GetCode() int32
	GetMessage() string
}

// Unwrap est l'idiome (data, error) de la couche mince Go, à appeler juste après un
// Execute() du SDK généré :
//
//	svc, ctx, err := nc.Transactions() // nc : *NitaClient — voir nita_client.go
//	resp, _, err := svc.PartenaireToCash(ctx).PartenaireToCashDtoV2(dto).Execute()
//	data, err := nita.Unwrap(resp, resp.GetData(), err)
//	// data : map[string]interface{} (type de ApisResponse.Data), typé automatiquement.
//
// Comportement :
//   - si err != nil (erreur de transport/désérialisation d'Execute), renvoyée telle
//     quelle (zero value, err) ;
//   - sinon si resp est nil, ou que son status n'est pas "success" — comparaison
//     INSENSIBLE À LA CASSE ("success"/"Success"/"SUCCESS" selon l'endpoint) —
//     renvoie une *NitaError construite depuis l'enveloppe (Status/Code/Message) et
//     portant data en pièce jointe (NitaError.Data) ;
//   - sinon renvoie (data, nil).
//
// T est inféré depuis l'argument positionnel `data` (jamais depuis `resp`, qui n'est
// typé que par l'interface non générique nitaStatus) : c'est une inférence de
// paramètre de fonction tout à fait ordinaire, qui fonctionne donc identiquement
// pour chacune des enveloppes générées, y compris sur le Go 1.18 fixé par go.mod.
func Unwrap[T any](resp nitaStatus, data T, err error) (T, error) {
	var zero T
	if err != nil {
		return zero, err
	}
	if resp == nil {
		return zero, newNitaError("", 0, "", nil)
	}
	if strings.EqualFold(resp.GetStatus(), "success") {
		return data, nil
	}
	return zero, newNitaError(resp.GetStatus(), resp.GetCode(), resp.GetMessage(), data)
}
