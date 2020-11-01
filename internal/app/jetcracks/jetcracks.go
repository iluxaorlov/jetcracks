package jetcracks

import (
	"fmt"
	"io/ioutil"
	"os"
)

const baseDir = ".config/JetBrains"

func Start(homeDir string) error {
	dir := homeDir + "/" + baseDir

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("There is no products")
			return nil
		}

		return err
	}

	names := make(map[int]string)
	point := 1

	message := "Select JetBrains product:\n"
	format := "%d) %s\n"

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		names[point] = file.Name()
		message += fmt.Sprintf(format, point, file.Name())
		point += 1
	}

	message += fmt.Sprintf(format, point, "Select all")
	fmt.Print(message)

	var selected int

	_, err = fmt.Fscan(os.Stdin, &selected)
	if err != nil {
		return err
	}

	if selected == point {
		for _, name := range names {
			productDir := dir + "/" + name

			if err := clear(productDir); err != nil {
				return err
			}
		}
	} else {
		name, ok := names[selected]
		if !ok {
			fmt.Println("No such product")
			return nil
		}

		productDir := dir + "/" + name

		if err := clear(productDir); err != nil {
			return err
		}
	}

	if err := os.RemoveAll(homeDir + "/.java/.userPrefs"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	fmt.Println("Success")

	return nil
}

func clear(productDir string) error {
	if err := os.RemoveAll(productDir + "/options/other.xml"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	if err := os.RemoveAll(productDir + "/eval"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}

	return nil
}
