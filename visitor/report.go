package visitor

type Report struct {
	employees []Employee
}

func NewReport() *Report {
	employees := [...]Employee{
		NewProgrammer("韩立"),
		NewProgrammer("南宫屏"),
		NewProgrammer("光脚大汉"),
		NewProgrammer("时光殿主"),
		NewProgrammer("南宫碗"),
		NewManager("紫菱仙子"),

		NewManager("陈师姐"),
	}
	return &Report{employees: employees[:]}
}

func (r *Report) Visit(visitor Visitor) {
	for _, employee := range r.employees {
		employee.Accept(visitor)
	}
}
