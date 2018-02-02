package main

import (
  "fmt"
  "strings"
  "log"
  "os/exec"
  "os"
  // "reflect"
  "regexp"

  "github.com/atotto/clipboard"
	"github.com/manifoldco/promptui"

)

func main() {
  var arguments = os.Args[1:]
  var command = strings.Join(arguments, " ")
  command_stdout, err := exec.Command("bash", "-c", command).Output()
  if err != nil {
    log.Fatal(err)
  }
  clean_output := remove_empty_line(string(command_stdout))
  select_items := strings.Split(clean_output, "\n")

  prompt := promptui.Select{
		Label: command,
		Items: select_items,
    Size: 10,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

  result = strings.TrimSpace(result)

	fmt.Printf("You choose %q\n", result)

  // yank to past_board
  clipboard.WriteAll(result)
}

func remove_empty_line(str string) string {
  regex, err := regexp.Compile("\n\n")
  if err != nil {
    log.Fatal(err)
  }
  return regex.ReplaceAllString(str, "\n")
}
