package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	input, err := os.ReadFile("input1.txt")
	if err != nil {
		return err
	}

	cmds, err := parseCommands(string(input))
	if err != nil {
		return err
	}

	root := reconstruct(cmds)
	updateDirSize(root)
	fmt.Println("total size", root.Size)

	total := 0
	iterateDirs(root, func(fi *FileInfo) {
		if fi.Size < 100000 {
			total += fi.Size
		}
	})
	fmt.Println("total of small", total)

	// printTree(root, "")
	var best *FileInfo
	totalSpace := 70000000
	requiredFree := 30000000
	missingFree := requiredFree - (totalSpace - root.Size)

	iterateDirs(root, func(fi *FileInfo) {
		if fi.Size > missingFree {
			if best == nil {
				best = fi
			} else {
				if best.Size > fi.Size {
					best = fi
				}
			}
		}
	})
	fmt.Println("best", best.Full(), best.Size)

	return nil
}

func printTree(node *FileInfo, prefix string) {
	fmt.Printf("%s- %s %s\n", prefix, node.Name, node.OriginalStat())
	for _, c := range node.Content {
		printTree(c, prefix+"  ")
	}
}

func updateDirSize(root *FileInfo) {
	var total int
	for _, n := range root.Content {
		if n.IsDir {
			updateDirSize(n)
		}
		total += n.Size
	}
	root.Size = total
}

func iterateDirs(root *FileInfo, fn func(*FileInfo)) {
	for _, c := range root.Content {
		if c.IsDir {
			fn(c)
			iterateDirs(c, fn)
		}
	}
}

func reconstruct(cmds Commands) *FileInfo {
	root := &FileInfo{Name: "/"}
	root.Parent = root

	workingFile := root

	var find func(node *FileInfo, name string) *FileInfo
	find = func(node *FileInfo, name string) *FileInfo {
		if name == "" {
			return node
		}

		dir, rest, _ := strings.Cut(name, "/")
		for _, f := range node.Content {
			if f.Name == dir {
				return find(f, rest)
			}
		}

		printTree(node, "")
		panic("did not find " + name)
	}

	for _, cmd := range cmds {
		switch cmd := cmd.(type) {
		case CD:
			switch {
			case cmd.Target == "..":
				workingFile = workingFile.Parent
			case strings.HasPrefix(cmd.Target, "/"):
				workingFile = find(root, cmd.Target[1:])
			default:
				workingFile = find(workingFile, cmd.Target)
			}
		case LS:
			// TODO: verify that there are no duplicates
			for _, f := range cmd.Files {
				f.Parent = workingFile
			}
			workingFile.Content = append(workingFile.Content, cmd.Files...)
		default:
			panic(fmt.Sprintf("unhandled %#v", cmd))
		}
	}

	return root
}

type Commands []Command

type Command any

func parseCommands(output string) (Commands, error) {
	rd := bufio.NewReader(strings.NewReader(output))
	done := false
	for !done {
		line, err := rd.ReadString('\n')
		done = errors.Is(err, io.EOF)
		if line == "" {
			continue
		}
	}

	cmds := Commands{}
	rxCommands := regexp.MustCompile(`(?m)\$ *`)
	for _, cmdout := range rxCommands.Split(output, -1) {
		if cmdout == "" {
			continue
		}
		cmd, err := parseCommand(cmdout)
		if err != nil {
			return cmds, fmt.Errorf("failed to parse %q: %w", cmdout, err)
		}
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}

var rxCommand = regexp.MustCompile(`(?s)([a-z]+) *([^\n]*)([^$]*)`)

func parseCommand(output string) (Command, error) {
	matches := rxCommand.FindStringSubmatch(output)
	if len(matches) == 0 {
		return nil, fmt.Errorf("invalid output %q", output)
	}

	cmd := CommandOutput{
		Command: strings.TrimSpace(matches[1]),
		Args:    strings.TrimSpace(matches[2]),
		Output:  strings.TrimSpace(matches[3]),
	}

	parse, ok := commandParse[cmd.Command]
	if !ok {
		return nil, fmt.Errorf("unknown command %q", cmd.Command)
	}

	return parse(cmd)
}

type CommandOutput struct {
	Command string
	Args    string
	Output  string
}

var commandParse = map[string]func(out CommandOutput) (Command, error){
	"cd": parseCD,
	"ls": parseLS,
}

type CD struct {
	Target string
}

func parseCD(out CommandOutput) (Command, error) {
	return CD{
		Target: out.Args,
	}, nil
}

type LS struct {
	Files []*FileInfo
}

type FileInfo struct {
	Name  string
	IsDir bool
	Size  int

	Parent  *FileInfo
	Content []*FileInfo
}

func (info *FileInfo) Full() string {
	if info.Parent == nil || info == info.Parent {
		return ""
	}
	return info.Parent.Full() + "/" + info.Name
}

func (info *FileInfo) OriginalStat() string {
	if info.IsDir {
		return fmt.Sprintf("(dir, size=%d)", info.Size)
	} else {
		return fmt.Sprintf("(file, size=%d)", info.Size)
	}
}

func parseLS(out CommandOutput) (Command, error) {
	ls := LS{}
	for _, line := range strings.Split(out.Output, "\n") {
		if line == "" {
			continue
		}
		num, name, _ := strings.Cut(line, " ")
		if num == "dir" {
			ls.Files = append(ls.Files, &FileInfo{
				Name:  name,
				IsDir: true,
			})
		} else {
			size, err := strconv.Atoi(num)
			if err != nil {
				return ls, fmt.Errorf("%q %q: %w", line, num, err)
			}
			ls.Files = append(ls.Files, &FileInfo{
				Name: name,
				Size: size,
			})
		}
	}
	return ls, nil
}
