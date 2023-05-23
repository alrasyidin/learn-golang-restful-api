package simple

type FooBarService struct {
	*FooService
	*BarService
}

// Constructor for FooBarService
func NewFooBarService(fooService *FooService, barService *BarService) *FooBarService {
	return &FooBarService{
		FooService: fooService,
		BarService: barService,
	}
}
