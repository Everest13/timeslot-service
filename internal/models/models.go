package models

import "time"

/*
Считаем что сами Orders хранятся в другом сервисе, в таком кейсе чтобы не было дублирования
воспринимаем order как набор данных для бронирования слота
*/
type Order struct {
	RequestID string    `json:"request_id"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Capacity  int       `json:"capacity"`
}

/*
Reservation — это лог всех зарезервированных таймслотов. Для конечного понимания, что за запрос
зарезервировал время. Тут должны находится только актуальные данные.
Это очень важно для наших коллег из отдела аналитики!
*/
type Reservation struct {
	RequestID  string `json:"request_id"`
	TimeslotID string `json:"timeslot_id"`
	Capacity   int    `json:"capacity"`
}

/*
TimeSlot — конкретный промежуток времени и максимальное кол-во
заказов в этот промежуток. Если заказ, занимает больше кол-во
времени, чем 1 промежуток (например 3 часа на сбор и доставку) значит, нам необходимо
забронировать Capacity из 2 промежутов сразу 8:00 - 9:59 и 10:00 - 11:59.
*/
type TimeSlot struct {
	ID       string    `json:"id"`
	From     time.Time `json:"date_from"`
	To       time.Time `json:"date_to"`
	Capacity int       `json:"capacity"`
}
