/*
Anonymous Functions : İsimsiz fonksiyonlar
Yazıldıkları yerde direkt olarak çalışırlar.
Çalışırken diğer fonksiyonlardaki gibi parametre verilemediği için fonksiyonun
sonuna parametre eklenerek çalışıtırılırlar.
*/
package main

import "fmt"

func main() {
	text := "Champion Galatasaray"
	func(a string) {
		fmt.Println(a)
	}(text)
}
