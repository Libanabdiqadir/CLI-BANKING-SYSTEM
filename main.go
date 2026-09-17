package main

func main() {
	store := NewMemoryAccountStore()

	RunCLI(store)
}