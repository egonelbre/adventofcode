package main

import "fmt"

type Password [8]byte

func Pass(text string) *Password {
	pass := &Password{}
	if len(text) != len(pass) {
		panic("invalid length")
	}
	for i := range pass {
		pass[i] = text[i]
	}
	return pass
}

func (password *Password) Inc() *Password {
	for i := len(password) - 1; i >= 0; i-- {
		password[i]++
		if password[i] > 'z' {
			password[i] = 'a'
			continue
		}
		break
	}
	return password
}

func (password *Password) IncAt(pos int) *Password {
	for i := pos + 1; i < len(password); i++ {
		password[i] = 'a'
	}
	for i := pos; i >= 0; i-- {
		password[i]++
		if password[i] >= 'z' {
			password[i] = 'a'
			continue
		}
		break
	}
	return password
}

func (password *Password) String() string { return string(password[:]) }

func Valid(password *Password) (bool, int) {
	for i, r := range password {
		if r == 'i' || r == 'o' || r == 'l' {
			return false, i
		}
	}

	doubles, skip := 0, 0
	for i, r := range password[1:] {
		if skip > 0 {
			skip--
			continue
		}
		if password[i] == r {
			doubles++
			skip = 1
		}
	}
	if doubles < 2 {
		return false, -1
	}

	seq := false
	for i, r := range password[2:] {
		if password[i]+2 == r && password[i+1]+1 == r {
			seq = true
			break
		}
	}
	if !seq {
		return false, -1
	}

	return true, -1
}

func NextValid(start *Password) *Password {
	t := *start
	cursor := &t
	for {
		valid, at := Valid(cursor)
		if valid {
			return cursor
		}
		if at >= 0 {
			cursor.IncAt(at)
		} else {
			cursor.Inc()
		}
	}
}

func NextValid2(start *Password) *Password {
	t := *start
	cursor := &t
	for {
		valid, _ := Valid(cursor)
		if valid {
			return cursor
		}
		cursor.Inc()
	}
}

func IsValid(password *Password) bool {
	ok, _ := Valid(password)
	return ok
}

func main() {
	fmt.Println(Pass("aaaaaaaa"))
	fmt.Println(Pass("aaaaaaaz"))
	fmt.Println(Pass("hijklmmn").Inc())

	fmt.Println("---")
	fmt.Println(IsValid(Pass("aaaaefgh")))
	fmt.Println(IsValid(Pass("hxbxxyzz")))
	fmt.Println("---")
	fmt.Println(IsValid(Pass("hijklmmn")), NextValid(Pass("hijklmmn")))
	fmt.Println(IsValid(Pass("abbceffg")), NextValid(Pass("abbceffg")))
	fmt.Println(IsValid(Pass("abbcegjk")), NextValid(Pass("abbcegjk")))
	fmt.Println(NextValid(Pass("abcdefgh")), "abcdffaa")
	fmt.Println(NextValid(Pass("ghijklmn")), "ghjaabcc")

	fmt.Println("--------------------")
	next := NextValid(Pass("hxbxwxba"))
	fmt.Println(next)
	fmt.Println(NextValid(next.Inc()))
}
