package main

import "fmt"

func main() {
	var viking, saxon int

	fmt.Scan(&viking, &saxon)

	vikingVS(viking, saxon)
}

func vikingVS(viking, saxon int) {
	/* I.S. Terdefinisi  jumlah viking dan saxon int
	   F.S.  Menampilkan string siapa pemenangnya*/

	if (4 * viking) > saxon {
		fmt.Println("viking menang")
	} else if (4 * viking) < saxon {
		fmt.Println("saxon menang")
	} else {
		fmt.Println("imbang")
	}
}
