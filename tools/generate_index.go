// From TEACH_ME_DEV/hub/generate_index.go.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type HubCourseYAML struct {
	Name    string
	Summary string
	Author  string
	Shared  string
	From    string
	License string
	Level   string
	Hours   float64
	Tags    []string
	Lessons int
}

func parseCourseYAML(b []byte) HubCourseYAML {
	var cy HubCourseYAML
	scanner := bufio.NewScanner(bytes.NewReader(b))
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		val := strings.TrimSpace(parts[1])
		val = strings.Trim(val, `"'`)

		switch key {
		case "name":
			cy.Name = val
		case "summary":
			cy.Summary = val
		case "author":
			cy.Author = val
		case "shared":
			cy.Shared = val
		case "from":
			cy.From = val
		case "license":
			cy.License = val
		case "level":
			cy.Level = val
		case "hours":
			cy.Hours, _ = strconv.ParseFloat(val, 64)
		case "lessons":
			cy.Lessons, _ = strconv.Atoi(val)
		case "tags":
			val = strings.Trim(val, "[]")
			rawTags := strings.Split(val, ",")
			for _, t := range rawTags {
				t = strings.TrimSpace(strings.Trim(t, `"'`))
				if t != "" {
					cy.Tags = append(cy.Tags, t)
				}
			}
		}
	}
	return cy
}

type CourseJSON struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type IndexEntry struct {
	Folder  string   `json:"folder"`
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	Summary string   `json:"summary"`
	Author  string   `json:"author"`
	Shared  string   `json:"shared"`
	Level   string   `json:"level"`
	Hours   float64  `json:"hours"`
	Tags    []string `json:"tags"`
	Lessons int      `json:"lessons"`
	Commit  string   `json:"commit,omitempty"`
}

type HubIndex struct {
	Format    int          `json:"format"`
	Generated string       `json:"generated"`
	Courses   []IndexEntry `json:"courses"`
}

func main() {
	coursesDir := "courses"
	entries := []IndexEntry{}

	dirs, err := os.ReadDir(coursesDir)
	if err == nil {
		for _, d := range dirs {
			if !d.IsDir() || strings.HasPrefix(d.Name(), ".") {
				continue
			}
			folder := d.Name()
			cPath := filepath.Join(coursesDir, folder)

			entry := IndexEntry{
				Folder:  folder,
				ID:      folder,
				Name:    folder,
				Level:   "beginner",
				Lessons: 1,
			}

			// Read course.yaml
			if yb, err := os.ReadFile(filepath.Join(cPath, "course.yaml")); err == nil {
				cy := parseCourseYAML(yb)
				if cy.Name != "" {
					entry.Name = cy.Name
				}
				entry.Summary = cy.Summary
				entry.Author = cy.Author
				entry.Shared = cy.Shared
				entry.Level = cy.Level
				entry.Hours = cy.Hours
				entry.Tags = cy.Tags
				entry.Lessons = cy.Lessons
			}

			// Read course.json
			if jb, err := os.ReadFile(filepath.Join(cPath, "course.json")); err == nil {
				var cj CourseJSON
				if err := json.Unmarshal(jb, &cj); err == nil {
					if cj.ID != "" {
						entry.ID = cj.ID
					}
					if entry.Name == folder && cj.Title != "" {
						entry.Name = cj.Title
					}
					if entry.Summary == "" {
						entry.Summary = cj.Description
					}
				}
			}

			entries = append(entries, entry)
		}
	}

	index := HubIndex{
		Format:    1,
		Generated: time.Now().UTC().Format(time.RFC3339),
		Courses:   entries,
	}

	out, err := json.MarshalIndent(index, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error encoding index: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile("index.json", out, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "error writing index.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Wrote index.json with %d courses\n", len(entries))
}
