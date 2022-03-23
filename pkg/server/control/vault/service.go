package vault

import (
	"sort"

	"github.com/NexClipper/sudory/pkg/server/database"
	"github.com/NexClipper/sudory/pkg/server/database/prepare"
	"github.com/NexClipper/sudory/pkg/server/macro/nullable"
	servicev1 "github.com/NexClipper/sudory/pkg/server/model/service/v1"
	stepv1 "github.com/NexClipper/sudory/pkg/server/model/service_step/v1"
	"github.com/pkg/errors"
)

//Service
type Service struct {
	ctx database.Context
}

func NewService(ctx database.Context) *Service {
	return &Service{ctx: ctx}
}

func (vault Service) Create(model servicev1.ServiceAndSteps) (*servicev1.DbSchemaServiceAndSteps, error) {
	//create service
	record_service := &servicev1.DbSchema{Service: model.Service}
	if err := vault.ctx.Create(record_service); err != nil {
		return nil, errors.Wrapf(err, "database create")
	}
	//create steps
	record_steps := make([]stepv1.DbSchema, len(model.Steps))
	for i := range model.Steps {
		record := &stepv1.DbSchema{ServiceStep: model.Steps[i]}
		if err := vault.ctx.Create(record); err != nil {
			return nil, errors.Wrapf(err, "database create")
		}
		record_steps[i] = *record
	}

	return &servicev1.DbSchemaServiceAndSteps{DbSchema: *record_service, Steps: record_steps}, nil
}

func (vault Service) Get(uuid string) (*servicev1.DbSchemaServiceAndSteps, error) {
	//get service
	where := "uuid = ?"
	args := []interface{}{
		uuid,
	}
	service := &servicev1.DbSchema{}
	if err := vault.ctx.Where(where, args...).Get(service); err != nil {
		return nil, errors.Wrapf(err, "database get where=%s args=%+v", where, args)
	}

	//find step
	where = "service_uuid = ?"
	args = []interface{}{
		service.Uuid,
	}
	steps := make([]stepv1.DbSchema, 0)
	if err := vault.ctx.Where(where, args...).Find(&steps); err != nil {
		return nil, errors.Wrapf(err, "database find where=%s args=%+v", where, args)
	}
	//sort -> Sequence ASC
	sort.Slice(steps, func(i, j int) bool {
		return nullable.Int32(steps[i].Sequence).Value() < nullable.Int32(steps[j].Sequence).Value()
	})

	return &servicev1.DbSchemaServiceAndSteps{DbSchema: *service, Steps: steps}, nil
}

func (vault Service) Find(where string, args ...interface{}) ([]servicev1.DbSchemaServiceAndSteps, error) {
	//find service
	records := make([]servicev1.DbSchema, 0)
	if err := vault.ctx.Where(where, args...).Find(&records); err != nil {
		return nil, errors.Wrapf(err, "database find where=%s args=%+v", where, args)
	}
	//make result
	var services = make([]servicev1.DbSchemaServiceAndSteps, len(records))
	for i := range records {
		service := records[i]
		//set service
		services[i].DbSchema = service
		//find step
		where := "service_uuid = ?"
		args := []interface{}{
			service.Uuid,
		}
		steps := make([]stepv1.DbSchema, 0)
		if err := vault.ctx.Where(where, args...).Find(&steps); err != nil {
			return nil, errors.Wrapf(err, "database find where=%s args=%+v", where, args)
		}
		//sort -> Sequence ASC
		sort.Slice(steps, func(i, j int) bool {
			return nullable.Int32(steps[i].Sequence).Value() < nullable.Int32(steps[j].Sequence).Value()
		})

		//set steps
		services[i].Steps = steps
	}

	return services, nil
}

func (vault Service) Query(query map[string]string) ([]servicev1.DbSchemaServiceAndSteps, error) {
	//parse query
	preparer, err := prepare.NewParser(query)
	if err != nil {
		return nil, errors.Wrapf(err, "prepare newParser query=%+v", query)
	}

	//find service
	records := make([]servicev1.DbSchema, 0)
	if err := vault.ctx.Prepared(preparer).Find(&records); err != nil {
		return nil, errors.Wrapf(err, "database find query=%+v", query)
	}

	//make result
	var services = make([]servicev1.DbSchemaServiceAndSteps, len(records))
	for i := range records {
		service := records[i]
		//set service
		services[i].DbSchema = service
		//find step
		where := "service_uuid = ?"
		args := []interface{}{
			service.Uuid,
		}
		steps := make([]stepv1.DbSchema, 0)
		if err := vault.ctx.Where(where, args...).Find(&steps); err != nil {
			return nil, errors.Wrapf(err, "database find where=%s args=%+v", where, args)
		}
		//sort -> Sequence ASC
		sort.Slice(steps, func(i, j int) bool {
			return nullable.Int32(steps[i].Sequence).Value() < nullable.Int32(steps[j].Sequence).Value()
		})

		//set steps
		services[i].Steps = steps
	}

	return services, nil
}

func (vault Service) Update(model servicev1.Service) (*servicev1.DbSchema, error) {
	where := "uuid = ?"
	args := []interface{}{
		model.Uuid,
	}
	record := &servicev1.DbSchema{Service: model}
	if err := vault.ctx.Where(where, args...).Update(record); err != nil {
		return nil, errors.Wrapf(err, "database update where=%s args=%+v", where, args)
	}

	return record, nil
}

func (vault Service) Delete(uuid string) error {

	//delete step
	where := "service_uuid = ?"
	args := []interface{}{
		uuid,
	}
	step := &stepv1.DbSchema{}
	if err := vault.ctx.Where(where, args...).Delete(step); err != nil {
		return errors.Wrapf(err, "database delete where=%s args=%+v", where, args)
	}
	//delete service
	where = "uuid = ?"
	args = []interface{}{
		uuid,
	}
	service := &servicev1.DbSchema{}
	if err := vault.ctx.Where(where, args...).Delete(service); err != nil {
		return errors.Wrapf(err, "database delete where=%s args=%+v", where, args)
	}
	return nil
}