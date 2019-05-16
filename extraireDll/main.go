package main

import (
	"bufio"
    "fmt"
	"os"
	"path/filepath"
)

func main() {

	rootPath:=getInput("Entrez le chemin du repertoire contenant les DLL: ")
	standardConfig:=getInput("Entrez le nom de votre profil de compilations pour les DLL clients: ")
	remoteConfig:=getInput("Entrez le nom de votre profil de compilations pour les DLL serveurs: ")

	err:=os.Mkdir("DllClient", 0755)
	check(err)

	err=os.Mkdir("DllServeur", 0755)
	check(err)

	fileListDll, err := os.Open("./ListeDll.txt")
	check(err)
	defer fileListDll.Close()

	scanner := bufio.NewScanner(fileListDll)

	for scanner.Scan() {
		dllPath:=findFile(rootPath+"/"+"C_"+scanner.Text()+"/"+standardConfig,{"*"}string[])
		io.Copy(dllPath,"./DllClient/"+strin)
	}
	
	err = scanner.Err()
	
}
func getInput(s string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(s)
	
	rootPath, err := reader.ReadString('\n')
	check(err)
}

func findFile(dirPath string, pattern []string) string {
	for _, v := range pattern {
		matches, err := filepath.Glob(dirPath + v)
		check(err)

		if len(matches) != 0 {
			fmt.Println("Found : ", matches)
		}
	}
}
func check(e error) {
    if e != nil {
        panic(e)
    }
}