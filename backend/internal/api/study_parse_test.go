package api

import (
	"encoding/json"
	"testing"
)

func TestStudyParseCatalogTimezoneAndRequiredEnd(t *testing.T) {
	const courseID = "11111111-1111-4111-8111-111111111111"
	const lectureID = "22222222-2222-4222-8222-222222222222"
	const otherID = "33333333-3333-4333-8333-333333333333"
	courses := []map[string]any{{"id": courseID, "lectures": []map[string]any{{"id": lectureID}}}, {"id": otherID, "lectures": []map[string]any{}}}
	for _, test := range []struct {
		name, course, lecture, end, zone string
		valid                            bool
	}{
		{"known", courseID, lectureID, "2026-10-23T10:00:00+08:00", "Asia/Shanghai", true},
		{"unknown-course", lectureID, "", "2026-10-23T10:00:00+08:00", "Asia/Shanghai", false},
		{"wrong-lecture", otherID, lectureID, "2026-10-23T10:00:00+08:00", "Asia/Shanghai", false},
		{"missing-end", courseID, "", "", "Asia/Shanghai", false},
		{"wrong-offset", courseID, "", "2026-10-23T10:00:00Z", "Asia/Shanghai", false},
		{"winter", courseID, "", "2026-01-20T10:00:00-05:00", "America/New_York", true},
		{"summer", courseID, "", "2026-07-20T10:00:00-04:00", "America/New_York", true},
		{"summer-wrong", courseID, "", "2026-07-20T10:00:00-05:00", "America/New_York", false},
		{"spring-gap", courseID, "", "2026-03-08T02:30:00-05:00", "America/New_York", false},
		{"fall-first", courseID, "", "2026-11-01T01:30:00-04:00", "America/New_York", true},
		{"fall-second", courseID, "", "2026-11-01T01:30:00-05:00", "America/New_York", true},
	} {
		t.Run(test.name, func(t *testing.T) {
			row := map[string]any{"id": otherID, "courseId": test.course, "title": "线性代数", "chapter": "第一章", "endsAt": test.end}
			if test.lecture != "" {
				row["lectureId"] = test.lecture
			}
			body, _ := json.Marshal([]any{row})
			err := validStudyRowsFor(body, courses, test.zone)
			if (err == nil) != test.valid {
				t.Fatal(err)
			}
		})
	}
	row := map[string]any{"id": otherID, "courseId": courseID, "title": "手工条目", "chapter": "第一章", "endsAt": ""}
	if err := validateStudyEntry(row); err != nil {
		t.Fatal("manual empty end compatibility", err)
	}
}

func TestStudyParseTimezoneRejectedBeforeModelCall(t *testing.T) {
	for _, zone := range []string{"Local", "Mars/Olympus", "+08:00", ""} {
		raw, _ := json.Marshal(map[string]string{"text": "课表", "timezone": zone})
		if _, err := parseStudyInput(raw); err == nil {
			t.Fatal(zone)
		}
	}
	for _, zone := range []string{"UTC", "Asia/Shanghai", "America/New_York"} {
		raw, _ := json.Marshal(map[string]string{"text": "课表", "timezone": zone})
		if _, err := parseStudyInput(raw); err != nil {
			t.Fatal(zone, err)
		}
	}
}
