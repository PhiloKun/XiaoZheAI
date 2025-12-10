package main

import "xiaozhe-ai-backend/cmd"

func main() {
	defer cmd.Clean()
	cmd.Start()
}
