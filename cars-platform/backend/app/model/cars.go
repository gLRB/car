package model

import (
	"gorm.io/gorm"
	"strconv"
)

type NewReq struct {
	Token     string `json:"token,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Rst       string `json:"rst,omitempty"`
	Dst       string `json:"dst,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Total     string `json:"total,omitempty"`
	CarNumber string `json:"car_number,omitempty"`
	Message   string `json:"message,omitempty"`
}

type EditReq struct {
	Token     string `json:"token,omitempty"`
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	StartDate string `json:"start_date,omitempty"`
	StartTime string `json:"start_time,omitempty"`
	EndDate   string `json:"end_date,omitempty"`
	EndTime   string `json:"end_time,omitempty"`
	Rst       string `json:"rst,omitempty"`
	Dst       string `json:"dst,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Total     string `json:"total,omitempty"`
	Remain    string `json:"remain,omitempty"`
	CarNumber string `json:"car_number,omitempty"`
	Message   string `json:"message,omitempty"`
}

type ShowReq struct {
	Rst       string `form:"rst"`
	StartDate string `form:"start_data,omitempty"`
	StartTime string `form:"start_time,omitempty"`
	TimeStamp string `form:"time_stamp,omitempty"`
	Phone     string `form:"phone"`
}

type Cars struct {
	gorm.Model
	Name      string `gorm:"comment:名称"`
	StartDate string `gorm:"comment:出发日期"`
	StartTime string `gorm:"comment:出发时间"`
	EndDate   string `gorm:"comment:到达日期"`
	EndTime   string `gorm:"comment:到达时间"`
	Rst       string `gorm:"comment:出发地"`
	Dst       string `gorm:"comment:目的地"`
	Phone     string `gorm:"comment:手机号码"`
	Total     int64  `gorm:"comment:总座位"`
	Remain    int64  `gorm:"comment:剩余座位"`
	CarNumber string `gorm:"comment:车牌号"`
	Message   string `gorm:"comment:备注"`
	State     bool   `gorm:"comment:是否被删除"`
}

func InsertCarMsg(db *gorm.DB, req NewReq) error {
	total, _ := strconv.ParseInt(req.Total, 10, 64)
	car := Cars{
		Name:      req.Name,
		StartDate: req.StartDate,
		StartTime: req.StartTime,
		EndDate:   req.EndDate,
		EndTime:   req.EndTime,
		Rst:       req.Rst,
		Dst:       req.Dst,
		Phone:     req.Phone,
		Total:     total,
		CarNumber: req.CarNumber,
		Remain:    total,
		Message:   req.Message,
		State:     true,
	}
	if err := db.Create(&car).Error; err != nil {
		return err
	}
	return nil
}

func SearchCars(db *gorm.DB, req ShowReq) ([]Cars, error) {
	var cars []Cars
	var resCars []Cars
	db.Where("phone LIKE ? and rst LIKE ? and start_date LIKE ?",
		"%"+req.Phone+"%", "%"+req.Rst+"%", "%"+req.StartDate+"%").Order("id DESC").Find(&cars)
	//resCars = cars
	//copy(resCars, cars)

	if req.StartDate != "" && req.StartTime != "" && req.TimeStamp != "" {
		for _, car := range cars {
			if verifyTime(car.StartTime, req.StartTime, req.TimeStamp) {
				resCars = append(resCars, car)
			}
		}
		return resCars, nil
	}
	return cars, nil
}

func EditCar(db *gorm.DB, req EditReq) error {
	total, _ := strconv.ParseInt(req.Total, 10, 64)
	remain, _ := strconv.ParseInt(req.Remain, 10, 64)
	car := Cars{
		Name:      req.Name,
		Rst:       req.Rst,
		Dst:       req.Dst,
		Phone:     req.Phone,
		Total:     total,
		Remain:    remain,
		CarNumber: req.CarNumber,
		Message:   req.Message,
		StartDate: req.StartDate,
		StartTime: req.StartTime,
		EndDate:   req.EndDate,
		EndTime:   req.EndTime,
	}

	db.Where("id = ?", req.ID).Updates(car)

	return nil
}

func verifyTime(source string, time string, stamp string) bool {

	sTime, _ := strconv.ParseInt(source[:2], 10, 64)
	baseTime, _ := strconv.ParseInt(time[:2], 10, 64)
	stampTime, _ := strconv.ParseInt(stamp, 10, 64)
	if sTime > baseTime+stampTime || sTime < baseTime-stampTime {
		return false
	}
	return true

}
