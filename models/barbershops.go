package models

import (
	"context"
	"fmt"
	"github.com/larship/barbershop/database"
)

type Barbershop struct {
	Uuid         string        `json:"uuid"`
	Name         string        `json:"name"`
	City         string        `json:"city"`
	Address      string        `json:"address"`
	HaircutTypes []HaircutType `json:"haircut-types"`
	Hairdressers []Hairdresser `json:"hairdressers"`
}

func GetBarbershopByUuid(barbershopUuid string) *Barbershop {
	var barbershop Barbershop

	sql := `
		SELECT *
		FROM barbershops
		WHERE uuid = $1
	`

	err := database.DB.GetConnection().QueryRow(context.Background(), sql, barbershopUuid).Scan(&barbershop.Uuid, &barbershop.Name, &barbershop.City, &barbershop.Address)
	if err != nil {
		fmt.Printf("Ошибка получения парикмахерской: %v", err)
		return nil
	}

	return &barbershop
}

func GetBarbershops(city string) []Barbershop {
	sql := `
		SELECT *
		FROM barbershops
		WHERE city = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, city)
	if err != nil {
		fmt.Printf("Ошибка получения парикхамерских: %v", err)
		return nil
	}

	barbershopsMap := map[string]*Barbershop{}
	for rows.Next() {
		var barbershopItem Barbershop
		err = rows.Scan(&barbershopItem.Uuid, &barbershopItem.Name, &barbershopItem.City, &barbershopItem.Address)
		barbershopsMap[barbershopItem.Uuid] = &barbershopItem
	}

	// Получим типы стрижек, которые делают в парикмахерской
	sql = `
		SELECT bht.barbershop_uuid, ht.*
		FROM barbershops_haircut_types bht
		INNER JOIN haircut_types ht ON ht.uuid = bht.haircut_type_uuid
	`
	rows, err = database.DB.GetConnection().Query(context.Background(), sql)
	if err != nil {
		fmt.Printf("Ошибка получения типов стрижек для парикмахерских: %v", err)
		return nil
	}

	for rows.Next() {
		var barbershopUuid string
		var haircutTypeItem HaircutType
		err = rows.Scan(&barbershopUuid, &haircutTypeItem.Uuid, &haircutTypeItem.Name)

		if val, ok := barbershopsMap[barbershopUuid]; ok {
			val.HaircutTypes = append(val.HaircutTypes, haircutTypeItem)
		}
	}

	// Получим мастеров парикмахерской
	sql = `
		SELECT bh.barbershop_uuid, h.*
		FROM barbershops_hairdressers bh
		INNER JOIN hairdressers h ON h.uuid = bh.hairdresser_uuid
	`

	// TODO Тут бы в самом запросе отфильтроваться по списку UUID, чтобы не тащить всех мастеров

	rows, err = database.DB.GetConnection().Query(context.Background(), sql)
	if err != nil {
		fmt.Printf("Ошибка получения мастеров парикмахерских: %v", err)
		return nil
	}

	for rows.Next() {
		var barbershopUuid string
		var hairdresserItem Hairdresser
		err = rows.Scan(&barbershopUuid, &hairdresserItem.Uuid, &hairdresserItem.FullName)

		if val, ok := barbershopsMap[barbershopUuid]; ok {
			val.Hairdressers = append(val.Hairdressers, hairdresserItem)
		}
	}

	var barbershopsList []Barbershop
	for _, val := range barbershopsMap {
		barbershopsList = append(barbershopsList, *val)
	}

	return barbershopsList
}
