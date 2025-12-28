package main

import "github.com/philokun/xiaozhe-ai-backend/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
