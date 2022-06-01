package lab2

import (
	"fmt"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestPrefixToPostfixAnswers(c *C) {
	resultTest1, errTest1 := PrefixToInfix("+ 3 * - 2 2 1")
	c.Check(resultTest1, Equals, "(3+((2-2)*1))")
	c.Check(errTest1, IsNil)

	resultTest2, errTest2 := PrefixToInfix("+ 1 1")
	c.Check(resultTest2, Equals, "(1+1)")
	c.Check(errTest2, IsNil)

	resultTest3, errTest3 := PrefixToInfix("- 1 / 2 3")
	c.Check(resultTest3, Equals, "(1-(2/3))")
	c.Check(errTest3, IsNil)

	resultTest4, errTest4 := PrefixToInfix("- 3 / 5 + 4 * 3 - 3 / + 2 3 10")
	c.Check(resultTest4, Equals, "(3-(5/(4+(3*(3-((2+3)/10))))))")
	c.Check(errTest4, IsNil)

}

func (s *MySuite) TestPrefixToPostfixErrors(c *C) {
	_, errTest1 := PrefixToInfix("$ 2 2")
	c.Check(errTest1, NotNil)
	c.Check(fmt.Sprintf("%s", errTest1), Equals, "error in input")

	_, errTest2 := PrefixToInfix("1 1 +")
	c.Check(errTest2, NotNil)
	c.Check(fmt.Sprintf("%s", errTest2), Equals, "invalid prefix equation")

	_, errTest3 := PrefixToInfix("+ 1 2 3")
	c.Check(errTest3, NotNil)
	c.Check(fmt.Sprintf("%s", errTest3), Equals, "error in input")

}

func ExamplePrefixToInfix() {
	res, _ := PrefixToInfix("- 3 / 5 10")
	fmt.Println(res)

	// Output:
	// (3-(5/10))
}
