package main

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			return true
		}
		return false
	}
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if i == 0 {
				if flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--
				}
			} else if i == len(flowerbed)-1 {
				if flowerbed[i-1] == 0 {
					flowerbed[i] = 1
					n--
				}
			} else {
				if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
					flowerbed[i] = 1
					n--
				}
			}
		}
		if n == 0 {
			return true
		}
	}
	return false
}
