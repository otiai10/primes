package main

func main() {
	cmd := getCommand()
	cmd.Prepare()
	cmd.Perform()
}
