package traveller_test

import (
	"fmt"
	"testing"

	"github.com/pwnedgod/traveller"
	"github.com/stretchr/testify/suite"
)

type pathSubTestCase struct {
	in              string
	caseInsensitive bool
	expected        []traveller.Matcher
	err             error
}

type PathTestSuite struct {
	suite.Suite
}

func TestRunPathTestSuite(t *testing.T) {
	suite.Run(t, new(PathTestSuite))
}

func (s PathTestSuite) TestCallP() {
	expectedMp := []traveller.Matcher{
		traveller.MatchExact{Value: "something"},
		traveller.MatchMulti{},
		traveller.MatchExact{Value: "something"},
		traveller.MatchPattern{Pattern: "so*th*ing"},
	}
	mp := traveller.P("something.**.something.so*th*ing")
	s.Equal(expectedMp, mp)
}

func (s PathTestSuite) TestCallPCI() {
	expectedMp := []traveller.Matcher{
		traveller.MatchPattern{Pattern: "something", Options: traveller.MatchPatternOptions{CaseInsensitive: true}},
		traveller.MatchMulti{},
		traveller.MatchPattern{Pattern: "something", Options: traveller.MatchPatternOptions{CaseInsensitive: true}},
		traveller.MatchPattern{Pattern: "so*th*ing", Options: traveller.MatchPatternOptions{CaseInsensitive: true}},
	}
	mp := traveller.PCI("something.**.something.so*th*ing")
	s.Equal(expectedMp, mp)
}

func (s PathTestSuite) TestCallMustPathPanic() {
	s.Panics(func() {
		traveller.MustPath("***", true)
	})
}

func (s PathTestSuite) TestCallMustPath() {
	expectedMp := []traveller.Matcher{
		traveller.MatchExact{Value: "something"},
		traveller.MatchMulti{},
		traveller.MatchExact{Value: "something"},
		traveller.MatchPattern{Pattern: "so*th*ing"},
	}
	mp := traveller.MustPath("something.**.something.so*th*ing", false)
	s.Equal(expectedMp, mp)
}

func (s PathTestSuite) TestCallPath() {
	cases := []pathSubTestCase{
		// Case sensitive.
		{
			in:              "something",
			caseInsensitive: false,
			expected:        []traveller.Matcher{traveller.MatchExact{Value: "something"}},
		},
		{
			in:              "some*",
			caseInsensitive: false,
			expected:        []traveller.Matcher{traveller.MatchPattern{Pattern: "some*"}},
		},
		{
			in:              "**",
			caseInsensitive: false,
			expected:        []traveller.Matcher{traveller.MatchMulti{}},
		},
		{
			in:              "***.***",
			caseInsensitive: false,
			err:             traveller.ErrInvalidPath,
		},
		{
			in:              "**something*",
			caseInsensitive: false,
			err:             traveller.ErrInvalidPath,
		},
		{
			in:              "nested1.**.*nest*",
			caseInsensitive: false,
			expected:        []traveller.Matcher{traveller.MatchExact{Value: "nested1"}, traveller.MatchMulti{}, traveller.MatchPattern{Pattern: "*nest*"}},
		},
		{
			in:              "nested1\\.*nest*",
			caseInsensitive: false,
			expected:        []traveller.Matcher{traveller.MatchPattern{Pattern: "nested1.*nest*"}},
		},

		// Case insensitive.
		{
			in:              "something",
			caseInsensitive: true,
			expected:        []traveller.Matcher{traveller.MatchPattern{Pattern: "something", Options: traveller.MatchPatternOptions{CaseInsensitive: true}}},
		},
		{
			in:              "some*",
			caseInsensitive: true,
			expected:        []traveller.Matcher{traveller.MatchPattern{Pattern: "some*", Options: traveller.MatchPatternOptions{CaseInsensitive: true}}},
		},
		{
			in:              "**",
			caseInsensitive: true,
			expected:        []traveller.Matcher{traveller.MatchMulti{}},
		},
		{
			in:              "***.***",
			caseInsensitive: true,
			err:             traveller.ErrInvalidPath,
		},
		{
			in:              "**something*",
			caseInsensitive: true,
			err:             traveller.ErrInvalidPath,
		},
		{
			in:              "nested1.**.*nest*",
			caseInsensitive: true,
			expected:        []traveller.Matcher{traveller.MatchPattern{Pattern: "nested1", Options: traveller.MatchPatternOptions{CaseInsensitive: true}}, traveller.MatchMulti{}, traveller.MatchPattern{Pattern: "*nest*", Options: traveller.MatchPatternOptions{CaseInsensitive: true}}},
		},
		{
			in:              "nested1\\.*nest*",
			caseInsensitive: true,
			expected:        []traveller.Matcher{traveller.MatchPattern{Pattern: "nested1.*nest*", Options: traveller.MatchPatternOptions{CaseInsensitive: true}}},
		},
	}

	for i, c := range cases {
		s.Run(fmt.Sprintf("Case #%d", i+1), func() {
			mp, err := traveller.Path(c.in, c.caseInsensitive)
			if c.err == nil {
				s.NoError(err)
				s.Equal(c.expected, mp)
			} else {
				s.ErrorIs(err, c.err)
			}
		})
	}
}
