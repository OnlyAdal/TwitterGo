package bd

import (
	"context"

	"github.com/OnlyAdal/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx := context.TODO()

	db := MongoCn.Database(DatabaseName)
	col := db.Collection("usuarios")

	condition := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condition).Decode(&resultado)

	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
