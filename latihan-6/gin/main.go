package main

import "latihan-6-gin/routers"

func main() {
	const PORT = ":5000"

	routers.StartServer().Run(PORT)
}
