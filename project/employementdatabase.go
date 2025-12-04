package project

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type employee struct {
	ID     int
	Name   string
	Age    int
	Salary int
}

type repo interface {
	Add(e employee) (int, error)
	Getall() []employee
	SearchByName(name string) []employee
	Searchbyid(id int) (employee, bool)
	updateSalary(id int, Newsalary int) error
	Deleteid(id int) error
	Save(filename string) error
	Load(filename string) error
}

type Inmemoryrepo struct {
	mu     sync.Mutex
	store  map[int]employee
	order  []int
	nextID int
}

func NewInmemoryrepo() *Inmemoryrepo {
	return &Inmemoryrepo{
		store:  make(map[int]employee),
		order:  []int{},
		nextID: 1,
	}
}

func (r *Inmemoryrepo) Add(e employee) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	id := r.nextID
	e.ID = id
	r.store[id] = e
	r.order = append(r.order, id)
	r.nextID++
	return id, nil
}
func (r *Inmemoryrepo) Getall() []employee {
	r.mu.Lock()
	defer r.mu.Unlock()
	emps := make([]employee, 0, len(r.order))
	for _, id := range r.order {
		if e, ok := r.store[id]; ok {
			emps = append(emps, e)
		}
	}
	return emps
}

func (r *Inmemoryrepo) SearchByName(name string) []employee {
	r.mu.Lock()
	defer r.mu.Unlock()
	name = strings.ToLower(strings.TrimSpace(name))
	results := make([]employee, 0)
	for _, e := range r.store {
		if strings.Contains(strings.ToLower(e.Name), name) {
			results = append(results, e)
		}
	}
	return results

}

func (r *Inmemoryrepo) Searchbyid(id int) (employee, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	e, ok := r.store[id]
	return e, ok
}

func (r *Inmemoryrepo) updateSalary(id int, Newsalary int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	e, ok := r.store[id]
	if !ok {
		return errors.New("Employee not found")
	}
	e.Salary = Newsalary
	r.store[id] = e
	return nil
}
func (r *Inmemoryrepo) Deleteid(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.store[id]; !ok {
		return errors.New("Employee not found")
	}
	delete(r.store, id)
	for i, v := range r.order {
		if v == id {
			r.order = append(r.order[i:], r.order[i+1:]...)
			break
		}
	}
	return nil
}

var (
	Employees   []employee
	employeemap = make(map[int]employee)
)

func saveemployees(filename string) error {
	data, err := json.MarshalIndent(Employees, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
func Loademployees(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &Employees)
	if err != nil {
		return err
	}
	return nil
}
func Addemployee(id int, name string, age int, position string, salary int) {
	emp := employee{
		ID:     id,
		Name:   name,
		Age:    age,
		Salary: salary,
	}
	Employees = append(Employees, emp)

}
func Viewemployee() {
	for _, e := range Employees {
		fmt.Printf("ID :%d | Name :%s | Age:%d | Salary:%d\n", e.ID, e.Name, e.Age, e.Salary)
	}
}
func Searchemployeebyname(name string) {
	found := false
	for _, e := range Employees {
		if strings.EqualFold(e.Name, name) {
			fmt.Printf("ID :%d | Name :%s | Age:%d | Salary:%d\n", e.ID, e.Name, e.Age, e.Salary)
			found = true
		}
	}
	if !found {
		fmt.Println("please enter a valid name")
	}
}

func Updatesalary(id int, newsalary int) {
	updated := false
	for i, e := range Employees {
		if e.ID == id {
			Employees[i].Salary = newsalary
			employeemap[i] = Employees[i]
			fmt.Printf("ID :%d  Newsalary :%d\n", id, newsalary)
			updated = true
			break
		}
	}
	if !updated {
		fmt.Println("salary not updated")
	}

}

func Employeeproject() {
	Loademployees("employees.json")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("__employee management system__")
		fmt.Println("1.Add employee")
		fmt.Println("2.List employee")
		fmt.Println("3.search by name")
		fmt.Println("4.update salary")
		fmt.Println("5.Delete employee")
		fmt.Println("6.save and exit")

		fmt.Println("Enter choice:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {

		case "1":
			fmt.Println("Enter Id:")
			idstr, _ := reader.ReadString('\n')
			idstr = strings.TrimSpace(idstr)
			id, _ := strconv.Atoi(idstr)

			fmt.Println("Enter name:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("Enter Age:")
			agestr, _ := reader.ReadString('\n')
			agestr = strings.TrimSpace(agestr)
			age, _ := strconv.Atoi(agestr)

			fmt.Println("Enter salary:")
			salarystr, _ := reader.ReadString('\n')
			salarystr = strings.TrimSpace(salarystr)
			salary, _ := strconv.Atoi(salarystr)

			Addemployee(id, name, age, salary)
			fmt.Println("Employee added succesfully")

		case "2":
			fmt.Println("view all employee")
			Viewemployee()

		case "3":

			fmt.Println("Enter employee name to search")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			Searchemployeebyname(name)

		case "4":
			fmt.Println("Enter employee ID to update salary:")
			idstr, _ := reader.ReadString('\n')
			idstr = strings.TrimSpace(idstr)
			id, _ := strconv.Atoi(idstr)

			fmt.Println("Enter new salary:")
			salarystr, _ := reader.ReadString('\n')
			salarystr = strings.TrimSpace(salarystr)
			newsalary, _ := strconv.Atoi(salarystr)

			Updatesalary(id, newsalary)

		case "5":
			saveemployees("employees.json")
			fmt.Println("employees saved succesfully.")
			return
		default:
			fmt.Println("please enter a valid choice.")

		}
	}
}
