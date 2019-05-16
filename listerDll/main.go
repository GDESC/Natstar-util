package main

import (
    "bufio"
    "fmt"
	"os"
	"strings"	
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez le chemin de votre export: ")
	
	filepath, err := reader.ReadString('\n')
	check(err)
	
	file, err := os.Open(strings.Split(filepath,"\n")[0])
	check(err)

	
	deleteFileIfExists("./ListeDll.txt")
	dllFile, err:= os.Create("./ListeDll.txt")
	check(err)	

	deleteFileIfExists("./ListeDllServeur.txt")
	remoteDllFile, err:= os.Create("./ListeDllServeur.txt")
	check(err)

	defer file.Close()
	defer dllFile.Close()
	defer remoteDllFile.Close()

	listRemoteDll, listDll:="",""
	nbDll,nbRemoteDll:=0,0
	isRemote:=false

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if (strings.Contains(scanner.Text(), "REMOTE = YES")) {
			isRemote=true
		}

		if (strings.Contains(scanner.Text(), "End of Library")) {
			if (isRemote) {
				listRemoteDll+=extractDllName(scanner.Text())
				nbRemoteDll++
			}

			listDll+=extractDllName(scanner.Text())
			nbDll++
			isRemote=false
        }
	}
	
	err = scanner.Err()
	check(err)

	dllFile.WriteString(listDll)
	remoteDllFile.WriteString(listRemoteDll)

	fmt.Println("nb DLL: ", nbDll)
	fmt.Println("nb DLL Serveur: ", nbRemoteDll)
}

func extractDllName(dllEndLine string) string{
	resTMP:=strings.Split(dllEndLine,"Library ")
	res:=strings.Split(resTMP[len(resTMP)-1]," ")[0]+"\n"
	return res
}

func deleteFileIfExists(filepath string) {
	if _, err := os.Stat(filepath); os.IsExist(err) {
		os.Remove(filepath)
	}
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}
