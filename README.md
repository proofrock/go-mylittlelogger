# go-mylittlelogger

A small, useful, non-structured logger for Go(lang)

# why?

Because I need a non-structured logger, where I can substitute the behaviour
of `log.Fatal` (call `os.Exit()`) during tests, and also disable the log
altogether.
