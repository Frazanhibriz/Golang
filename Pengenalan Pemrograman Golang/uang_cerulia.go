package main

import "fmt"

func main() {
	var quantum_shard, galactum, nebula_crown, stellaris, lumin int

	fmt.Scan(&lumin)

	stellaris = lumin / 20

	lumin = lumin % 20

	nebula_crown = stellaris / 2

	stellaris = stellaris % 2

	galactum = nebula_crown / 3

	nebula_crown = nebula_crown % 3

	quantum_shard = galactum / 10

	galactum = galactum % 10

	fmt.Println(quantum_shard, galactum, nebula_crown, stellaris, lumin)
}
