package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type Person struct {
	Name string
	Age int
}

type User struct {
	Name   string
	Age    int
	Salary int
}

type By func(p1, p2 *Person) bool

type personSorter struct {
	persons []Person
	by By
}

func (s *personSorter) Len() int {
	return len(s.persons)
}

func (s *personSorter) Less(i, j int) bool {
	return s.by(&s.persons[i], &s.persons[j])
}

func (s *personSorter) Swap(i, j int) {
	s.persons[i], s.persons[j] = s.persons[j], s.persons[i]
}

func (by By) Sort(persons []Person) {
	ps := &personSorter{	
		persons: persons,
		by: by,
	}

	sort.Sort(ps)
}


func main(){
	people := []Person{
		{Name: "Alice", Age: 20},
		{Name: "Bob", Age: 25},
		{Name: "Anna", Age: 20},
	}

	ageAsc := func (p1, p2 *Person) bool {
		return p1.Age < p2.Age
	}

	ageDesc := func (p1, p2 *Person) bool {
		return p1.Age > p2.Age
	}

	nameAsc := func (p1, p2 * Person) bool {
		return p1.Name < p2.Name
	}

	lenNameAsc := func (p1, p2 * Person) bool {
		return len(p1.Name) < len(p2.Name)
	}

	By(ageAsc).Sort(people)
	fmt.Println("Sort ASC by age:", people)
	By(nameAsc).Sort(people)
	fmt.Println("Sort ASC by name:", people)
	By(ageDesc).Sort(people)
	fmt.Println("Sort DESC by age:", people)
	By(lenNameAsc).Sort(people)
	fmt.Println("Sort ASC by length name:", people)

	// SORT.SLICE
	stringSlice := []string{"banana", "apple", "cherry", "grapes", "guzava"}
	sort.Slice(stringSlice, func(i,j int)bool {
		return stringSlice[i][0] < stringSlice[j][0]
	})

	fmt.Println("Sorted stringSlice:", stringSlice)

	//1.1
	sort.Slice(stringSlice, func(i, j int) bool {
		return len(stringSlice[i]) < len(stringSlice[j])
	})
	fmt.Println("Sorted stringSlice by len:", stringSlice)

	//1.2
	sort.Slice(stringSlice, func(i, j int) bool {
		return stringSlice[i][len(stringSlice[i])-1] < stringSlice[j][len(stringSlice[j])-1]
	})
	fmt.Println("Sorted stringSlice by last char:", stringSlice)

	//1.3
	sort.Slice(stringSlice, func(i, j int) bool {
		return countVowels(stringSlice[i]) < countVowels(stringSlice[j])
	})
	fmt.Println("Sorted stringSlice by vowels count:", stringSlice)

	//2.1
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sort ASC by age:", people)

	//2.2
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age > people[j].Age
	})
	fmt.Println("Sort DESC by age:", people)

	//2.3
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Println("Sort ASC by name:", people)

	//2.4
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name > people[j].Name
	})
	fmt.Println("Sort DESC by name:", people)

	//3
	sort.Slice(stringSlice, func(i, j int) bool {
		lower1 := strings.ToLower(stringSlice[i])
		lower2 := strings.ToLower(stringSlice[j])

		s1HasPrefix := strings.HasPrefix(lower1, "a")
		s2HasPrefix := strings.HasPrefix(lower2, "a")

		if s1HasPrefix != s2HasPrefix {
			return s1HasPrefix
		}

		if len(lower1) != len(lower2) {
			return len(lower1) < len(lower2)
		}

		return lower1 < lower2
	})
	fmt.Println("Sort by name prefix 'a' and len:", stringSlice)

	//4
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Stable sort of people:", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Unstable sort of people:", people)

	users := []User{
		{Name: "Alex", Age: 15, Salary: 400},
		{Name: "Bob", Age: 25, Salary: 4000},
		{Name: "Mark", Age: 14, Salary: 2400},
		{Name: "Rauf", Age: 19, Salary: 200},
		{Name: "Olga", Age: 40, Salary: 1000},
		{Name: "Alisa", Age: 24, Salary: 5400},
		{Name: "Stiv", Age: 58, Salary: 800},
	} 

	// Unstable Salary Desc -> Age ASC -> Name ASC
	sort.Slice(users, func(i, j int) bool {
		if users[i].Salary != users[j].Salary {
			return users[i].Salary > users[j].Salary
		}

		if users[i].Age != users[j].Age {
			return users[i].Age < users[j].Age
		}

		return users[i].Name < users[j].Name
	})
	fmt.Println("Users unstable sort:", users)

	// Stable Name ASC -> Age ASC -> Salary DESC
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Name < users[j].Name	
	})
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Age < users[j].Age
	})
	sort.SliceStable(users, func(i, j int) bool {
		return users[i].Salary > users[j].Salary
	})
	fmt.Println("Users stable sort:", users)

	nums := []int64{1, 5, 2, 3, 1}
	sort.Slice(nums , func(i, j int) bool {
		return nums[i]%3 < nums[j]%3
	})
	fmt.Println("Bad comparator sort:", nums)

	sort.Slice(nums, func(i, j int) bool {
		return rand.Intn(2) == 0
	})

	words := []string{
		"яблоко",
		"арбуз",
		"банан",
		"ёж",
		"ель",
	}

	sort.Slice(words, func(i, j int) bool{
		return []rune(words[i])[0] < []rune(words[j])[0]
	})

	fmt.Println("Words sorted:", words)

	// SORT.SEARCH
	nums1 := []int{1,3,5,7,9,11}

	idx := sort.Search(len(nums1), func(i int) bool {
		return nums1[i] >= 6 
	})
	fmt.Println("Idx:", nums1[idx])

	idx = sort.Search(len(nums1), func(i int) bool {
		return nums1[i] >= 7
	})

	if idx == len(nums1) && nums1[idx-1] != 7 {
		fmt.Println("7 is not in the nums")
	} else {
		fmt.Println("7 is in the nums")
	}

	idx = sort.Search(len(nums1), func(i int) bool {
		return nums1[i] >= 8
	})

	fmt.Println("Insertion point:", idx)

}


func countVowels(str string) int {
	vowels := "aeiouAEIOU"
	count := 0

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count += 1
		}
	}
	return count
}