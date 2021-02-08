package models

import (
	"context"
	"fmt"
	"github.com/larship/beautyshop/database"
)

type Beautyshop struct {
	Uuid    string   `json:"uuid"`
	Name    string   `json:"name"`
	City    string   `json:"city"`
	Address string   `json:"address"`
	Workers []Worker `json:"workers"`
}

func GetBeautyshopByUuid(beautyshopUuid string) *Beautyshop {
	var beautyshop Beautyshop

	sql := `
		SELECT *
		FROM beautyshops
		WHERE uuid = $1
	`

	err := database.DB.GetConnection().QueryRow(context.Background(), sql, beautyshopUuid).Scan(&beautyshop.Uuid, &beautyshop.Name, &beautyshop.City, &beautyshop.Address)
	if err != nil {
		fmt.Printf("Ошибка получения салона красоты: %v", err)
		return nil
	}

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
		err = rows.Scan(&beautyshopItem.Uuid, &beautyshopItem.Name, &beautyshopItem.City, &beautyshopItem.Address)
		beautyshopsMap[beautyshopItem.Uuid] = &beautyshopItem
	}

	// Получим типы услуг, которые делают в салоне красоты (выборка по мастерам)
	// sql = `
	// 	SELECT bst.beautyshop_uuid, st.*
	// 	FROM beautyshops_service_types bst
	// 	INNER JOIN service_types st ON st.uuid = bst.service_type_uuid
	// `
	// rows, err = database.DB.GetConnection().Query(context.Background(), sql)
	// if err != nil {
	// 	fmt.Printf("Ошибка получения типов стрижек для парикмахерских: %v", err)
	// 	return nil
	// }
	//
	// for rows.Next() {
	// 	var beautyshopUuid string
	// 	var serviceTypeItem ServiceType
	// 	err = rows.Scan(&beautyshopUuid, &serviceTypeItem.Uuid, &serviceTypeItem.Name)
	//
	// 	if val, ok := beautyshopsMap[beautyshopUuid]; ok {
	// 		val.ServiceTypes = append(val.ServiceTypes, serviceTypeItem)
	// 	}
	// }

	// Получим мастеров салона красоты
	sql = `
		SELECT bw.beautyshop_uuid, w.*
		FROM beautyshops_workers bw
		INNER JOIN workers w ON w.uuid = bw.worker_uuid
	`

	// TODO Тут бы в самом запросе отфильтроваться по списку UUID, чтобы не тащить всех мастеров

	rows, err = database.DB.GetConnection().Query(context.Background(), sql)
	if err != nil {
		fmt.Printf("Ошибка получения мастеров салонов красоты: %v", err)
		return nil
	}

	for rows.Next() {
		var beautyshopUuid string
		var workerItem Worker
		err = rows.Scan(&beautyshopUuid, &workerItem.Uuid, &workerItem.FullName, &workerItem.Description)

		if val, ok := beautyshopsMap[beautyshopUuid]; ok {
			val.Workers = append(val.Workers, workerItem)
		}
	}

	var beautyshopsList []Beautyshop
	for _, val := range beautyshopsMap {
		beautyshopsList = append(beautyshopsList, *val)
	}

	return beautyshopsList
}
