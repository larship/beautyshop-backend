package models

import (
	"context"
	"fmt"
	"github.com/larship/beautyshop/database"
	"time"
)

type ScheduleItem struct {
	Uuid        string
	Beautyshop  Beautyshop
	ClientUuid  string // @TODO Подгружать полностью клиента
	Worker      Worker
	ServiceType ServiceType
	StartDate   time.Time
	EndDate     time.Time
}

func GetScheduleItems(beautyshopUuid string, from string, to string) []ScheduleItem {
	sql := `
		SELECT
			s.uuid, s.start_date, s.end_date,
			b.name, b.city, b.address,
			w.full_name, w.description
			st.name
		FROM schedule s
		INNER JOIN beautyshops b ON b.uuid = s.beautyshop_uuid
		INNER JOIN workers w ON w.uuid = s.worker_uuid
		INNER JOIN service_types st ON st.uuid = s.service_type_uuid
		WHERE s.beautyshop_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, beautyshopUuid)
	if err != nil {
		fmt.Printf("Ошибка получения элементов расписания: %v", err)
		return nil
	}

	var schedule []ScheduleItem

	for rows.Next() {
		var item ScheduleItem
		// var beautyshop Beautyshop
		// var worker Worker
		// var serviceType ServiceType
		// @TODO Получать остальные структуры
		var tmp string

		err = rows.Scan(&item.Uuid, &item.StartDate, &item.EndDate, &tmp, &tmp, &tmp, &tmp, &tmp)
		if err != nil {
			fmt.Printf("Ошибка получения данных элементов расписания: %v", err)
		}

		schedule = append(schedule, item)
	}

	return schedule
}
