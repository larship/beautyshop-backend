package models

import (
	"context"
	"fmt"
	"github.com/larship/barbershop/database"
	"time"
)

type ScheduleItem struct {
	Uuid string
	Barbershop     Barbershop
	ClientUuid string // @TODO Подгружать полностью клиента
	Hairdresser Hairdresser
	HaircutType HaircutType
	StartDate time.Time
	EndDate time.Time
}

func GetScheduleItems(barbershopUuid string, from string, to string) []ScheduleItem {
	sql := `
		SELECT
			s.uuid, s.start_date, s.end_date,
			b.name, b.city, b.address,
			h.full_name,
			ht.name
		FROM schedule s
		INNER JOIN barbershops b ON b.uuid = s.barbershop_uuid
		INNER JOIN hairdressers h ON h.uuid = s.hairdresser_uuid
		INNER JOIN haircut_types ht ON ht.uuid = s.haircut_type_uuid
		WHERE s.barbershop_uuid = $1
	`

	rows, err := database.DB.GetConnection().Query(context.Background(), sql, barbershopUuid)
	if err != nil {
		fmt.Printf("Ошибка получения элементов расписания стрижек: %v", err)
		return nil
	}

	var schedule []ScheduleItem

	for rows.Next() {
		var item ScheduleItem
		// var barbershop Barbershop
		// var hairdresser Hairdresser
		// var haircutType HaircutType
		// @TODO Получать остальные структуры
		var tmp string

		err = rows.Scan(&item.Uuid, &item.StartDate, &item.EndDate, &tmp, &tmp, &tmp, &tmp, &tmp)
		if err != nil {
			fmt.Printf("Ошибка получения данных элементов расписания стрижек: %v", err)
		}

		schedule = append(schedule, item)
	}

	return schedule
}
