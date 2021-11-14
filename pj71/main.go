package main

type Entry struct {
	Name  string
	Value int
}

type List []Entry

func (l List) Less(i, j int) bool {
	if l[i].Value == l[j].Value {
		return (l[i].Name < l[j].Name)
	} else {
		return (l[i].Value < l[j].Value)
	}
}

func main() {
	m := map[string]int{"S": 1, "J": 4, "A": 3, "N": 3}
	lt := List{}
	for k, v := range m {
		e := Entry{k, v}
		lt = append(lt, e)
	}

}
