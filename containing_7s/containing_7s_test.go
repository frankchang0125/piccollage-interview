package containing7s

import (
	"strconv"
	"strings"
    "testing"
)

func TestSolution(t *testing.T) {
	var ans, expectAns int	
    
    ans = Solution(0)
	expectAns = correctAns(0)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
	
	ans = Solution(1)
	expectAns = correctAns(1)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(7)
	expectAns = correctAns(7)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution(17)
	expectAns = correctAns(17)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(20)
	expectAns = correctAns(20)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
	
	ans = Solution(56)
	expectAns = correctAns(56)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
    
    ans = Solution(70)
	expectAns = correctAns(70)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
    
    ans = Solution(100)
	expectAns = correctAns(100)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(723)
	expectAns = correctAns(723)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
    }
	
	ans = Solution(823)
	expectAns = correctAns(823)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
    
    ans = Solution(1000)
	expectAns = correctAns(1000)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(2207)
	expectAns = correctAns(2207)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(3499)
	expectAns = correctAns(3499)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(7777)
	expectAns = correctAns(7777)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}	

	ans = Solution(8677)
	expectAns = correctAns(8677)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(8697)
	expectAns = correctAns(8697)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(8797)
	expectAns = correctAns(8797)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(9408)
	expectAns = correctAns(9408)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(9725)
	expectAns = correctAns(9725)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}

	ans = Solution(12645209)
	expectAns = correctAns(12645209)
	if ans != expectAns {
		t.Errorf("Expect answer to be %d, but got %d.", expectAns, ans)
	}
}

func correctAns(num int) int {
	count := 0
	for i := 1; i <= num; i++ {
		s := strconv.Itoa(i)
		if strings.Contains(s, "7") {
			count++
		}
	}
	return count
}
