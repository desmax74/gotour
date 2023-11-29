//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Από τις Προδιαγραφές της Γλώσσας:
// μια γρήγορη δήλωση μεταβλητής μπορεί να δηλώσει ξανά μεταβλητές,
// δεδομένου ότι είχαν αρχικά δηλωθεί νωρίτερα, στο ίδιο τμήμα κώδικα,
// με τον ίδιο τύπο και τουλάχιστον μια από τις μεταβλητές που δεν
// είναι κενές (_), είναι καινούργια.

// Δείγμα προγράμματος, προκειμένου να παρουσιαστούν ορισμένοι από τους μηχανισμούς
// πίσω από τον τελεστή σύντομης δήλωσης, σε περιπτώσεις όπου γίνεται εκ νέου,
// σύντομη δήλωση.
package main

import "fmt"

// Ο user είναι ένας τύπος struct, που δηλώνει πληροφορίες χρήστη.
type user struct {
	id   int
	name string
}

func main() {

	// Δηλώστε την μεταβλητή σφάλματος.
	var err1 error

	// Ο τελεστής σύντομης δήλωσης μεταβλητής, θα δηλώσει
	// την u και θα ξαναδηλώσει την err1.
	u, err1 := getUser()
	if err1 != nil {
		return
	}

	fmt.Println(u)

	// Ο τελεστής σύντομης δήλωσης μεταβλητής, θα
	// ξαναδηλώσει την u και θα δηλώσει την err2.
	u, err2 := getUser()
	if err2 != nil {
		return
	}

	fmt.Println(u)
}

// Η getUser επιστρέφει έναν δείκτη διεύθυνσης τύπου user.
func getUser() (*user, error) {
	return &user{1432, "Betty"}, nil
}