package double

import "testing"

type SpySearcher struct {
	phone           string
	searchWasCalled bool
	whatIsFirstName string
}

func (ss *SpySearcher) Search(people []*Person, firstName, lastName string) *Person {
	ss.searchWasCalled = true
	ss.whatIsFirstName = firstName
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}

func TestFindCallsSearchAndReturnsPerson(t *testing.T) {
	fakePhone := "+31 65 222 333"
	phonebook := &Phonebook{}
	spy := &SpySearcher{phone: fakePhone}

	phone, _ := phonebook.Find(spy, "Jane", "Doe")

	if !spy.searchWasCalled {
		t.Errorf("Expected to call 'Search' in 'Find', but it wasn't.")
	}

	if spy.whatIsFirstName != "ANUCHIT" {
		t.Error("xxxx")
	}

	if phone != fakePhone {
		t.Errorf("Want '%s', got '%s'", fakePhone, phone)
	}
}
