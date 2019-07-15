package containing7s

import (
    "strconv"
)

func Solution(num int) int {
    numStr := strconv.Itoa(num)
    digits := len(numStr)
    numOf7 := 0

    for i := 0; i < digits; i++ {
        d := int(numStr[i] - '0')
        if d == 7 {
            others, _ := strconv.Atoi(numStr[i+1:])
            numOf7 += 7 * numOf7Pow10(digits - i) + others + 1
            break
        } else if d < 7 {
            numOf7 += d * numOf7Pow10(digits - i)
        } else {
            // d > 7
            numOf7 += (d - 1) * numOf7Pow10(digits - i) + pow10(digits - i - 1)
        }
    }

    return numOf7
}

func numOf7Pow10(digits int) int {
    if digits <= 1 {
        return 0
    } else if digits == 2 {
        return 1
    }

    result := 1
    for i := 3; i <= digits; i++ {
        result = 9 * result + pow10(i-2)
    }
    return result
}

func pow10(x int) int {
    result := 1
    for i := 0; i < x; i++ {
        result *= 10
    }
    return result
}
