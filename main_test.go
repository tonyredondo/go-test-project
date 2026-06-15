package main

import (
	"testing"
	"time"
)

/*
func TestMain(m *testing.M) {
		// Gets the CodeOwners from the file
		codeOwners := GetCodeOwners()
		SetGlobalEventFinishHook(func(testsAny []any) {
			for _, testAny := range testsAny {
				test := testAny.(Test)

				// This is where you can add your custom logic for handling test events.
				fmt.Println("********** Test finished:", test.Name())

				// Get tags from the test
				cmd, _ := test.GetTag("test.command")
				fmt.Println(cmd)

				// Get the codeowners from the test
				if file, ok := test.GetTag("test.source.file"); ok {
					if entry, ok := codeOwners.Match(file.(string)); ok {
						fmt.Println(entry.Owners)
					}
				}

				// Set a tag
				test.SetTag("test.tag", "test_value")
				tv, _ := test.GetTag("test.tag")
				fmt.Println(tv)
			}
		})
	os.Exit(m.Run())
}
*/

func TestDataDog_Parallelism(t *testing.T) {
	t.Parallel()
	// mTracer := mocktracer.Start()
	// defer mTracer.Stop()
	type testcase struct {
		name string
		text string
	}
	tests := []testcase{
		{name: "test1", text: "something to print"},
		{name: "test2", text: "something to print"},
		{name: "test3", text: "something to print"},
		{name: "test4", text: "something to print"},
		{name: "test5", text: "something to print"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(_t *testing.T) {
			// _t.Parallel()
			// Simulate some work
			for i := 0; i < 1; i++ {
				t.Run(tc.text, func(t *testing.T) {
					// t.Parallel()
					// Simulate some work
					_ = i

					// span := mTracer.StartSpan("normal_span:" + tc.text)
					// defer span.Finish()
				})
			}
			// Print the text
			t.Log(tc.text)
		})
	}
	<-time.After(1 * time.Second)

	// fmt.Println("Finished spans:", len(mTracer.FinishedSpans()))
}

func TestDataDog_Parallelism2(t *testing.T) {
	t.Parallel()
	<-time.After(1 * time.Second)
}

func TestDataDog_Parallelism3(t *testing.T) {
	t.Parallel()
	<-time.After(1 * time.Second)
}

func TestDataDog_Parallelism4(t *testing.T) {
	t.Parallel()
	<-time.After(1 * time.Second)
}
