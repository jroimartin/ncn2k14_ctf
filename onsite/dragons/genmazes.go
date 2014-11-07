package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"text/template"
)

func main() {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	ciphertext := make([]byte, aes.BlockSize+len(mazes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], mazes)
	mazesTemplate.Execute(os.Stdout, ciphertext)
}

const key = "\x39\x82\x50\x85\x9f\x30\x31\xde\xca\xf8\x4f\x0b\xf5\x74\x7a\x25\xfa" +
	"\xb4\x7b\x0e\x76\xe9\xb7\xfb\x60\xf8\xac\xad\x7b\x83\xd8\xe9"

var mazesTemplate = template.Must(template.New("").Parse(mazesCode))

const mazesCode = `package main

import (
	"crypto/aes"
	"crypto/cipher"
)

func generateMaze() (m *maze) {
	m = new(maze)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	iv := mazes[:aes.BlockSize]
	ciphertext := mazes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	for i, ch := range ciphertext {
		l := i / (levelSizeX*levelSizeY)
		x := (i % (levelSizeX*levelSizeY)) / levelSizeY
		y := (i % (levelSizeX*levelSizeY)) % levelSizeY
		m[l][x][y] = mazeEntity(ch)
	}
	return m
}

var mazes = []byte{ {{range .}}{{.}},{{end}} }

const key = "\x39\x82\x50\x85\x9f\x30\x31\xde\xca\xf8\x4f\x0b\xf5\x74\x7a\x25\xfa" +
	"\xb4\x7b\x0e\x76\xe9\xb7\xfb\x60\xf8\xac\xad\x7b\x83\xd8\xe9"
`

var mazes = []byte{
		// level 0
		' ', '#', ' ', ' ', ' ', '#', '#', '#', '#', '#', '#', '#', 'u', '#', '#', '#',
		' ', '#', ' ', '#', ' ', '#', '#', ' ', ' ', ' ', '#', '#', ' ', '#', '#', '#',
		' ', '#', ' ', '#', 'm', '#', '#', ' ', '#', ' ', '#', '#', ' ', '#', '#', '#',
		' ', '#', ' ', '#', '#', '#', '#', ' ', '#', ' ', '#', '#', ' ', 'm', ' ', '#',
		' ', '#', ' ', '#', ' ', '#', '#', 'm', '#', ' ', '#', '#', ' ', '#', '#', '#',
		' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', '#', '#', ' ', '#', '#', '#',
		'#', '#', '#', '#', ' ', ' ', 'm', '#', '#', 'm', '#', '#', ' ', '#', '#', '#',
		' ', ' ', ' ', ' ', ' ', '#', '#', '#', '#', ' ', '#', '#', ' ', ' ', ' ', '#',
		' ', '#', '#', '#', ' ', '#', '#', '#', '#', ' ', '#', '#', ' ', '#', ' ', '#',
		' ', '#', '#', '#', ' ', '#', '#', '#', '#', ' ', '#', '#', ' ', '#', ' ', '#',
		' ', '#', '#', '#', ' ', '#', '#', '#', '#', ' ', '#', '#', '#', '#', ' ', '#',
		' ', '#', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'm', ' ', ' ', ' ', ' ',
		' ', '#', '#', '#', 'm', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' ',
		' ', '#', '#', '#', ' ', '#', ' ', 'f', '#', '#', '#', '#', '#', '#', '#', ' ',
		'm', '#', '#', '#', ' ', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', 'm',
		'#', '#', '#', '#', ' ', ' ', ' ', 'm', '#', ' ', '#', '#', '#', '#', '#', '#',
		// level 1
		'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', ' ', '#', '#', '#',
		'm', '#', '#', 'm', '#', '#', '#', '#', '#', '#', '#', '#', ' ', '#', '#', '#',
		' ', '#', '#', ' ', '#', '#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', '#', '#',
		' ', '#', '#', ' ', '#', '#', ' ', '#', '#', '#', '#', '#', ' ', '#', '#', '#',
		' ', '#', '#', ' ', '#', '#', ' ', '#', ' ', 'm', ' ', '#', ' ', '#', '#', '#',
		' ', '#', '#', ' ', '#', '#', ' ', '#', ' ', ' ', 'm', '#', ' ', '#', '#', '#',
		' ', '#', '#', ' ', ' ', ' ', ' ', '#', 'm', ' ', ' ', '#', ' ', '#', '#', '#',
		' ', '#', '#', ' ', '#', '#', '#', '#', ' ', ' ', ' ', '#', ' ', '#', '#', '#',
		' ', ' ', '#', ' ', '#', '#', '#', '#', ' ', ' ', 'm', '#', ' ', '#', '#', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#', ' ', '#', '#', '#',
		'#', ' ', '#', '#', '#', '#', ' ', '#', ' ', ' ', ' ', '#', ' ', '#', '#', '#',
		'#', ' ', '#', '#', 'm', '#', ' ', '#', ' ', 'm', ' ', '#', ' ', '#', '#', '#',
		'#', ' ', '#', '#', ' ', '#', ' ', '#', ' ', ' ', 'm', '#', ' ', '#', '#', '#',
		'#', ' ', '#', '#', ' ', '#', 'd', '#', ' ', ' ', ' ', '#', ' ', '#', '#', '#',
		'#', ' ', '#', '#', ' ', '#', '#', '#', '#', '#', '#', '#', ' ', '#', '#', '#',
		'#', 'm', ' ', ' ', ' ', ' ', ' ', 'm', '#', 'd', ' ', ' ', ' ', '#', '#', '#',
}
