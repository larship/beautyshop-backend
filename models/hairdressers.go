package models

import (
	"context"
	"fmt"
	"github.com/google/uuid"
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

func AddHairdresser(barbershopUuid string, hairdresserFullName string) bool {
	var sql string
	var err error
	barbershop := GetBarbershopByUuid(barbershopUuid)

	if barbershop == nil {
		fmt.Printf("При добавлении парикмахера не смогли найти парикмахерскую с uuid = %s", barbershopUuid)
		return false
	}

	hairdresserUuid := uuid.New().String()

	sql = `
		INSERT INTO hairdressers
		VALUES ($1, $2)
	`

	_, err = database.DB.GetConnection().Exec(context.Background(), sql, hairdresserUuid, hairdresserFullName)

	if err != nil {
		fmt.Printf("Ошибка добавления парикмахера: %v", err)
		return false
	}

	sql = `
		INSERT INTO barbershops_hairdressers
		VALUES ($1, $2)
	`

	_, err = database.DB.GetConnection().Exec(context.Background(), sql, barbershopUuid, hairdresserUuid)

	if err != nil {
		fmt.Printf("Ошибка добавления связи парикмахерская-парикмахер: %v", err)
		return false
	}

	return true
}
