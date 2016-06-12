package main

import (
	"github.com/alikhajeh1/goa_app1/app"
	"github.com/goadesign/goa"
)

// NumbersController implements the numbers resource.
type NumbersController struct {
	*goa.Controller
}

// NewNumbersController creates a numbers controller.
func NewNumbersController(service *goa.Service) *NumbersController {
	return &NumbersController{Controller: service.NewController("NumbersController")}
}

func sieveOfEratosthenes(N int) (primes []int) {
    b := make([]bool, N)
    for i := 2; i < N; i++ {
        if b[i] == true { continue }
        primes = append(primes, i)
        for k := i * i; k < N; k += i {
            b[k] = true
        }
    }
    return
}

// Prime runs the prime action.
func (c *NumbersController) Show(ctx *app.ShowNumbersContext) error {
  numbers := app.GoaExampleNumbers {
      Numbers: sieveOfEratosthenes(*ctx.LessThan),
  }

	return ctx.OK(&numbers)
}
