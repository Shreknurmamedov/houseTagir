package relatives

type Relatives struct {
	Member string
	Name   string
}

func AddRelativesInfo() []Relatives {
	return []Relatives{
		{Member: "Бабушка", Name: "Раижат"},
		{Member: "Дедушка", Name: "Ягибег"},
		{Member: "Дядя", Name: "Муххамад"},
		{Member: "Тетя", Name: "Глафира"},
		{Member: "Прабабушка", Name: "Эйлихан"},
	}
}
