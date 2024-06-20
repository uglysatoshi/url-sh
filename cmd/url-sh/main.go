package main

import (
    "url-sh/internal/config"
)

func main() {
    _ = config.MustLoad()
}
