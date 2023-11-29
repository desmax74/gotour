//go:build OMIT

// Όλα τα υλικά είναι αδειοδοτημένα υπό την Άδεια Apache Έκδοση 2.0, Ιανουάριος 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Δείγμα προγράμματος, που παρουσιάζει πως συμπεριφέρονται οι σχεσιακοί πίνακες,
// όταν γίνεται ανάγνωση ενός κλειδιού, που δεν υπάρχει.
package main

import "fmt"

func main() {

	// Δημιουργείστε έναν σχεσιακό πίνακα, προκειμένου να παρακολουθήσετε
	// τα σκορ παικτών, σε έναν αγώνα.
	scores := make(map[string]int)

	// Διαβάστε το στοιχείο στο κλειδί "anna". Δεν υπάρχει, επομένως λαμβάνουμε
	// την μηδενική τιμή για τον τύπο των τιμών, του σχεσιακού πίνακα.
	score := scores["anna"]

	fmt.Println("Score:", score)

	// Αν πρέπει να ελέγξουμε για την παρουσία ενός κλειδιού, χρησιμοποιούμε
	// μια εκχώρηση 2 μεταβλητών. Η 2η μεταβλητή είναι μια bool.
	score, ok := scores["anna"]

	fmt.Println("Score:", score, "Present:", ok)

	// Μπορούμε να χρησιμοποιήσουμε την συμπεριφορά με την μηδενική τιμή, ώστε
	// να γράψουμε βολικο κώδικα, όπως αυτός:
	scores["anna"]++

	// Χωρίς αυτή την συμπεριφορά, θα έπρεπε να γράψουμε κώδικα
	// με αμυντικό τρόπο, όπως παρακάτω:
	if n, ok := scores["anna"]; ok {
		scores["anna"] = n + 1
	} else {
		scores["anna"] = 1
	}

	score, ok = scores["anna"]
	fmt.Println("Score:", score, "Present:", ok)
}