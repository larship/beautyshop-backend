package models

import (
	"context"
	"fmt"
	"github.com/larship/barbershop/database"
)

type HaircutType struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}

func GetHaircutTypes(barbershopUuid string) []HaircutType {
	sql := `
		SELECT ht.*
		FROM haircut_types ht
		INNER JOIN barbershops_haircut_types bht ON bht.haircut_type_uuid = ht.uuid
		WHERE bht.barbershop_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, barbershopUuid)
	if err != nil {
		fmt.Printf("Ошибка получения типов стрижек: %v", err)
		return nil
	}

	var haircutTypesList []HaircutType

	for rows.Next() {
		var item HaircutType
		err = rows.Scan(&item.Uuid, &item.Name)
		haircutTypesList = append(haircutTypesList, item)
	}

	return haircutTypesList
}
