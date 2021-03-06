package aggtext

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var sentencesTests = []struct {
	input    string
	expected []string
}{
	{
		`Lorem Ipsum is simply dummy text of the printing and typesetting industry.
		Lorem Ipsum has been the industry's standard dummy text ever since the 1500s,
		when an unknown printer took a galley of type and scrambled it to make a type specimen book.
		It has survived not only five centuries, but also the leap into electronic typesetting,
		remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset
		sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like
		Aldus PageMaker including versions of Lorem Ipsum.`,
		[]string{
			"Lorem Ipsum is simply dummy text of the printing and typesetting industry.",
			"Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
			"It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.",
			"It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.",
		},
	},
	{
		`Contrary to popular belief, Lorem Ipsum is not simply random text.
		It has roots in a piece of classical Latin literature from 45 BC,
		making it over 2000 years old. Richard McClintock, a Latin professor at
		Hampden-Sydney College in Virginia, looked up one of the more obscure
		Latin words, consectetur, from a Lorem Ipsum passage, and going through the
		cites of the word in classical literature, discovered the undoubtable source.
		Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum"
		(The Extremes of Good and Evil) by Cicero, written in 45 BC. This book is a treatise on the
		theory of ethics, very popular during the Renaissance!
		The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.`,
		[]string{
			"Contrary to popular belief, Lorem Ipsum is not simply random text.",
			"It has roots in a piece of classical Latin literature from 45 BC, making it over 2000 years old.",
			"Richard McClintock, a Latin professor at Hampden-Sydney College in Virginia, looked up one of the more obscure Latin words, consectetur, from a Lorem Ipsum passage, and going through the cites of the word in classical literature, discovered the undoubtable source.",
			`Lorem Ipsum comes from sections 1.10.32 and 1.10.33 of "de Finibus Bonorum et Malorum" (The Extremes of Good and Evil) by Cicero, written in 45 BC.`,
			"This book is a treatise on the theory of ethics, very popular during the Renaissance!",
			`The first line of Lorem Ipsum, "Lorem ipsum dolor sit amet..", comes from a line in section 1.10.32.`,
		},
	},
	{
		`			Lorem ipsum - tai fiktyvus tekstas naudojamas spaudos ir grafinio dizaino pasaulyje jau nuo XVI a. pradžios. Lorem Ipsum tapo
		standartiniu fiktyviu tekstu, kai nežinomas spaustuvininkas atsitiktine tvarka išdėliojo raides atspaudų prese ir tokiu
		būdu sukūrė    raidžių egzempliorių.
		Šis tekstas išliko beveik nepasikeitęs ne tik penkis amžius, bet ir įžengė i kopiuterinio
		grafinio dizaino laikus? Jis išpopuliarėjo XX a. šeštajame dešimtmetyje, kai buvo išleisti Letraset
		lapai su Lorem Ipsum ištraukomis, o vėliau -leidybinė sistema AldusPageMaker, kurioje buvo ir Lorem Ipsum versija.`,
		[]string{
			" Lorem ipsum - tai fiktyvus tekstas naudojamas spaudos ir grafinio dizaino pasaulyje jau nuo XVI a. pradžios.",
			"Lorem Ipsum tapo standartiniu fiktyviu tekstu, kai nežinomas spaustuvininkas atsitiktine tvarka išdėliojo raides atspaudų prese ir tokiu būdu sukūrė raidžių egzempliorių.",
			"Šis tekstas išliko beveik nepasikeitęs ne tik penkis amžius, bet ir įžengė i kopiuterinio grafinio dizaino laikus?",
			"Jis išpopuliarėjo XX a. šeštajame dešimtmetyje, kai buvo išleisti Letraset lapai su Lorem Ipsum ištraukomis, o vėliau -leidybinė sistema AldusPageMaker, kurioje buvo ir Lorem Ipsum versija.",
		},
	},
	{
		"XVI a. pradžios. Geras sakinys. Nuo senu laiku!",
		[]string{
			"XVI a. pradžios.",
			"Geras sakinys.",
			"Nuo senu laiku!",
		},
	},
	{"Traling... dots.. are Ignored?",
		[]string{
			"Traling... dots.. are Ignored?",
		},
	},
	{
		"Traling... Uppercase Dots.. Not",
		[]string{
			"Traling...",
			"Uppercase Dots..",
			"Not",
		},
	},
	{
		"Traling!!!! exclamation Marks!! Not?? This one.",
		[]string{
			"Traling!!!! exclamation Marks!!",
			"Not??",
			"This one.",
		},
	},
}

func TestSentences(t *testing.T) {
	for _, test := range sentencesTests {
		got := Sentences(test.input)

		assert.Len(t, got, len(test.expected))
		assert.Equal(t, test.expected, got)
	}
}

var isSentenceEndTests = []struct {
	input    []rune
	expected bool
}{
	{
		[]rune(" A"),
		true,
	},
	{
		[]rune(" Š"),
		true,
	},
	{
		[]rune(".Š"),
		false,
	},
	{
		[]rune(" Š"),
		true,
	},
	{
		[]rune(" !"),
		false,
	},
	{
		[]rune(""),
		false,
	},
}

func TestIsSentenceEnd(t *testing.T) {
	for _, test := range isSentenceEndTests {
		got := isSentenceEnd(test.input...)

		assert.Equal(t, test.expected, got)
	}
}
