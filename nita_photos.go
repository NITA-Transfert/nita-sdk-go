// Couche mince écrite à la main (NON générée) — voir nita_client.go.
//
// Encodage Base64 des pièces d'identité (KYC) pour les endpoints d'envoi
// (PartenaireToCash), qui attendent les photos en Base64 dans le
// corps JSON. Le partenaire fournit des octets bruts, typiquement via os.ReadFile.
package nita

import "encoding/base64"

// ToBase64 encode des octets bruts en Base64 standard (sans préfixe "data:").
func ToBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// EncodePhotos encode les pièces d'identité de l'expéditeur en Base64, prêtes à
// être copiées dans un DTO d'envoi (champs photoPieceIdentiteRecto /
// photoPieceIdentiteVerso / photoIdentite). Un paramètre nil est omis de la map
// renvoyée — seules les clés correspondant à une photo fournie sont présentes.
//
//	recto, _ := os.ReadFile("cni-recto.jpg")
//	verso, _ := os.ReadFile("cni-verso.jpg")
//	photos := nita.EncodePhotos(recto, verso, nil)
//	dto.SetPhotoPieceIdentiteRecto(photos["photoPieceIdentiteRecto"])
//	dto.SetPhotoPieceIdentiteVerso(photos["photoPieceIdentiteVerso"])
func EncodePhotos(recto, verso, portrait []byte) map[string]string {
	fields := map[string]string{}
	if recto != nil {
		fields["photoPieceIdentiteRecto"] = ToBase64(recto)
	}
	if verso != nil {
		fields["photoPieceIdentiteVerso"] = ToBase64(verso)
	}
	if portrait != nil {
		fields["photoIdentite"] = ToBase64(portrait)
	}
	return fields
}
