# seatgeek

My initial approach was to boil section names down to just the number using
regular expressions. I was excited about my initial success was with Citi Field
and was quickly disappointed in Dodger Stadium for their reuse of so many
numbers. My next approach was to scale back my efforts and try to reduce section
names to their actual names. At this point, I realized that more tickets had
too little data than too much data. If I was unable to map a ticket to a section
I tried the reverse; mapping sections to tickets. This was a bad idea as I had
two approaches to manage. My final solution was to combine my bottom-up and
top-down approaches into a single top-down approach by applying a series of
heuristics to every word to come up with a best guess.

It was important for me to try and make my models as fat as possible while
keeping my controllers clean. I think this was accomplished given the simplicity
of my `main` function versus how much of the functionality is tied to the ticket
struct. The ticket struct owns all of the various heuristics necessary to boil
down to the best guess. This architecture still has a way to go but I fear if
I work on it any longer I will have built a SeatGeek competitor ;)

It's now clear to me that I should have started by examining my data and writing
test cases such that I could have jumped write to my final approach. It would
have saved me a lot of time spinning my wheels on a reusable architecture.

Regardless, I have taken the liberty of writing some test cases around the
majority of logic for ticket normalization. Further testing would require me to
use more dependency injection and interfaces to be able to mock out CSV files
or at least the data they produce.

A large portion of my time went into exploring the build script so I could
recreate the environment in Go. Assuming you have go installed on your machine,
you should be able to run the autograder as normal just by specifying `go` as
the language.
