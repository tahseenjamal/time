package main

import (
	"fmt"
	"time"
)

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end)

}

func main() {

	start, _ := time.Parse(time.RFC822, "01 Jan 15 10:00 UTC")

	end, _ := time.Parse(time.RFC822, "01 Jan 16 10:00 UTC")

	in, _ := time.Parse(time.RFC822, "01 Jan 15 20:00 UTC")

	out, _ := time.Parse(time.RFC822, "01 Jan 17 10:00 UTC")

	if inTimeSpan(start, end, in) {

		fmt.Println(in, "is between", start, "and", end, ".")

	}

	if !inTimeSpan(start, end, out) {

		fmt.Println(out, "is not between", start, "and", end, ".")

	}

	india, _ := time.LoadLocation("Asia/Calcutta")

	//accepted time format in text by Cassandra
	//http://docs.datastax.com/en/archived/cql/3.0/cql/cql_reference/timestamp_type_r.html
	fmt.Println(time.Now().In(india).Format("2006-01-02 15:04:05-07:00"))

	fmt.Println(time.Now().Format("2006-01-02 15:04:05-07:00"))

	a, _ := time.Parse("2006-01-02 15:04:05-07:00", "2017-04-02 09:59:54+05:30")

	if a.After(time.Now()) {

		fmt.Println(a, " is after ", time.Now().Format("2006-01-02 15:04:05-07:00"))

	} else {

		fmt.Println(a, " is before ", time.Now().Format("2006-01-02 15:04:05-07:00"))

	}

	fmt.Println(time.Now().Sub(a).Minutes())

}
