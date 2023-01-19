package traveller_test

import (
	"fmt"
	"testing"

	"github.com/pwnedgod/traveller"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type swipe struct {
	Plain   string
	Meaning string
	Peace   int
}

type facade struct {
	Party      string
	Barrel     any
	Retirement float64
	Tiger      string
	unexported string
}

type offroad struct {
	Surprise []int
	Social   []float64
	Bake     string
	Deprive  string
}

type philosophy struct {
	Deliver string
	Thirsty string
	Job     any
	Peace   map[string]string
	Tire    offroad
}

type electric struct {
	Add           float64
	Employee      []string
	Contradiction any
	Shelter       int
	Penny         []int
	Issue         []float64
	Beach         string
	Volcano       string
	Knowledge     philosophy
}

type mode struct {
	Flight int
	Fence  electric
}

type Embedded struct {
	Inheritance int
}

type open struct {
	Cave        map[string]float64
	Swipe       map[string]swipe
	Powder      string
	Mill        []int
	Carbon      any
	Consumption any
	Wall        string
	Brother     []float64
}

type speech struct {
	Enthusiasm string
	Locate     int
	Couple     []int
	Investment any
	Betray     []int
	Slippery   int
	Possession float64
	Explosion  string
	Critic     []int
	College    []string
	Policeman  any
	Traction   map[string]int
	Fax        string
	Create     mode
	Slide      open
}

type habit struct {
	Jet      facade
	Decline  map[string]string
	Clean    []int
	Elephant []string
	Hate     speech
}

type bulb struct {
	Sunshine   int
	Headache   [5]int
	Brother    [1]string
	Barrel     float64
	Worth      []string
	Cup        map[string]any
	Federation habit
	Identity   float64
	Scramble   []string
	Band       string
	Embedded
}

func makeBulb() bulb {
	return bulb{
		Sunshine: 121,
		Headache: [5]int{1021, 2930, 3718, 2848, 1366},
		Brother:  [1]string{"HFLMPL2rQcwFzlx8dw7D"},
		Barrel:   370.734,
		Worth:    []string{"9MkyQvrHMuJIQjmNgETf", "WgNTqZEG6KuSnqCocyiV", "RGnXLTPCadctoltAPnXs"},
		Cup: map[string]any{
			"Blasphemy": "ONr7QDhcZJNgiSnZByaH",
			"Favour":    "VL6foOIq436n8gevZi7K",
			"Houseplant": map[string]string{
				"Mislead":   "yDlqlodvPqwJFB5o8hKq",
				"Machinery": "nnGbiSSEYt01kotPuVHS",
			},
		},
		Federation: habit{
			Jet: facade{
				Party: "ZPGANa8QAKvR7AFzXwCn",
				Barrel: map[string]int{
					"Outside": 34,
					"Pumpkin": 420,
				},
				Retirement: 69.999,
				Tiger:      "zJwMx0qTQaYWqKsOABNf",
				unexported: "EmzONvxGZb0hc2MQK9Bu",
			},
			Decline: map[string]string{
				"Victory":  "43ZeSUgDdanbNBemUydH",
				"Instinct": "8bJD76KwNbdBMZE6L1ex",
			},
			Clean:    []int{517, 440, 168, 357, 871, 455},
			Elephant: []string{"6XlDuw5bkQiGS6o43LiL", "iJqQGu2bGu9uEv0qaCtD", "IPY0XpzqYHPG9Vpy6T0n", "t8jQLbWHG1MMjgZpFNgi"},
			Hate: speech{
				Enthusiasm: "NMEyTLqjVf6yg35qqVRh",
				Locate:     871,
				Couple:     []int{515, 133},
				Investment: "lXZDwVsLnI6FNV1e243J",
				Betray:     []int{254, 750, 614},
				Slippery:   868,
				Possession: 925.840,
				Explosion:  "iZ6bwDNLksDCUG5JU4JE",
				Critic:     []int{744, 684, 151, 243, 507},
				College:    []string{"ZgN3U1snmVcUgWNylw0A", "jSIvrxCNBok6lUwHfQgw", "z4xCxlMbQdES77U3hh3k", "a380jUF8DVIE8GNQ3E1G"},
				Policeman:  1,
				Traction: map[string]int{
					"Crosswalk": 5199,
					"Excess":    1999,
				},
				Fax: "xauUunFmmIWKt9tHhHaS",
				Create: mode{
					Flight: 740,
					Fence: electric{
						Add:           787.90,
						Employee:      []string{"wURCf15YLDnVSxeY9yf5", "JYTetPTxJo7elWO1ACqm", "NVYrrQYUzV9MXG0T6LDC", "a7SuQKccyub3yCktGkFt", "xUhmjrkjPmblnxr8JDDu"},
						Contradiction: "ETOCwPMU4Em6bscP5Hma",
						Shelter:       555,
						Penny:         []int{874, 864, 582, 807},
						Issue:         []float64{673.774, 567.674},
						Beach:         "YS6zSh6R90AlPhiEWHSu",
						Volcano:       "WMYHYn40fVjCrYNItYyK",
						Knowledge: philosophy{
							Deliver: "YOSiMHf7S9FCwCvlbnyr",
							Thirsty: "A3GPqW2maQgxzRSxluwm",
							Job: []any{
								"wFxV0sctomzJpypv6Ok8",
								802,
								822,
								swipe{
									Plain:   "gJ5jRBNRbdSK9buzDa0z",
									Meaning: "T2rOHuXIac5WhPzZJn92",
									Peace:   696969,
								},
								[2]string{
									"a4LvOwoVQKImJsdt5uNf",
									"NKjzEKSXbBjXnUgFYYnp",
								},
							},
							Peace: map[string]string{
								"Demonstrator": "2OaXDIBbUMzOuKbZ6xGl",
								"Accent":       "EXdexUnE7OW2I2tWE9kZ",
							},
							Tire: offroad{
								Surprise: []int{449, 686, 727, 715, 81},
								Social:   []float64{959.690, 256.303},
								Bake:     "uogamB1F9U4tFMaVS3BU",
								Deprive:  "i5JMVutvt9KOIMaoHKst",
							},
						},
					},
				},
				Slide: open{
					Cave: map[string]float64{
						"Accent": 686.303,
						"Desert": 13.796,
					},
					Swipe: map[string]swipe{
						"Deserted": {
							Plain:   "St1ABpJxt6l5ktcDnXs6",
							Meaning: "1qC401Uo5gXh2363Aqk2",
							Peace:   99214,
						},
					},
					Powder: "fHFHSVzIxqxIXgMGpgzZ",
					Mill:   []int{896},
					Carbon: 7000.42,
					Consumption: swipe{
						Plain:   "ax3CwALI1upOMk7XIqAi",
						Meaning: "WXTy6UrwVwm4A2gt4gV8",
						Peace:   999999,
					},
					Wall:    "ePLBmgKVncBSAhjit9jb",
					Brother: []float64{3.456, 7.123},
				},
			},
		},
		Identity: 1.12,
		Scramble: []string{},
		Band:     "dWoZA2QqGf9An6Ew25eC",
		Embedded: Embedded{
			Inheritance: 9876,
		},
	}
}

type generalSubTestCase interface {
	DoTest(*assert.Assertions)
}

type getAllSubTestCase[I any, T any] struct {
	in       I
	mp       []traveller.Matcher
	expected []T
	options  []traveller.TravellerOption
}

func (c getAllSubTestCase[I, T]) DoTest(assert *assert.Assertions) {
	actual := traveller.GetAll[T](c.in, c.mp, c.options...)
	assert.ElementsMatch(c.expected, actual)
}

type getSubTestCase[I any, T any] struct {
	in       I
	mp       []traveller.Matcher
	expected T
	notFound bool
	options  []traveller.TravellerOption
}

func (c getSubTestCase[I, T]) DoTest(assert *assert.Assertions) {
	actual, ok := traveller.Get[T](c.in, c.mp, c.options...)
	if c.notFound {
		assert.False(ok)
	} else {
		assert.True(ok)
		assert.Equal(c.expected, actual)
	}
}

type setAllBySubTestCase[I any, T any] struct {
	in         I
	mp         []traveller.Matcher
	setter     traveller.SetterFunc[T]
	expectedFn func() I
	count      int
	options    []traveller.TravellerOption
}

func (c setAllBySubTestCase[I, T]) DoTest(assert *assert.Assertions) {
	in := c.in
	count := traveller.SetAllBy(&in, c.mp, c.setter, c.options...)
	assert.Equal(c.count, count)
	assert.Equal(c.expectedFn(), in)
}

type setBySubTestCase[I any, T any] struct {
	in         I
	mp         []traveller.Matcher
	setter     traveller.SetterFunc[T]
	expectedFn func() I
	changed    bool
	options    []traveller.TravellerOption
}

func (c setBySubTestCase[I, T]) DoTest(assert *assert.Assertions) {
	in := c.in
	changed := traveller.SetBy(&in, c.mp, c.setter, c.options...)
	if c.changed {
		assert.True(changed)
	} else {
		assert.False(changed)
	}
	assert.Equal(c.expectedFn(), in)
}

type GeneralTestSuite struct {
	suite.Suite
}

func TestRunGeneralTestSuite(t *testing.T) {
	suite.Run(t, new(GeneralTestSuite))
}

func (s GeneralTestSuite) TestCallGetAll() {
	cases := []generalSubTestCase{
		getAllSubTestCase[bulb, bulb]{in: makeBulb(), mp: []traveller.Matcher{}, expected: []bulb{makeBulb()}},
		getAllSubTestCase[bulb, string]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Band"}},
			expected: []string{"dWoZA2QqGf9An6Ew25eC"},
		},
		getAllSubTestCase[bulb, any]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Cup"}, traveller.MatchPattern{Pattern: "*"}},
			expected: []any{"ONr7QDhcZJNgiSnZByaH", "VL6foOIq436n8gevZi7K", map[string]string{"Mislead": "yDlqlodvPqwJFB5o8hKq", "Machinery": "nnGbiSSEYt01kotPuVHS"}},
		},
		getAllSubTestCase[bulb, any]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchMulti{}, traveller.MatchPattern{Pattern: "*ll*"}},
			expected: []any{"ePLBmgKVncBSAhjit9jb", []int{896}, []string{"ZgN3U1snmVcUgWNylw0A", "jSIvrxCNBok6lUwHfQgw", "z4xCxlMbQdES77U3hh3k", "a380jUF8DVIE8GNQ3E1G"}},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchMulti{}},
			expected: []int{121, 1021, 2930, 3718, 2848, 1366, 34, 420, 517, 440, 168, 357, 871, 455, 871, 515, 133, 254, 750, 614, 868, 744, 684, 151, 243, 507, 1, 5199, 1999, 740, 555, 874, 864, 582, 807, 802, 822, 696969, 449, 686, 727, 715, 81, 99214, 896, 999999, 9876},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Sunshine"}},
			expected: []int{},
			options:  []traveller.TravellerOption{traveller.WithIgnoreStruct(true)},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchPattern{Pattern: "*unsh*"}},
			expected: []int{},
			options:  []traveller.TravellerOption{traveller.WithIgnoreStruct(true)},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchMulti{}},
			expected: []int{},
			options:  []traveller.TravellerOption{traveller.WithIgnoreStruct(true)},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Cup"}, traveller.MatchExact{Value: "Blasphemy"}},
			expected: []int{},
			options:  []traveller.TravellerOption{traveller.WithIgnoreMap(true)},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchMulti{}, traveller.MatchPattern{Pattern: "*c*"}},
			expected: []int{696969, 999999, 871, 1, 9876},
			options:  []traveller.TravellerOption{traveller.WithIgnoreMap(true)},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchMulti{}},
			expected: []int{121, 1021, 2930, 3718, 2848, 1366, 517, 440, 168, 357, 871, 455, 871, 515, 133, 254, 750, 614, 868, 744, 684, 151, 243, 507, 1, 740, 555, 874, 864, 582, 807, 802, 822, 696969, 449, 686, 727, 715, 81, 896, 999999, 9876},
			options:  []traveller.TravellerOption{traveller.WithIgnoreMap(true)},
		},
		getAllSubTestCase[bulb, string]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Brother"}, traveller.MatchExact{Value: 0}},
			expected: []string{},
			options:  []traveller.TravellerOption{traveller.WithIgnoreArray(true)},
		},
		getAllSubTestCase[bulb, string]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Headache"}, traveller.MatchPattern{Pattern: "*"}},
			expected: []string{},
			options:  []traveller.TravellerOption{traveller.WithIgnoreArray(true)},
		},
		getAllSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchMulti{}},
			expected: []int{121, 34, 420, 871, 868, 1, 1999, 5199, 740, 555, 99214, 999999, 9876},
			options:  []traveller.TravellerOption{traveller.WithIgnoreArray(true)},
		},
	}

	for i, c := range cases {
		s.Run(fmt.Sprintf("Case #%d", i+1), func() {
			c.DoTest(s.Assert())
		})
	}
}

func (s GeneralTestSuite) TestCallMustGetPanic() {
	s.Panics(func() {
		traveller.MustGet[string](makeBulb(), []traveller.Matcher{traveller.MatchExact{Value: "NonExistant"}})
	})
}

func (s GeneralTestSuite) TestCallMustGet() {
	x := makeBulb()
	val := traveller.MustGet[string](x, []traveller.Matcher{traveller.MatchExact{Value: "Brother"}, traveller.MatchExact{Value: 0}})
	s.Equal(x.Brother[0], val)
}

func (s GeneralTestSuite) TestCallGet() {
	cases := []generalSubTestCase{
		getSubTestCase[bulb, bulb]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{},
			expected: makeBulb(),
		},
		getSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Sunshine"}, traveller.MatchExact{Value: "Inner"}},
			notFound: true,
		},
		getSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Sunshine"}},
			expected: 121,
		},
		getSubTestCase[bulb, float64]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchPattern{Pattern: "*rr*"}},
			expected: 370.734,
		},
		getSubTestCase[bulb, float64]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Cup"}, traveller.MatchExact{Value: "Blasphemy"}},
			notFound: true,
		},
		getSubTestCase[bulb, string]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Cup"}, traveller.MatchExact{Value: "Blasphemy"}},
			expected: "ONr7QDhcZJNgiSnZByaH",
		},
		getSubTestCase[bulb, string]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchMulti{}},
			expected: "HFLMPL2rQcwFzlx8dw7D",
		},
		getSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Inheritance"}},
			expected: 9876,
			options:  []traveller.TravellerOption{},
		},
		getSubTestCase[bulb, int]{
			in:       makeBulb(),
			mp:       []traveller.Matcher{traveller.MatchExact{Value: "Inheritance"}},
			notFound: true,
			options:  []traveller.TravellerOption{traveller.WithNoFlatEmbeds(true)},
		},
	}

	for i, c := range cases {
		s.Run(fmt.Sprintf("Case #%d", i+1), func() {
			c.DoTest(s.Assert())
		})
	}
}

func (s GeneralTestSuite) TestCallSetAllPanic() {
	s.Panics(func() {
		traveller.SetAll(69, []traveller.Matcher{}, 0)
	})
}

func (s GeneralTestSuite) TestCallSetAll() {
	actual := makeBulb()
	traveller.SetAll(&actual, []traveller.Matcher{traveller.MatchExact{Value: "Cup"}, traveller.MatchExact{"Houseplant"}, traveller.MatchPattern{Pattern: "*"}}, "this has been edited")
	expected := makeBulb()
	expected.Cup["Houseplant"].(map[string]string)["Mislead"] = "this has been edited"
	expected.Cup["Houseplant"].(map[string]string)["Machinery"] = "this has been edited"
}

func (s GeneralTestSuite) TestCallSetAllBy() {
	editStr := " edited"

	cases := []generalSubTestCase{
		setAllBySubTestCase[bulb, string]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchMulti{}},
			setter: func(oldVal string) (any, bool, bool) {
				if oldVal == "YOSiMHf7S9FCwCvlbnyr" {
					return nil, true, false
				}
				return oldVal + editStr, true, true
			},
			expectedFn: func() bulb {
				x := makeBulb()
				x.Cup["Blasphemy"] = x.Cup["Blasphemy"].(string) + editStr
				x.Cup["Favour"] = x.Cup["Favour"].(string) + editStr
				addToMap(x.Cup["Houseplant"].(map[string]string), editStr)
				addToMap(x.Federation.Decline, editStr)
				addToMap(x.Federation.Hate.Create.Fence.Knowledge.Peace, editStr)
				x.Federation.Jet.Party += editStr
				x.Federation.Jet.Tiger += editStr
				x.Federation.Hate.Enthusiasm += editStr
				x.Federation.Hate.Investment = x.Federation.Hate.Investment.(string) + editStr
				x.Federation.Hate.Explosion += editStr
				x.Federation.Hate.Fax += editStr
				x.Federation.Hate.Create.Fence.Contradiction = x.Federation.Hate.Create.Fence.Contradiction.(string) + editStr
				x.Federation.Hate.Create.Fence.Beach += editStr
				x.Federation.Hate.Create.Fence.Volcano += editStr
				x.Federation.Hate.Create.Fence.Knowledge.Thirsty += editStr
				x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[0] = x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[0].(string) + editStr
				x.Federation.Hate.Create.Fence.Knowledge.Tire.Bake += editStr
				x.Federation.Hate.Create.Fence.Knowledge.Tire.Deprive += editStr
				x.Federation.Hate.Slide.Swipe["Deserted"] = swipe{
					Plain:   x.Federation.Hate.Slide.Swipe["Deserted"].Plain + editStr,
					Meaning: x.Federation.Hate.Slide.Swipe["Deserted"].Meaning + editStr,
					Peace:   x.Federation.Hate.Slide.Swipe["Deserted"].Peace,
				}
				x.Federation.Hate.Slide.Powder += editStr
				x.Federation.Hate.Slide.Consumption = swipe{
					Plain:   x.Federation.Hate.Slide.Consumption.(swipe).Plain + editStr,
					Meaning: x.Federation.Hate.Slide.Consumption.(swipe).Meaning + editStr,
					Peace:   x.Federation.Hate.Slide.Consumption.(swipe).Peace,
				}
				x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[3] = swipe{
					Plain:   x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[3].(swipe).Plain + editStr,
					Meaning: x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[3].(swipe).Meaning + editStr,
					Peace:   x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[3].(swipe).Peace,
				}
				x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[4] = [2]string{
					x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[4].([2]string)[0] + editStr,
					x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)[4].([2]string)[1] + editStr,
				}
				x.Federation.Hate.Slide.Wall += editStr
				x.Band += editStr
				addToSlice(x.Brother[:], editStr)
				addToSlice(x.Federation.Elephant, editStr)
				addToSlice(x.Federation.Hate.College, editStr)
				addToSlice(x.Federation.Hate.Create.Fence.Employee, editStr)
				addToSlice(x.Worth, editStr)
				addToSlice(x.Scramble, editStr)
				return x
			},
			count: 49,
		},
		setAllBySubTestCase[bulb, string]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchMulti{}},
			setter: func(string) (any, bool, bool) {
				return "this has been edited", false, true
			},
			expectedFn: func() bulb {
				x := makeBulb()
				x.Brother[0] = "this has been edited"
				return x
			},
			count: 1,
		},
		setAllBySubTestCase[bulb, any]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchExact{Value: "<Nonexistant>"}},
			setter: func(any) (any, bool, bool) {
				return "this has been edited", true, true
			},
			expectedFn: func() bulb { return makeBulb() },
			count:      0,
		},
		setAllBySubTestCase[bulb, any]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchExact{Value: "Federation"}, traveller.MatchExact{Value: "Hate"}, traveller.MatchExact{Value: "Create"}, traveller.MatchExact{Value: "Fence"}, traveller.MatchExact{Value: "Knowledge"}, traveller.MatchExact{Value: "Job"}, traveller.MatchPattern{Pattern: "*"}},
			setter: func(any) (any, bool, bool) {
				return swipe{
					Plain:   "plain plain",
					Meaning: "meaning meaning",
					Peace:   25,
				}, true, true
			},
			expectedFn: func() bulb {
				x := makeBulb()
				sl := x.Federation.Hate.Create.Fence.Knowledge.Job.([]any)
				for i := range sl {
					sl[i] = swipe{
						Plain:   "plain plain",
						Meaning: "meaning meaning",
						Peace:   25,
					}
				}
				return x
			},
			count: 5,
		},
	}

	for i, c := range cases {
		s.Run(fmt.Sprintf("Case #%d", i+1), func() {
			c.DoTest(s.Assert())
		})
	}
}

func (s GeneralTestSuite) TestCallSetPanic() {
	s.Panics(func() {
		traveller.Set(69, []traveller.Matcher{}, 0)
	})
}

func (s GeneralTestSuite) TestCallSet() {
	actual := makeBulb()
	traveller.Set(&actual, []traveller.Matcher{traveller.MatchExact{Value: "Cup"}, traveller.MatchExact{"Houseplant"}, traveller.MatchPattern{Pattern: "*"}}, "this has been edited")
	expected := makeBulb()
	expected.Cup["Houseplant"].(map[string]string)["Mislead"] = "this has been edited"
}

func (s GeneralTestSuite) TestCallSetBy() {
	cases := []generalSubTestCase{
		setBySubTestCase[bulb, string]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchExact{Value: "Cup"}, traveller.MatchExact{Value: "Houseplant"}, traveller.MatchExact{Value: "Machinery"}},
			setter: func(string) (any, bool, bool) {
				return "this has been edited", true, true
			},
			expectedFn: func() bulb {
				x := makeBulb()
				x.Cup["Houseplant"].(map[string]string)["Machinery"] = "this has been edited"
				return x
			},
			changed: true,
		},
		setBySubTestCase[bulb, any]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchExact{Value: "<Nonexistant>"}},
			setter: func(any) (any, bool, bool) {
				return "this has been edited", true, true
			},
			expectedFn: func() bulb { return makeBulb() },
			changed:    false,
		},
		setBySubTestCase[bulb, any]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchExact{Value: "Federation"}, traveller.MatchExact{Value: "Hate"}, traveller.MatchExact{Value: "Create"}, traveller.MatchExact{Value: "Fence"}, traveller.MatchExact{Value: "Contradiction"}},
			setter: func(any) (any, bool, bool) {
				return map[string]any{}, true, true
			},
			expectedFn: func() bulb {
				x := makeBulb()
				x.Federation.Hate.Create.Fence.Contradiction = map[string]any{}
				return x
			},
			changed: true,
		},
		setBySubTestCase[bulb, string]{
			in: makeBulb(),
			mp: []traveller.Matcher{traveller.MatchMulti{}},
			setter: func(oldVal string) (any, bool, bool) {
				if oldVal == "HFLMPL2rQcwFzlx8dw7D" {
					return nil, true, false
				}
				return "this has been edited", true, true
			},
			expectedFn: func() bulb {
				x := makeBulb()
				x.Worth[0] = "this has been edited"
				return x
			},
			changed: true,
		},
	}

	for i, c := range cases {
		s.Run(fmt.Sprintf("Case #%d", i+1), func() {
			c.DoTest(s.Assert())
		})
	}
}

func addToMap(x map[string]string, extra string) {
	for key := range x {
		x[key] += extra
	}
}

func addToSlice(x []string, extra string) {
	for i := range x {
		x[i] += extra
	}
}
