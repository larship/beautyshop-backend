package models

import (
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/larship/beautyshop/database"
	"time"
)

type Beautyshop struct {
	Uuid        string       `json:"uuid"`
	Name        string       `json:"name"`
	City        string       `json:"city"`
	Address     string       `json:"address"`
	Coordinates pgtype.Point `json:"coordinates"`
	OpenHour    uint16       `json:"openHour"`
	CloseHour   uint16       `json:"closeHour"`
	CreatedDate time.Time    `json:"createdDate"`
	Workers     []*Worker    `json:"workers"`
}

func GetBeautyshopByUuid(beautyshopUuid string) *Beautyshop {
	var beautyshop Beautyshop

	sql := `
		SELECT *
		FROM beautyshops
		WHERE uuid = $1
	`

	err := database.DB.GetConnection().QueryRow(context.Background(), sql, beautyshopUuid).Scan(&beautyshop.Uuid,
		&beautyshop.Name, &beautyshop.City, &beautyshop.Address, &beautyshop.Coordinates, &beautyshop.OpenHour,
		&beautyshop.CloseHour, &beautyshop.CreatedDate)
	if err != nil {
		fmt.Printf("Ошибка получения салона красоты: %v", err)
		return nil
	}

	beautyshop.Workers = GetWorkers(beautyshop.Uuid)

	return &beautyshop
}

func GetBeautyshops(city string) []*Beautyshop {
	sql := `
		SELECT *
		FROM beautyshops
		WHERE city = $1
		ORDER BY created_date ASC
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, city)
	if err != nil {
		fmt.Printf("Ошибка получения салонов красоты: %v", err)
		return nil
	}

	var beautyshopsList []*Beautyshop
	for rows.Next() {
		var beautyshopItem Beautyshop
		err = rows.Scan(&beautyshopItem.Uuid, &beautyshopItem.Name, &beautyshopItem.City,
			&beautyshopItem.Address, &beautyshopItem.Coordinates, &beautyshopItem.OpenHour, &beautyshopItem.CloseHour,
			&beautyshopItem.CreatedDate)

		beautyshopsList = append(beautyshopsList, &beautyshopItem)
	}

	// @TODO Запрашивать сотрудников пачкой
	for _, val := range beautyshopsList {
		val.Workers = GetWorkers(val.Uuid)
	}

	return beautyshopsList
}
