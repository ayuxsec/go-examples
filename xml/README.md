```go
import "encoding/xml"
```

**Unmarshal XML to struct:**

```go
type Person struct {
    XMLName xml.Name `xml:"person"`
    Name    string   `xml:"name"`
    Age     int      `xml:"age"`
}

data := []byte(`<person><name>John</name><age>30</age></person>`)
var p Person
xml.Unmarshal(data, &p)
// p.Name = "John", p.Age = 30
```

**Nested elements:**

```go
type Book struct {
    Title  string `xml:"title"`
    Author struct {
        Name string `xml:"name"`
    } `xml:"author"`
}

data := []byte(`<book><title>Go</title><author><name>Pike</name></author></book>`)
var b Book
xml.Unmarshal(data, &b)
// b.Author.Name = "Pike"
```

**Attributes:**

```go
type Item struct {
    ID    string `xml:"id,attr"`
    Value string `xml:",chardata"`
}

data := []byte(`<item id="123">content</item>`)
var i Item
xml.Unmarshal(data, &i)
// i.ID = "123", i.Value = "content"
```

**Arrays/slices:**

```go
type Library struct {
    Books []Book `xml:"book"`
}

data := []byte(`<library><book><title>A</title></book><book><title>B</title></book></library>`)
var lib Library
xml.Unmarshal(data, &lib)
// len(lib.Books) = 2
```

**Decoder for streaming/large files:**

```go
decoder := xml.NewDecoder(strings.NewReader(xmlString))
for {
    token, _ := decoder.Token()
    if token == nil { break }
    if se, ok := token.(xml.StartElement); ok {
        if se.Name.Local == "person" {
            var p Person
            decoder.DecodeElement(&p, &se)
        }
    }
}
```
