package zhuk

type Type int

const (
	Simple       Type = iota
	Ass               // Жопный
	BlackRapper       // Чёрный репер
	KPopper           // Кейпопер
	GameDesigner      // Игровой дизайнер
	Hercules          // Силач
	Fisher            // Рыбак
	Deutsch           // Немец
	Alcoholic         // Алкоголик
	FilmCritic        // Кинокритик
	CProgrammer       // С-программист
	Cooker            // Повар
	Billionaire       // Миллиардер
	Banker            // Банкир
	Lover             // Любовник
	Collector         // Коллектор
	Homeless          // Бездомный
	Butler            // Дворецкий
	Streamer          // Стример
	Courier           // Курьер
	Farmer            // Фермер
)

type Zhuk struct {
	Type Type
	Name string
}
