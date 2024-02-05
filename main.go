package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println("Hello Go")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("testCalculate!")
	fmt.Println("--------------")

	for {
		fmt.Print("->")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1)
		text = strings.Replace(text, " ", "", -1)

		chr := rune(text[0])

		if chr == 'I' || chr == 'i' || chr == 'V' || chr == 'v' || chr == 'X' || chr == 'x' || chr == 'Х' || chr == 'х' {
			fmt.Println(CalculateRome(text))
		} else if int(chr-'0') <= 9 && int(chr-'0') >= 0 {
			fmt.Println(CalculateArb(text))
		} else {
			err := errors.New("incorrect symbol, number of index: 0")
			panic(err)
		}
	}
}

func CalculateArb(str string) int {
	_leftSide := 0
	var _operator rune
	_rightSide := 0

	_tmpInt := 0
	var _chngSide bool = false

	for index, value := range str {
		_value := int(value - '0')

		if _value >= 0 && _value <= 9 {
			if _tmpInt == 0 {
				_tmpInt = _value
			} else {
				_tmpInt = _tmpInt*10 + _value
				if _tmpInt > 10 {
					err := errors.New("\"10\" it is max operand count")
					panic(err)
				}
			}
		} else if value == '+' || value == '-' || value == '/' || value == '*' {
			if _chngSide == true {
				panic(index)
			}
			_operator = value
			_tmpInt = 0
			_chngSide = true
		} else {
			err := errors.New(fmt.Sprint("incorrect symbol, number of index: ", index))
			panic(err)
		}

		if _chngSide == false {
			_leftSide = _tmpInt
		} else if _chngSide {
			_rightSide = _tmpInt
		}
	}
	var ansvr int
	switch _operator {
	case '+':
		ansvr = _leftSide + _rightSide
	case '-':
		ansvr = _leftSide - _rightSide
	case '/':
		ansvr = _leftSide / _rightSide
	case '*':
		ansvr = _leftSide * _rightSide
	}
	return ansvr
}

func CalculateRome(str string) string {
	_leftSide := 0
	var _operator rune
	_rightSide := 0

	_countI := 0
	_tmpInt := 0
	var _chngSide bool = false

	for index, value := range str {
		if value == 'I' || value == 'i' {
			_countI++
			if _countI > 3 {
				err := errors.New(fmt.Sprint("incorrect number, number of index: ", index))
				panic(err)
			}
			_tmpInt++
		} else if value == 'V' || value == 'v' {
			_tmpInt = 5 - _countI
		} else if value == 'X' || value == 'x' || value == 'Х' || value == 'х' {
			_tmpInt = 10 - _countI
		} else if value == '+' || value == '-' || value == '/' || value == '*' {
			if _chngSide == true {
				panic(index)
			}
			_operator = value
			_countI = 0
			_tmpInt = 0
			_chngSide = true
		} else {
			err := errors.New(fmt.Sprint("incorrect symbol, number of index: ", index))
			panic(err)
		}

		if _tmpInt > 10 {
			err := errors.New("\"X\" it is max operand count")
			panic(err)
		}
		if _chngSide == false {
			_leftSide = _tmpInt
		} else if _chngSide {
			_rightSide = _tmpInt
		}
	}
	var ansvrInt int
	switch _operator {
	case '+':
		ansvrInt = _leftSide + _rightSide
	case '-':
		ansvrInt = _leftSide - _rightSide
	case '/':
		ansvrInt = _leftSide / _rightSide
	case '*':
		ansvrInt = _leftSide * _rightSide
	}
	if ansvrInt < 0 {
		err := errors.New("\"0\" it is min ansver with Rome num")
		panic(err)
	}
	var ansvr string = ""
	counter := ansvrInt / 10
	if counter >= 5 {
		ansvr = "L"

		if counter >= 10 {
			ansvr = "C"
			counter -= 5
		}
		counter -= 5
	}
	for i := 0; i < counter; i++ {
		ansvr += "X"
	}

	switch ansvrInt % 10 {
	case 1:
		ansvr += "I"
	case 2:
		ansvr += "II"
	case 3:
		ansvr += "III"
	case 4:
		ansvr += "IV"
	case 5:
		ansvr += "V"
	case 6:
		ansvr += "VI"
	case 7:
		ansvr += "VII"
	case 8:
		ansvr += "VIII"
	case 9:
		ansvr += "IX"
	}
	return ansvr
}
