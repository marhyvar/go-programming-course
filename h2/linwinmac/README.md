# Harjoitus 2a)

Harjoituksessa käännettiin yksinkertainen go-ohjelma kolmelle alustalle (Linux, Windows, Mac).

Käännöksen tekevät komennot
```
CGO_ENABLED=0 GOARCH=amd64 GOOS=windows go build
CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build
CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build
```

 Alla on kuvankaappaus tilanteesta, jossa linuxilla tehty staattinen käännös toimii Windows-alustalla.

![kuva](linwinmac.png)