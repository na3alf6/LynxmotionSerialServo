package adapters

import (

	"github.com/na3alf6/LynxmotionSerialServo/src/api/usecases"
)

type ServoController struct {
	Interactor usecases.ServoInteractor
}

func NewServoController(handler servoHandler) *ServoController {
	return &ServoController{
		Interactor: usecases.ServoInteractor{
			ServoRepository: &usecases.ServoRepository{
				ServoHandler: handler,
			},
		},
	}
}

func (controller *ServoController) Create(c Context) {
	c.Bind(&u)
	Servo, err := controller.Interactor.Add(u)
	if err != nil {
		c.JSON(500, NewError(err))
		return
	}
	c.JSON(200, Servo)
}
