package bd

import (
	"context"

	"github.com/OnlyAdal/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCn.Database(DatabaseName)
	col := db.Collection("usuarios")

	u.Password, _ = EncriptarPAssword(u.Password)
	result, err := col.InsertOne(ctx, u)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}
