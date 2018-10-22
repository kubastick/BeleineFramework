package beleine

import (
	"strings"
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func ReadSimplify(dir string) {

	if !strings.HasSuffix(dir, ".beleine") {
		fmt.Println("File name must end with .beleine")
		return
	}

	file, err := os.Open(dir)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := bufio.NewScanner(file)

	var page Page

	for data.Scan() {

		//TODO: Split this shitty code into functions or something like this idk

		if strings.HasPrefix(data.Text(),"(Label") {
			label := NewLabel()

			sizeIndex := strings.Index(data.Text(), "size")
			endValueIndedx := strings.Index(data.Text(), ")")

			if endValueIndedx != -1 {
				label.SetText(data.Text()[endValueIndedx + 1:])
			} else {
				panic(`There is no ")" ending in arguments section.`)
			}

			if sizeIndex != -1 {
				size := data.Text()[sizeIndex + 5 : sizeIndex + 6]

				i, err := strconv.Atoi(size)

				if err != nil {
					panic(err)
				}

				label.SetSize(i)
			}

			page.Attach(&label)
		}

		if err != nil {
			break
		}
	}

	fmt.Print(page.Render())
}