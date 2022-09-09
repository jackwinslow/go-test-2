# go-test-2
Please enter go run main.go to run the program. The first function executes as a part of Question 1, and then right after the second function executes as a part of Question 2.

Question 3 Answer: I believe my code does not resemble the theory about big O mentioned in the Golang Docs, since SliceStable makes O(nlog(n)log(n)) to data.Swap instead of O(nlog(n)) like in Slice. This is because it is not altering the order of equal duplicates which naturally saves time. O(nlog(n)log(n)) < O(nlog(n)), but my code takes longer to run SliceStable than to run Slice, contradicting the theory mentioned in the Golang Docs. 
