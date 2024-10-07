// 1-) Yapmak istediğimiz: Sudoku programını önce çözümsüz yazmak sonra da çözümünü sudoku kurallarına göre yapmasını istemek

// 2-) Plan Oluşturmak: Sudoku ızgarısını nasıl temsil edeceğini, hangi fonksiyonların gerekli olduğunu(örneğin, Sudoku'nun geçerli olup olmadığını kontrol eden bir fonksiyon) ve algoritmayı adım adım nasıl uygulayacağını belirlemek

// 3-) Sudoku ızgarasını temsil etmek için dizi(array) ve dilim(slice) yapılarından diziyi(array)'ı, sabit boyutlu bir sudoku yapacağımız için seçtik.

package main

import "fmt"

func main() {
	// 4-) Sudoku ızgarasını başlangıçta boş olmayan değerlerle tanımlıyoruz, 0'lar sudokudaki boş alanları temsil ediyor
	var sudoku = [9][9]int{
		{9, 0, 4, 6, 0, 3, 0, 0, 1},
		{3, 7, 0, 1, 0, 0, 2, 0, 6},
		{0, 0, 6, 0, 0, 9, 3, 0, 4},
		{0, 0, 1, 3, 0, 0, 9, 0, 5},
		{5, 6, 0, 0, 9, 1, 0, 0, 0},
		{8, 2, 0, 0, 0, 4, 6, 1, 0},
		{0, 0, 7, 9, 0, 0, 0, 4, 0},
		{4, 2, 5, 0, 1, 6, 7, 0, 0},
		{1, 0, 2, 0, 0, 7, 5, 0, 8},
	}

	fmt.Println("Başlangıç Sudoku:")
	printSudoku(sudoku)

	if !isValidSudoku(sudoku) {
		fmt.Println("Verilen Sudoku kurallara uygun değil.")
		return
	}

	if solveSudoku(&sudoku) {
		fmt.Println("Çözülmüş Sudoku:")
		printSudoku(sudoku)
	} else {
		fmt.Println("Sudoku çözülemedi.")
	}

	// Sudoku ızgarasını yazdırıyoruz
	// fmt.Println(sudoku)

	// if isValidSudoku(sudoku) {
	// 	fmt.Println("Sudoku geçerli.")
	// } else {
	// 	fmt.Println("Sudoku geçersiz.")
	// }

	// printSudoku(sudoku)

	// fmt.Println("Başlangıç Sudoku:")
	// printSudoku(sudoku)

	// if solveSudoku(&sudoku) {
	// 	fmt.Println("Çözülmüş Sudoku:")
	// 	printSudoku(sudoku)
	// } else {
	// 	fmt.Println("Sudoku çözülemedi.")
	// }
}

// 5-) Her satır, sütun ve 3x3'lük alanda 1'den 9'a kadar her bir rakamın sadece 1 kez yazılıp yazılmadığını kontrol etmek için, kısacası Sudoku'nun temel kuralına uygun hareket etmek için, veri doğrulama işlemi yapmamız, yani hatalı veri girmemek ve hatalı veri girişinde sudokuyu ve çözümünü tespit edebilmek için, kuralların veri girişine göre ihlal edilip edilmediğini kontrol etmek için veri doğrulama yapıyoruz.

// Satır doğrulama fonksiyonu
func checkRow(sudoku [9][9]int, row int) bool {
	seen := [10]bool{} // 1'den 9'a kadar olan sayıları takip etmek için bir dizi(köşeli parantezin içine 10 yazılmasının sebebi boş alanları belirtirken 0 rakamını da kullandığımızdan dolayıdır.)
	for i := 0; i < 9; i++ {
		num := sudoku[row][i] // satırların içindeki hücreleri kontrol ettiğinin temsilidir.
		if num != 0 {         // Eğer hücre boş değilse
			if seen[num] {
				return false // Aynı sayı tekrar bulunursa geçersiz
			}
			seen[num] = true
		}
	}
	return true
}

// Sütun doğrulama fonksiyonu
func checkColumn(sudoku [9][9]int, col int) bool {
	seen := [10]bool{} // 1'den 9'a kadar olan sayıları takip etmek için
	for i := 0; i < 9; i++ {
		num := sudoku[i][col]
		if num != 0 { // Eğer hücre boş değilse
			if seen[num] {
				return false // Aynı sayı tekrar bulundu
			}
			seen[num] = true
		}
	}
	return true
}

// 3x3 bölge doğrulama fonksiyonu
func checkBox(sudoku [9][9]int, startRow, startCol int) bool {
	seen := [10]bool{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			num := sudoku[startRow+i][startCol+j]
			if num != 0 {
				if seen[num] {
					return false // Aynı sayı tekrar bulundu
				}
				seen[num] = true
			}
		}
	}
	return true
}

// Genel doğrulama fonksiyonu
func isValidSudoku(sudoku [9][9]int) bool {
	// Satır kontrolü
	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if sudoku[i][j] != 0 {
				if seen[sudoku[i][j]] {
					return false // Aynı sayıyı birden fazla kez gördük
				}
				seen[sudoku[i][j]] = true
			}
		}
	}

	// Sütun kontrolü
	for i := 0; i < 9; i++ {
		seen := make(map[int]bool)
		for j := 0; j < 9; j++ {
			if sudoku[j][i] != 0 {
				if seen[sudoku[j][i]] {
					return false
				}
				seen[sudoku[j][i]] = true
			}
		}
	}

	// 3x3 bölge kontrolü
	for boxRow := 0; boxRow < 3; boxRow++ {
		for boxCol := 0; boxCol < 3; boxCol++ {
			seen := make(map[int]bool)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					val := sudoku[boxRow*3+i][boxCol*3+j]
					if val != 0 {
						if seen[val] {
							return false
						}
						seen[val] = true
					}
				}
			}
		}
	}

	return true // Eğer geçerliyse

	// // Satır ve sütunları kontrol et
	// for i := 0; i < 9; i++ {
	// 	if !checkRow(sudoku, i) || !checkColumn(sudoku, i) {
	// 		return false
	// 	}
	// }

	// // 3x3 bölgeleri kontrol et
	// for i := 0; i < 9; i += 3 {
	// 	for j := 0; j < 9; j += 3 {
	// 		if !checkBox(sudoku, i, j) {
	// 			return false
	// 		}
	// 	}
	// }

	// return true // Eğer tüm kontroller geçerse Sudoku geçerli
}

func printSudoku(sudoku [9][9]int) {
	for i := 0; i < 9; i++ {
		if i%3 == 0 && i != 0 {
			fmt.Println("---------------------") // 3x3 bloklar arası ayırıcı
		}
		for j := 0; j < 9; j++ {
			if j%3 == 0 && j != 0 {
				fmt.Print("| ") // Sütunlar arası ayırıcı
			}
			if sudoku[i][j] == 0 {
				fmt.Print(". ") // Boş hücreyi noktayla göster
			} else {
				fmt.Print(sudoku[i][j], " ")
			}
		}
		fmt.Println() // Satırın sonu
	}
}

func solveSudoku(sudoku *[9][9]int) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if sudoku[row][col] == 0 { // Boş hücre bul
				for num := 1; num <= 9; num++ { // 1'den 9'a kadar sayıları dene
					if isValidPlacement(sudoku, row, col, num) {
						sudoku[row][col] = num   // Geçerli ise sayıyı yerleştir
						if solveSudoku(sudoku) { // İlerle
							return true
						}
						sudoku[row][col] = 0 // Geri adım at
					}
				}
				return false // Eğer hiçbiri geçerli değilse
			}
		}
	}
	return true // Eğer tüm hücreler dolmuşsa
}

func isValidPlacement(sudoku *[9][9]int, row, col, num int) bool {
	// Satır ve sütun kontrolü
	for i := 0; i < 9; i++ {
		if sudoku[row][i] == num || sudoku[i][col] == num {
			return false
		}
	}

	// 3x3 kutu kontrolü
	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if sudoku[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true // Eğer geçerliyse
}
