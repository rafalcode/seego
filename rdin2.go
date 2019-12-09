
2. io.ByteReader
An alternative is also to use the io.ByteReader interface through a bufio.Reader, like this:

reader := bufio.NewReader(os.Stdin)
for {
    b, err := reader.ReadByte()
    if err != nil {
        break
    }

    // do stuff
}
