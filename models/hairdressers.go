package models

import (
	"context"
	"fmt"
	"github.com/larship/barbershop/database"
)

type Hairdresser struct {
	Uuid     string `json:"id"`
	FullName string `json:"full-name"`
}

func GetHairdressers(barbershopUuid string) []HaircutType {
	sql := `
		SELECT h.*
		FROM hairdressers h
		INNER JOIN barbershops_hairdressers bh ON bh.hairdresser_uuid = h.uuid
		WHERE bh.barbershop_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, barbershopUuid)
	if err != nil {
		fmt.Printf("Ошибка получения парикмахеров: %v", err)
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
