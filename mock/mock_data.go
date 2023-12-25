package mock

type TestTree struct {
	Name    string
	Surname string
}

func GetTreeMockData() []*TestTree {
	return []*TestTree{
		generateFakeData("eRay", "can"),
		generateFakeData("john", "doe"),
		generateFakeData("Bob", "Johnson"),
		generateFakeData("Eva", "Williams"),
		generateFakeData("Michael", "Brown"),
		generateFakeData("cristiano", "rONAldO"),
		generateFakeData("Messi", "surname"),
		generateFakeData("Talisca", "fake"),
		generateFakeData("Ndombele", "Brown"),
		generateFakeData("MicCael", "Brown"),
		generateFakeData("aliCe", "smith"),
	}

}

func generateFakeData(name, surname string) *TestTree {
	return &TestTree{
		Name:    name,
		Surname: surname,
	}
}
