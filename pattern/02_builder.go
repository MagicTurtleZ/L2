package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
	Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
*/

// продукт — создаваемый разными строителями объект
type IDEConfig struct {
	Editor       string
	Language     string
	Theme        string
	FontSize     int
	ShowLineNums bool
}

// конкретные строители реализуют строительные шаги, каждый по-своему.
type DefaultSettings struct {
	IDEConfig
}

func (ds *DefaultSettings) SetEditor() {
	ds.Editor = "VSCode"
}

func (ds *DefaultSettings) SetLanguage() {
	ds.Language = "English"
}

func (ds *DefaultSettings) SetTheme() {
	ds.Theme = "Light"
}

func (ds *DefaultSettings) SetFontSize() {
	ds.FontSize = 16
}

func (ds *DefaultSettings) SetShowLineNums() {
	ds.ShowLineNums = true
}

func (ds *DefaultSettings) GetConfig() IDEConfig {
	return ds.IDEConfig
}

type CustomSettings struct {
	IDEConfig
}

func (cs *CustomSettings) SetEditor() {
	cs.Editor = "GoLand"
}

func (cs *CustomSettings) SetLanguage() {
	cs.Language = "Russian"
}

func (cs *CustomSettings) SetTheme() {
	cs.Theme = "Dark"
}

func (cs *CustomSettings) SetFontSize() {
	cs.FontSize = 14
}

func (cs *CustomSettings) SetShowLineNums() {
	cs.ShowLineNums = false
}

func (cs *CustomSettings) GetConfig() IDEConfig {
	return cs.IDEConfig
}

// интерфейс строителя объявляет шаги конструирования продуктов, общие для всех видов строителей.
type IDEConfigBuilder interface {
	SetEditor()
	SetLanguage()
	SetTheme()
	SetFontSize()
	SetShowLineNums()
	GetConfig() IDEConfig
}

// директор определяет порядок вызова строительных шагов для производства той или иной конфигурации продуктов
type IDEDirector struct {
	IDEBuilder IDEConfigBuilder
}

func (d *IDEDirector) Construct() IDEConfig {
	d.IDEBuilder.SetEditor()
	d.IDEBuilder.SetLanguage()
	d.IDEBuilder.SetTheme()
	d.IDEBuilder.SetFontSize()
	d.IDEBuilder.SetShowLineNums()
	return d.IDEBuilder.GetConfig()
}

// клиентский код
func ExampleBuilder() {
	def := new(DefaultSettings)
	director := IDEDirector{
		IDEBuilder: def,
	}
	fmt.Println(director.Construct())

	cust := new(CustomSettings)
	director = IDEDirector{
		IDEBuilder: cust,
	}
	fmt.Println(director.Construct())

}

/*
	Плюсы:
1. позволяет создавать продукты пошагово
2. позволяет использовать один и тот же код для создания различных продуктов
3. изолирует сложный код сборки продукта от его основной бизнес-логики
	Минусы:
1. усложняет код программы из-за введения дополнительных классов/структур
2. клиент будет привязан к конкретному объекту строителя, т.к. интерфейс может не содержать
   нужного метода, тогда будет необходимо вносить в него правки
*/
