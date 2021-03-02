package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/larship/beautyshop/database"
)

type Beautyshop struct {
	Uuid        string       `json:"uuid"`
	Name        string       `json:"name"`
	City        string       `json:"city"`
	Address     string       `json:"address"`
	Coordinates pgtype.Point `json:"coordinates"`
	Workers     []Worker     `json:"workers"`
}

func GetBeautyshopByUuid(beautyshopUuid string) *Beautyshop {
	var beautyshop Beautyshop

	sql := `
		SELECT *
		FROM beautyshops
		WHERE uuid = $1
	`

	err := database.DB.GetConnection().QueryRow(context.Background(), sql, beautyshopUuid).Scan(&beautyshop.Uuid,
		&beautyshop.Name, &beautyshop.City, &beautyshop.Address, &beautyshop.Coordinates)
	if err != nil {
		fmt.Printf("Ошибка получения салона красоты: %v", err)
		return nil
	}

	beautyshop.Workers = GetWorkers(beautyshop.Uuid)

	return &beautyshop
}

func GetBeautyshops(city string) []Beautyshop {
	sql := `
		SELECT *
		FROM beautyshops
		WHERE city = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, city)
	if err != nil {
		fmt.Printf("Ошибка получения салонов красоты: %v", err)
		return nil
	}

	beautyshopsMap := map[string]*Beautyshop{}
	for rows.Next() {
		var beautyshopItem Beautyshop
		err = rows.Scan(&beautyshopItem.Uuid, &beautyshopItem.Name, &beautyshopItem.City,
			&beautyshopItem.Address, &beautyshopItem.Coordinates)
		beautyshopsMap[beautyshopItem.Uuid] = &beautyshopItem
	}

	for _, val := range beautyshopsMap {
		beautyshopsMap[val.Uuid].Workers = GetWorkers(val.Uuid)
	}

	var beautyshopsList []Beautyshop
	for _, val := range beautyshopsMap {
		beautyshopsList = append(beautyshopsList, *val)
	}

	return beautyshopsList
}
