package data

import "math/rand"

type Question struct {
	ID      int      `json:"id"`
	Text    string   `json:"text"`
	Answers []Answer `json:"answers"`
}

type Answer struct {
	Text  string `json:"text"`
	Value string `json:"value"`
}

type Character struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ImageURL    string `json:"image_url"`
}

// GetQuestions возвращает все вопросы теста
func GetQuestions() []Question {
	return []Question{
		{
			ID:   1,
			Text: "Что вы предпочитаете делать в свободное время?",
			Answers: []Answer{
				{Text: "Печь пироги и печенье", Value: "sovunya"},
				{Text: "Заниматься спортом и активным отдыхом", Value: "krosh"},
				{Text: "Экспериментировать и изобретать", Value: "pin"},
				{Text: "Писать стихи", Value: "barash"},
				{Text: "Читать книги и узнавать новое", Value: "losyash"},
				{Text: "Лежать на пляже и загорать", Value: "nusha"},
			},
		},
		{
			ID:   2,
			Text: "Ваша любимая погода?",
			Answers: []Answer{
				{Text: "Солнечная и ясная", Value: "krosh"},
				{Text: "Прохладная, самое то для вкусного чая", Value: "sovunya"},
				{Text: "Теплая и спокойная", Value: "nusha"},
				{Text: "Любая, главное чтобы можно было работать", Value: "pin"},
				{Text: "Дождливая и меланхоличная", Value: "barash"},
				{Text: "Безоблачное небо для изучения звезд", Value: "losyash"},
			},
		},
		{
			ID:   3,
			Text: "Какой ваш подход к решению проблем?",
			Answers: []Answer{
				{Text: "Анализировать и планировать", Value: "losyash"},
				{Text: "Действовать быстро и решительно", Value: "krosh"},
				{Text: "Найти вдохновение и потом все получится", Value: "barash"},
				{Text: "Искать нестандартные решения", Value: "pin"},
				{Text: "Главное доверять интуиции", Value: "sovunya"},
				{Text: "Проблемы не для меня, у меня всегда все в ажуре", Value: "nusha"},
			},
		},
		{
			ID:   4,
			Text: "Что для вас самое важное в жизни?",
			Answers: []Answer{
				{Text: "Знания и мудрость", Value: "losyash"},
				{Text: "Приключения и веселье", Value: "krosh"},
				{Text: "Любовь и дружба", Value: "barash"},
				{Text: "Творчество и изобретения", Value: "pin"},
				{Text: "Красота и популярность", Value: "nusha"},
				{Text: "Забота и помощь другим", Value: "sovunya"},
			},
		},
	}
}

// GetCharacters возвращает всех персонажей
func GetCharacters() map[string]Character {
	return map[string]Character{
		"krosh": {
			Name:        "Крош",
			Description: "Энергичный и непоседливый! Вы полны энтузиазма и любите приключения. Ваша жизненная энергия заряжает всех вокруг!",
			ImageURL:    "/static/images/krosh.png",
		},
		"sovunya": {
			Name:        "Совунья",
			Description: "Мудрая и заботливая! Вы всегда готовы помочь и дать хороший совет. На вас можно положиться в любой ситуации.",
			ImageURL:    "/static/images/sovunya.webp",
		},
		"pin": {
			Name:        "Пин",
			Description: "Изобретательный и творческий! Вы любите создавать что-то новое и решать сложные задачи. Ваши идеи меняют мир к лучшему!",
			ImageURL:    "/static/images/pin.png",
		},
		"barash": {
			Name:        "Бараш",
			Description: "Романтичный и мечтательный! Вы цените красоту и глубокие чувства. Ваша душа полна поэзии и вдохновения.",
			ImageURL:    "/static/images/barash.png",
		},
		"losyash": {
			Name:        "Лосяш",
			Description: "Умный и вдумчивый! Вы любите знания и стремитесь к саморазвитию. Ваша эрудиция впечатляет окружающих.",
			ImageURL:    "/static/images/losyash.png",
		},
		"nusha": {
			Name:        "Нюша",
			Description: "Красивая и жизнерадостная! Вы цените стиль и модные вещички. Вы украшение любой компании.",
			ImageURL:    "/static/images/nusha.png",
		},
	}
}

// GetRandomCharacter возвращает случайного персонажа
func GetRandomCharacter() Character {
	characters := GetCharacters()
	keys := make([]string, 0, len(characters))
	for k := range characters {
		keys = append(keys, k)
	}
	randomKey := keys[rand.Intn(len(keys))]
	return characters[randomKey]
}
