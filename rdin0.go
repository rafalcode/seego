ref. https://stackoverflow.com/questions/26811600/is-there-a-go-function-which-is-similar-to-getchar-until-eof￼
Two techniques pop to mind:

1. ioutil.ReadAll
From the documentation:

ReadAll reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.

For example:

byteSlice, err := ioutil.ReadAll(os.Stdin)
if err != nil {
    log.Fatal(err)
}

for b := range byteSlice {
    // do stuff
}
