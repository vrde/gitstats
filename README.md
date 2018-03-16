**Disclaimer: this is the first time I write some Go code, if you see some :poop: you can either :see_no_evil: or open a PR and tell me how to fix my code :two_hearts:**

# ghstats
**Problem**: I want to measure if a software project is becoming more and more successful. (And I want to learn some Go, so I need a simple project to start.)

**Solution**: Create a tool to extract stats from a GitHub Organization, specifically how fast issues are handled, bugs addressed, pull requests merged, and what is the ratio of core contributors vs new contributors. All this data should be easy to group on a weekly/monthly/quarterly/yearly basis.

# How the Proof of Concept looks like
For the proof of concept, downloading data from GitHub and creating a CSV file should be enough. After I have the CSV file I can do all kind of queries and understand which ones are the good ones, and implement them in code.

```bash
$ go run cmd/ghstats.go org/repo > issues.csv
```

# References
To write this program I'm reading books and online resources. One day I might also get a [t-shirt](https://medium.com/@ashleymcnamara/gophercon-2018-b9a97387b954)

## Books
- [The Go Programming Language](https://www.gopl.io/)

## Online Resources
- [Go Proverbs](https://go-proverbs.github.io/), words to live by when writing Go.
- [Pointers vs. values in parameters and return values](https://stackoverflow.com/questions/23542989/pointers-vs-values-in-parameters-and-return-values), helped me understand the different patterns regarding that topic.
- [Standard Package Layout – Ben Johnson – Medium](https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1), some neat ideas on how to organize your code, I really like the section "Root package is for domain types".
- [Go Project Layout – golang-learn – Medium](https://medium.com/golang-learn/go-project-layout-e5213cdcfaa2), where should I put the command line interface?

## Read it/Watch it later
- [Go Proverbs - Rob Pike - Gopherfest - November 18, 2015 - YouTube](https://www.youtube.com/watch?v=PAAkCSZUG1c&t=7m36s)
- [CodeReviewComments · golang/go Wiki](https://github.com/golang/go/wiki/CodeReviewComments)

## Tools
- [fatih/vim-go: Go development plugin for Vim](https://github.com/fatih/vim-go), really easy to use, a must-have if you are writing Go in [Neo]Vim.
- [golang/dep: Go dependency management tool](https://github.com/golang/dep), I still have to dig in that.
