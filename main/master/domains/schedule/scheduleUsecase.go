package schedule

import (
	"database/sql"
	"playUp/main/master/model"
	"sync"
)

// UserUsecase struct
type ScheduleUsecase struct {
	db           *sql.DB
	ScheduleRepo ScheduleRepository
}

type ScheduleUsecaseInterface interface {
	createScheduleBadmintonS(ns *schedule, res *model.Result, wg *sync.WaitGroup)
	updateScheduleBadmintonS(newSchedule *schedule, finalResult *resultMatch, wg *sync.WaitGroup, result *model.Result)

	getScheduleBadmintonS(scheduleList *[]*schedule, id, status string) error
	getScheduleDetailBadmintonS(res *model.Result, id *string)
	getScheduleBadmintonSALL(res *model.Result, page, limit *string)
}

func NewScheduleUsecase(db *sql.DB) ScheduleUsecaseInterface {
	return ScheduleUsecase{db, NewScheduleRepo(db)}
}
func (r ScheduleUsecase) createScheduleBadmintonS(ns *schedule, res *model.Result, wg *sync.WaitGroup) {
	defer wg.Done()

	err := r.ScheduleRepo.createScheduleBadmintonS(ns, res)
	if err != nil {

		res.Err = err
	}

}
func (r ScheduleUsecase) updateScheduleBadmintonS(newSchedule *schedule, finalResult *resultMatch, wg *sync.WaitGroup, result *model.Result) {
	defer wg.Done()
	err := r.ScheduleRepo.updateScheduleBadmintonS(newSchedule, finalResult)
	if err != nil {

		result.Err = err
	}

}
func (r ScheduleUsecase) getScheduleBadmintonS(scheduleList *[]*schedule, id, status string) error {
	err := r.ScheduleRepo.getScheduleBadmintonS(scheduleList, id, status)
	if err != nil {

		return err
	}
	return nil
}
func (r ScheduleUsecase) getScheduleDetailBadmintonS(res *model.Result, id *string) {

	err := r.ScheduleRepo.getScheduleDetailBadmintonS(res, id)
	if err != nil {
		res.Err = err
	}

}
func (r ScheduleUsecase) getScheduleBadmintonSALL(res *model.Result, page, limit *string) {
	err := r.ScheduleRepo.getScheduleBadmintonSALL(res, page, limit)
	if err != nil {
		res.Err = err
	}

}
