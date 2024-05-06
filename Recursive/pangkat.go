package main

import "fmt"

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	hasil := pangkat(m, n, 0)
	fmt.Println(hasil)
}

func pangkat(m, n, i int) int {
	if i == n {
		return 1
	} else {
		return m * pangkat(m, n, i+1)
	}
}

/*

program pangkat
kamus
	m, n, hasil : integer
algoritma
	input(m, n)
	hasil <- pangkat(m, n, 0)
	output(hasil)
endprogram

function pangkat(m, n, i : integer) -> integer
    if i == n then
        return 1
    else
        return m * pangkat(m, n, i+1)
    end if
end function

*/
