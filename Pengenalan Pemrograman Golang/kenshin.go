package main

import "fmt"

func main() {
	var n, kekuatan_kenshin, kecepatan_kenshin, kekuatan_lawan, kecepatan_lawan, musuh_kalah int

	fmt.Scan(&n, &kekuatan_kenshin, &kecepatan_kenshin)

	for i := 1; i <= n; i++ {
		fmt.Scan(&kekuatan_lawan, &kecepatan_lawan)
		if kekuatan_kenshin >= 3 && kecepatan_kenshin >= 3 {
			if kekuatan_lawan >= kekuatan_kenshin && kecepatan_lawan < kecepatan_kenshin {
				kecepatan_kenshin -= 6
				musuh_kalah++
			} else if kekuatan_lawan < kekuatan_kenshin && kecepatan_lawan >= kecepatan_kenshin {
				kekuatan_kenshin -= 6
				musuh_kalah++
			} else if kekuatan_lawan < kekuatan_kenshin && kecepatan_lawan < kecepatan_kenshin {
				if kekuatan_kenshin > kecepatan_kenshin {
					kekuatan_kenshin -= 6
					musuh_kalah++
				} else if kekuatan_kenshin < kecepatan_kenshin {
					kecepatan_kenshin -= 6
					musuh_kalah++
				} else if kekuatan_kenshin == kecepatan_kenshin {
					kecepatan_kenshin -= 6
					musuh_kalah++
				}
			}
		}
	}
	fmt.Println(musuh_kalah, kekuatan_kenshin, kecepatan_kenshin)
}
