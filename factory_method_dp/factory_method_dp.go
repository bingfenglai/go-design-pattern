package factory_method_dp

type Factory interface {
	CreateOperation() Operation
}

type OperationAddFactory struct {
}

func (o *OperationAddFactory) CreateOperation() Operation {
	return &OperationAdd{}
}

type OperationSubFactory struct {
}

func (o *OperationSubFactory) CreateOperation() Operation {
	return &OperationSub{}
}

type OperationMulFactory struct {
}

func (o *OperationMulFactory) CreateOperation() Operation {

	return &OperationMul{}
}

type OperationDivFactory struct {
}

func (o *OperationDivFactory) CreateOperation() Operation {
	return &OperationDiv{}
}
