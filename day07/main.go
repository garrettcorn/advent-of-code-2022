package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type fileSystem struct {
	root *dir
}

func (f fileSystem) String() string {
	return fmt.Sprintf("%v", f.root)
}

func (f fileSystem) DirsTotalSizeAtMost(size int) []dir {
	return f.root.MaxSizeDirs(size)
}

func (f fileSystem) DirsTotalSizeAtLeast(size int) []dir {
	return f.root.MinSizeDirs(size)
}

type file struct {
	name string
	size int
}

func (f file) String() string {
	return fmt.Sprintf("\t%v (file, size=%v)", f.name, f.size)
}

type dir struct {
	parentDir    *dir
	name         string
	childrenDirs []dir
	files        []file
}

func (d dir) MaxSizeDirs(size int) []dir {
	dirs := []dir{}
	if d.Size() <= size {
		dirs = append(dirs, d)
	}
	for _, cdr := range d.childrenDirs {
		dirs = append(dirs, cdr.MaxSizeDirs(size)...)
	}
	return dirs
}

func (d dir) MinSizeDirs(size int) []dir {
	dirs := []dir{}
	if d.Size() >= size {
		dirs = append(dirs, d)
	}
	for _, cdr := range d.childrenDirs {
		dirs = append(dirs, cdr.MinSizeDirs(size)...)
	}
	return dirs
}

func (d dir) String() string {
	result := fmt.Sprintf("%v (dir, size=%v)\n", d.name, d.Size())
	for _, f := range d.files {
		result += fmt.Sprintf("%v\n", f)
	}
	for _, cdr := range d.childrenDirs {
		for _, l := range strings.Split(fmt.Sprintf("%v", cdr), "\n") {
			if l != "" {
				result += fmt.Sprintf("\t%v\n", l)
			}
		}
	}
	return result
}

func (d dir) Size() int {
	sum := 0
	for _, f := range d.files {
		sum += f.size
	}
	for _, cds := range d.childrenDirs {
		sum += cds.Size()
	}
	return sum
}

func getFileSystemFromInput(input io.Reader) fileSystem {
	s := bufio.NewScanner(input)
	s.Split(bufio.ScanLines)

	fsys := fileSystem{}
	currentDir := &dir{parentDir: nil, childrenDirs: []dir{}, files: []file{}}
	for s.Scan() {
		line := s.Text()
		xpart := strings.Split(line, " ")
		// for idx, val := range xpart {
		// 	fmt.Printf("xpart[%v]=%v\n", idx, val)
		// }
		// fmt.Println()
		if xpart[0] == "$" {
			// command
			if xpart[1] == "cd" {
				// cd
				switch xpart[2] {
				case "/":
					currentDir.name = "/"
					fsys.root = currentDir
				case "..":
					currentDir = currentDir.parentDir
				default:
					// cd to a child dir
					for idx, d := range currentDir.childrenDirs {
						if xpart[2] == d.name {
							currentDir = &currentDir.childrenDirs[idx]
						}
					}
				}
			} else if xpart[1] == "ls" {
				// nothing to do for ls
			}
		} else if xpart[0] == "dir" {
			currentDir.childrenDirs = append(currentDir.childrenDirs, dir{name: xpart[1], parentDir: currentDir})
		} else {
			// file
			s, err := strconv.Atoi(xpart[0])
			if err != nil {
				panic(err)
			}
			currentDir.files = append(currentDir.files, file{name: xpart[1], size: s})
		}
	}
	return fsys
}

func readFile(filePath string) io.Reader {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	// defer f.Close()

	return f
}

func main() {
	filePath := "./input"
	fileContent := readFile(filePath)
	fileSystem := getFileSystemFromInput(fileContent)
	// fmt.Printf("%v\n", fileSystem)
	size := 100000
	dirs := fileSystem.DirsTotalSizeAtMost(size)
	sumOfDirs := sumDirSizes(dirs)
	fmt.Printf("PART1: %v\n", sumOfDirs)

	totalSpace := 70000000
	spaceNeededForUpdate := 30000000
	unusedSpace := totalSpace - fileSystem.root.Size()
	spaceNeeded := spaceNeededForUpdate - unusedSpace
	delCanidates := fileSystem.DirsTotalSizeAtLeast(spaceNeeded)
	sDir := smallestDir(delCanidates)
	fmt.Printf("PART2: %v\n", sDir.Size())

}

func smallestDir(dirs []dir) dir {
	result := dirs[0]
	for _, d := range dirs {
		if d.Size() < result.Size() {
			result = d
		}
	}
	return result
}

func sumDirSizes(dirs []dir) int {
	sum := 0
	for _, d := range dirs {
		sum += d.Size()
	}
	return sum
}
