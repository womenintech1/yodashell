# yodashell

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func execute() {
	//execute-uygulamak

	// burada pwd komutunu gerçekleştiriyoruz.
	// bunun çıktısını out değişkenimizde saklayabiliriz
	// ve hataları err içinde yakalayın

	out, err := exec.Command("ls", "-l").Output()

	// yürütme işlemimizde bir hata varsa
	// burada halledin

	if err != nil {
		fmt.Printf("%s", err)

		// != -Eşit değil işleci ( != ), işlenenler aynı değere sahip değilse true döndürür; aksi takdirde false döndürür
		// nil - Go'daki nil basitçe diğer dillerdeki NULL işaretçi değeridir.
		// %s -Sprintf . fmt.Printf("%s is %d years old\n", name, age) fmt. Printf işlevi konsola biçimlendirilmiş bir dize yazdırır.
		// %s bir dize değeri ve %d bir tamsayı değeri bekler.
	}

	// yukarıda tanımlanan out değişkeni []byte türünde olduğu için dönüştürmemiz gerekir
	// bunu bir dizeye dönüştürün, aksi takdirde konsolumuzda çöp yazdırıldığını göreceğiz
	// bu şekilde bir dizeye dönüştürüyoruz

	fmt.Println("Befehl erfolgreich durchgeführt")
	output := string(out[:])
	fmt.Println(output)

	//ls -l komutunu burada deneyelim

	// cd komutunu burada deneyelim

	out, err = exec.Command("cd").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Befehl erfolgreich durchgeführt")
	//[:] Kurz gesagt, der Operator [:] erlaubt es Ihnen, ein Slice aus einem Array zu erstellen, optional mit Start- und Endgrenzen
	output = string(out[:])
	fmt.Println(output)

	out, err = exec.Command("mkdir").Output()
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Println("Befehl erfolgreich durchgeführt")
	output = string(out[:])
	fmt.Println(output)

}

func main() {

	fmt.Println(isWindows())
}

func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
