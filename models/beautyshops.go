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
	Admins      []string     `json:"admins"`
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
	beautyshop.Admins = getBeautyshopAdminUuidList(beautyshop.Uuid)

	return &beautyshop
}

func GetBeautyshopListByAdmin(adminUuid string) []*Beautyshop {
	sql := `
		SELECT b.*
		FROM beautyshops b
		INNER JOIN beautyshops_admins ba ON ba.beautyshop_uuid = b.uuid
		WHERE ba.client_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, adminUuid)
	if err != nil {
		fmt.Printf("Ошибка получения списка администрируемых салонов красоты: %v", err)
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
	// @TODO Запрашивать администраторов пачкой
	for _, val := range beautyshopsList {
		val.Workers = GetWorkers(val.Uuid)
		val.Admins = getBeautyshopAdminUuidList(val.Uuid)
	}

	return beautyshopsList
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
	// @TODO Запрашивать администраторов пачкой
	for _, val := range beautyshopsList {
		val.Workers = GetWorkers(val.Uuid)
		val.Admins = getBeautyshopAdminUuidList(val.Uuid)
	}

	return beautyshopsList
}

func getBeautyshopAdminUuidList(beautyshopUuid string) []string {
	sql := `
		SELECT client_uuid
		FROM beautyshops_admins
		WHERE beautyshop_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, beautyshopUuid)
	if err != nil {
		fmt.Printf("Ошибка получения списка администраторов: %v", err)
		return nil
	}

	var adminUuidList []string
	for rows.Next() {
		var uuid string
		err = rows.Scan(&uuid)
		adminUuidList = append(adminUuidList, uuid)
	}

	return adminUuidList
}
