package main

import "testing"

func Test_parseReport(t *testing.T) {
	lines := []string{
		"Valve AA has flow rate=0; tunnels lead to valves BB, CC",
		"Valve BB has flow rate=12; tunnels lead to valves CC",
		"Valve CC has flow rate=32; tunnels lead to valves BB",
	}

	got := parseInput(lines)

	assertEqual(t, got.name, "AA")
	assertEqual(t, got.flowRate, 0)
	assertEqual(t, len(got.leadTo), 2)

	for _, gotLeadTo := range got.leadTo {
		if gotLeadTo.name == "BB" {
			assertEqual(t, gotLeadTo.name, "BB")
			assertEqual(t, gotLeadTo.flowRate, 12)
			assertEqual(t, len(gotLeadTo.leadTo), 1)

			assertEqual(t, gotLeadTo.leadTo[0].name, "CC")
		} else {
			assertEqual(t, gotLeadTo.name, "CC")
			assertEqual(t, gotLeadTo.flowRate, 32)
			assertEqual(t, len(gotLeadTo.leadTo), 1)

			assertEqual(t, gotLeadTo.leadTo[0].name, "BB")

		}
	}
}

func assertEqual[T comparable](t *testing.T, got, want T) {
	if got != want {
		t.Helper()
		t.Errorf("got %#v, want %#v", got, want)
	}
}
