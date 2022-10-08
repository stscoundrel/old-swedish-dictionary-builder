package dictionary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDictionariesContainExpectedAmountOfEntries(t *testing.T) {
	combined, first, latter := GetOldSwedishDictionary()

	assert.Equal(t, 22572, len(first), "Incorrect amount of entries for first volumes!")
	assert.Equal(t, 19172, len(latter), "Incorrect amount of entries for latter volumes!")
	assert.Equal(t, 41744, len(combined), "Incorrect amount of entries for combined volumes!")
	assert.Equal(t, len(first)+len(latter), len(combined), "Combined should consist of first & latter volumes!")
}

func TestDictionariesContainExpectedContent(t *testing.T) {
	dictionary, _, _ := GetOldSwedishDictionary()

	expected1 := DictionaryEntry{
		Headword:          "af bränna",
		PartOfSpeech:      "vb",
		GrammaticalAspect: "v.",
		Definitions: []string{
			"afbränna, genom eld förstöra. hans trähws the af brendhe  RK 2: 2757 . ib 1511. halff stadhen är affbrändh  BSH 5: 132 (  1506) . Jfr bränna af.",
		},
		AlternativeForms: []string{},
	}

	expected2 := DictionaryEntry{
		Headword:          "alder daghar",
		PartOfSpeech:      "nn",
		GrammaticalAspect: "pl.",
		Definitions: []string{
			"ålderdom.  &quot; tiill aller da[gha] &quot; MD (S) 242 . oppa sina aldher dagha  Lg 3: 650 .",
		},
		AlternativeForms: []string{
			"aller- )",
		},
	}

	expected3 := DictionaryEntry{
		Headword:          "fiädhrdher",
		PartOfSpeech:      "av",
		GrammaticalAspect: "adj.",
		Definitions: []string{
			" Jfr ofiädhradher.",
		},
		AlternativeForms: []string{},
	}

	expected4 := DictionaryEntry{
		Headword:          "vis",
		PartOfSpeech:      "av",
		GrammaticalAspect: "adj.",
		Definitions: []string{
			" L.",
			"1)  vetane, kunnig; vis, förståndig; vittnade om vishet el. förstånd. octouianus viis man ok vitar . . . lät kalla sik ena visasta spakano  Bu 62 .  &quot; johannes crisotomus vis mästare &quot; ib 49 .  &quot; iak huxaþe mina systor vara dara. oc fan iak sannare at hon är wrþen en þe visaste &quot; ib 187 .  &quot; juþa ok hälzt þe försto. ok þe visasto. läto sk þät dighart angra &quot; ib 197 .  ib 100 .  &quot; var härra . . . sannaþe han vara . . . wisastan mästara þere const &quot; ib 181 .  Bil 711 .  &quot; en wiis philosophus &quot; ib 535 .  &quot; mön war mykit wiis (j boklighe konst) &quot; ib.  &quot; mön giordhe sik wisare än mästarin &quot; ib 213 .  &quot; vise mästara . . . letadho hwru wärdlin wardh &quot; MB 1: 31 .  &quot; thw äst en wiisman oc höffuisk (vitr naðr oc listugr) &quot; Di 36 .  ib 35, 280 .  Al 157, 163, 179 .  &quot; helsar iak thik porum wiis &quot; ib 4364 . salomon wise  KS 69 (170, 76) .  &quot; aff visom läkiarom &quot; KL 362 . osniellan ok litit vndirstandande forma iak at göra ympnit sniellan ok visan  Bir 1: 255 .  &quot; man kan icke skapa menniskien stragx saa wiis och kloch, som henne bör wara &quot; FM 489 ( 1510) .  &quot; mangha waro ok ära ydher tyl handa som . . . ära wysane ok klokana. än iak är &quot; Gers Frest Inl. görs henna storligha behof wardha wysz ok klok tyl ath waktha sik ok stridha moth swa dane fiende  ib.  &quot; tyl ath göra menniskene wysa ok frunomstogha moth dyöffwlsens forsath ib. fore alla män tha ära the wiis &quot; Al 40 .  &quot; hwa är swa viis (för oviis?) at han rädhis ey sin skadha (tam quis inops mentis quod non timet a noceumentis) &quot; GO 987 ; se Kock, Medeltidsordspr. 2: 387. thz ordh sagdhe en tunga wiis mz wiso radhe oc wise akt haffwer thu thz ordhith sakt  Al 8342, 8343 .  &quot; medh wise föreakt &quot; KS 82 (203, 90) . medh . . . wise daghtingan  ib 75 (185, 82). mz blidho änlite. viso oc vyrdhelico  Bo 27 . - förståndig, som förståg si på (ngt). jak är ympnyt sniällir oc viis til radha (sapiens ad consilia)  Bir 3: 464 . the äre vise til alla ära  Fr 2737 .  RK 2: 2924 . ",
			"2)  bekant med (ngt), egande kännedom om el. erfarenhet af (ngt). vara vis, ega kännedom om, veta. med gen. the frw heet olipiadis wään ok fagher thäs war han wiis  Al 150 .  &quot; han wan idhar stadh miteradis ok brände han op thäs ärin ij wiis &quot; ib 2408 (båda dessa ställen måhända att föra till följ. ord 6). - med prep. um. alla qwinnor waro bundna vnder thetta band, vtan the waro wiisa om gudz wilia  MB 1: 134 . - vara el. komma i erfarenhet (af ngt). med gen. k cristiern hade tänkt swänskom saa harth eeth riiss gudh them naadhe som täss haffde warith wiiss  RK 3: 2401 .  &quot; vara sik vis, veta. med ack. alexander thz war thik wis thz iak haffwer thik lätith winna alla men A 8595 (måhända att föra till följ. ord 6). - vardha vis, lära känna, få veta, få besked om, erfara (ngt) ij thänna tima vardh konungin viis huilkin ära ok huilkin priis herra iwan hafdhe wunnith thär &quot; Iv 1829 . helena wilde wiis wardha hwilkit korssit wars herra ware  Bil 87 . - med gen. el. ack. hwar som aarla riis han wardhir mangs viis  GO 9 .  ib 489 .  &quot; miin fru thäs viis nu vordhin är at ij swa lönlika hafuin varith här &quot; Iv 1354 .  Al 7616 .  &quot; hans nampn vil iak nu vardha viis &quot; Iv 1266 . - erfara, komma i erfarenhet af, röna, upplefva. med gen. el. ack. tha wart han thz äuintyr wiis  Iv (Cod. B) 691 .  &quot; lithet äpter warth han thz wiiss &quot; RK 3: 1519 .  &quot; wy sculom sigher wynna och saa priis saa wardher han (Tord Bonde) mz k kalr täss wiiss ib 32. - vardha manz vis, träda i sexuell förbindels med en man, hafva köttsligt umgånge med en man. justina wilde ey wardha mans viis &quot; Bil 447 .  &quot; en nunna som aldre war manz viis &quot; ib 686 . - Jfr bok-, iäm-, lagh-, o-, rät-, skiäl-, vrang-vis.",
		},
		AlternativeForms: []string{
			"viis . ",
			"wiiss . ",
			"viis . ",
			"wiss RK 2: 2924 (i rimsl. med pris.) wysz  Gers Frest Inl. ack. m. visan. f. wysa  Gers Frest Inl. dat. f. wise  KS 75 (185, 82), 82 (203, 90) ;  Al 8343 . ",
			"viso . ",
			"wisir MB 1: 573,  &quot; hvarest denna form uppgifves tillhöra Cod. B. vise &quot; ib 31 ;  Fr 2737 . ",
			"viisa MB 1: 134 . ",
			"wyse Gers Frest Ind. dat. visom. best. form nom. m. wise  MB 1: 7 ;  KS 69 (170, 76) . ack. m. wisa  ST 413 . ",
			"wiso ib 415 . ",
			"visare . ",
			"wysane Gers Frest Inl. superl. visaster), ",
		},
	}

	assert.Equal(t, expected1, dictionary[100])
	assert.Equal(t, expected2, dictionary[666])
	assert.Equal(t, expected3, dictionary[6666])
	assert.Equal(t, expected4, dictionary[39666])
}
