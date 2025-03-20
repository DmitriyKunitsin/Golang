package main

import "fmt"

func main() {
	nibbles := []byte{0x12, 0x34}
	head := []byte{0x14, 0x06}
	MergingPacket(nibbles, head)
	nibbles = []byte{0x12, 0x34}
	head = []byte{0x14, 0x06}
	for _ , value := range nibbles {
		head  = append(head, value)
	}
	for i, _ := range head {
		fmt.Printf("%#x\t", head[i])
	}
}

func MergingPacket(nibles []byte, head []byte) {
	for _, value := range nibles {
		head = append(head, value)
	}
	fmt.Printf("%d \n", len(head))
	head[1] = head[1]<<4 | (nibles[0] >> 4) & 0x0F 

	for i := 2; i < len(head); i ++ {
		tmp := head[i] & 0x0F
		fmt.Printf("%#x\t", tmp)
		if i + 1 >= len(head) {
			head[i] = (head[i] & 0x0F)
		} else {
			head[i] = ( tmp  | (head[i+1] & 0xF0))
		}
	}
	fmt.Printf("\n\n")

	for _, value := range head {
		fmt.Printf("%#x\t", value)
	}
}
func zipOnePacket(nibles []byte) {
	for i := 0; i < len(nibles); i++ {
		fmt.Printf("%#x\t", nibles[i])
		tmp := nibles[i] & 0xF0
		fmt.Printf("%#x\t", tmp)
		nibles[i] = nibles[i]<<4 | (tmp>>4)&0x0F
		fmt.Printf("%#x\n", nibles[i])
	}
}
