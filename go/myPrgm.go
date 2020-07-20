package main
 import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
    "bufio"
    "strconv"
)

func check(e error) {
	if e != nil{
		panic(e)
	}
}
/* readMe function takes the file name,
opens it,
& returns the content of the file -anything but numbers */

func readMe(filename string) string {
        filebuffer, err := ioutil.ReadFile(filename)
        if err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
	check(err)
        inputdata := string(filebuffer)
        data := bufio.NewScanner(strings.NewReader(inputdata))
        data.Split(bufio.ScanRunes)

	var result string
        result = ""
        for data.Scan() {
		if (IsNumeric(string(data.Text())) == false) {
		//	fmt.Println(data.Text())
			result = result + data.Text()
		}
	}
	//fmt.Println("the concatenated one ", result)
	return result
}



// function parses out all numbers in the imported file data
func IsNumeric(s string) bool {
   _, err := strconv.ParseFloat(s, 64)
   return err == nil
}


/* this function takes newfile string and creates a file
it also takes a string, content, & put it in the created file */

func write_output (content string, newfile_name string) {
	d1 :=[]byte(content)
	err := ioutil.WriteFile(newfile_name, d1, 0644)
	check(err)
}

 func main() {
    // Program Name is always the first (implicit) argument
    file_name := os.Args[1]
    newfile_name := os.Args[2]
    //fmt.Printf("Program Name: %s\n", file_name)
    //write_output(file_name, newfile_name)
    content := readMe(file_name)
    //fmt.Println(readMe(file_name))
    write_output(content, newfile_name)
}
