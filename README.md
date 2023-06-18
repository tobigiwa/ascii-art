# ASCII-ART

![my badge](https://badgen.net/badge/hello/world/red?icon=github)
### Description
**This is a Go-based project that generates ASCII art text banners in styles, including shadow, standard, and thinkertoy.**


### Usage
![front_page](/images/frontpage.png)

The page shows where to enter our test and select desired banner


![standard banner](/images/standard.png)


The text was from a quote found on the book ![Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/), showing the **STANDARD BANNER**


> “Go is unique and even experienced programmers have to unlearn a few things and think
differently about software. Learning Go does a good job of working through the big
features of the language while pointing out idiomatic code, pitfalls,
and design patterns along the way.” —Aaron Schlesinger, Sr. Engineer, Microsoft

![shadow banner](/images/shadow.png)

**SHADOW BANNER**


![thinkertoy banner](/images/thinkertoy.png)

**THINKERTOY BANNER**

### Build

> go run cmd/web/ascii-art.go

- To produce the executable

> go build cmd/web/ascii-art.go

## Contributing

Contributions to this project are welcome! To contribute, fork the repository and create a pull request with your changes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.

### Implementation details
- I used the go embed directive to include the templates in the binary for faster loading.
- Wrote a small middleware to feature HTPP verb routes.
- Appropriate test and benchmarking



