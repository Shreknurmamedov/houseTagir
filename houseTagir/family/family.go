package family

type Family struct {
	Member string
	Name   string
}

func AddFamilyInfo() []Family {
	return []Family{
		{Member: "Мать", Name: "Сунахалум"},
		{Member: "Отец", Name: "Нарик"},
		{Member: "Старшая дочь", Name: "Эмина"},
		{Member: "Младшая дочь", Name: "Диана"},
		{Member: "Сын", Name: "Тагир"},
	}
}
