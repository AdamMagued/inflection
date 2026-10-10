package inflection

import (
	"strings"
	"testing"
)

var inflections = map[string]string{
	"star":        "stars",
	"STAR":        "STARS",
	"Star":        "Stars",
	"bus":         "buses",
	"fish":        "fish",
	"mouse":       "mice",
	"query":       "queries",
	"ability":     "abilities",
	"agency":      "agencies",
	"movie":       "movies",
	"archive":     "archives",
	"index":       "indices",
	"wife":        "wives",
	"safe":        "saves",
	"half":        "halves",
	"move":        "moves",
	"salesperson": "salespeople",
	"person":      "people",
	"spokesman":   "spokesmen",
	"man":         "men",
	"woman":       "women",
	"basis":       "bases",
	"diagnosis":   "diagnoses",
	"diagnosis_a": "diagnosis_as",
	"datum":       "data",
	"medium":      "media",
	"stadium":     "stadia",
	"analysis":    "analyses",
	"node_child":  "node_children",
	"child":       "children",
	"experience":  "experiences",
	"day":         "days",
	"comment":     "comments",
	"foobar":      "foobars",
	"newsletter":  "newsletters",
	"old_news":    "old_news",
	"news":        "news",
	"series":      "series",
	"species":     "species",
	"quiz":        "quizzes",
	"perspective": "perspectives",
	"ox":          "oxen",
	"photo":       "photos",
	"buffalo":     "buffaloes",
	"tomato":      "tomatoes",
	"dwarf":       "dwarves",
	"elf":         "elves",
	"information": "information",
	"equipment":   "equipment",
	"criterion":   "criteria",
	"foot":        "feet",
	"goose":       "geese",
	"moose":       "moose",
	"tooth":       "teeth",
	"milk":        "milk",
	"salt":        "salt",
	"time":        "time",
	"water":       "water",
	"paper":       "paper",
	"music":       "music",
	"help":        "help",
	"luck":        "luck",
	"oil":         "oil",
	"progress":    "progress",
	"rain":        "rain",
	"research":    "research",
	"shopping":    "shopping",
	"software":    "software",
	"traffic":     "traffic",
	"zombie":      "zombies",
	"campus":      "campuses",
	"harddrive":   "harddrives",
	"drive":       "drives",
	"leaf":        "leaves",
	"wolf":        "wolves",
}

// storage is used to restore the state of the global variables
// on each test execution, to ensure no global state pollution
type storage struct {
	singulars    RegularSlice
	plurals      RegularSlice
	irregulars   IrregularSlice
	uncountables []string
}

var backup = storage{}

func init() {
	AddIrregular("criterion", "criteria")
	backup.singulars = GetSingular()
	backup.plurals = GetPlural()
	backup.irregulars = GetIrregular()
	backup.uncountables = GetUncountable()
}

func restore() {
	SetSingular(backup.singulars)
	SetPlural(backup.plurals)
	SetIrregular(backup.irregulars)
	SetUncountable(backup.uncountables)
}

func TestPlural(t *testing.T) {
	for key, value := range inflections {
		if v := Plural(strings.ToUpper(key)); v != strings.ToUpper(value) {
			t.Errorf("%v's plural should be %v, but got %v", strings.ToUpper(key), strings.ToUpper(value), v)
		}

		if v := Plural(strings.Title(key)); v != strings.Title(value) {
			t.Errorf("%v's plural should be %v, but got %v", strings.Title(key), strings.Title(value), v)
		}

		if v := Plural(key); v != value {
			t.Errorf("%v's plural should be %v, but got %v", key, value, v)
		}
	}
}

func TestSingular(t *testing.T) {
	for key, value := range inflections {
		if v := Singular(strings.ToUpper(value)); v != strings.ToUpper(key) {
			t.Errorf("%v's singular should be %v, but got %v", strings.ToUpper(value), strings.ToUpper(key), v)
		}

		if v := Singular(strings.Title(value)); v != strings.Title(key) {
			t.Errorf("%v's singular should be %v, but got %v", strings.Title(value), strings.Title(key), v)
		}

		if v := Singular(value); v != key {
			t.Errorf("%v's singular should be %v, but got %v", value, key, v)
		}
	}
}

func TestAddPlural(t *testing.T) {
	defer restore()
	ln := len(pluralInflections)
	AddPlural("", "")
	if ln+1 != len(pluralInflections) {
		t.Errorf("Expected len %d, got %d", ln+1, len(pluralInflections))
	}
}

func TestAddSingular(t *testing.T) {
	defer restore()
	ln := len(singularInflections)
	AddSingular("", "")
	if ln+1 != len(singularInflections) {
		t.Errorf("Expected len %d, got %d", ln+1, len(singularInflections))
	}
}

func TestAddIrregular(t *testing.T) {
	defer restore()
	ln := len(irregularInflections)
	AddIrregular("", "")
	if ln+1 != len(irregularInflections) {
		t.Errorf("Expected len %d, got %d", ln+1, len(irregularInflections))
	}
}

func TestAddUncountable(t *testing.T) {
	defer restore()
	ln := len(uncountableInflections)
	AddUncountable("", "")
	if ln+2 != len(uncountableInflections) {
		t.Errorf("Expected len %d, got %d", ln+2, len(uncountableInflections))
	}
}

func TestGetPlural(t *testing.T) {
	plurals := GetPlural()
	if len(plurals) != len(pluralInflections) {
		t.Errorf("Expected len %d, got %d", len(plurals), len(pluralInflections))
	}
}

func TestGetSingular(t *testing.T) {
	singular := GetSingular()
	if len(singular) != len(singularInflections) {
		t.Errorf("Expected len %d, got %d", len(singular), len(singularInflections))
	}
}

func TestGetIrregular(t *testing.T) {
	irregular := GetIrregular()
	if len(irregular) != len(irregularInflections) {
		t.Errorf("Expected len %d, got %d", len(irregular), len(irregularInflections))
	}
}

func TestGetUncountable(t *testing.T) {
	uncountables := GetUncountable()
	if len(uncountables) != len(uncountableInflections) {
		t.Errorf("Expected len %d, got %d", len(uncountables), len(uncountableInflections))
	}
}

func TestSetPlural(t *testing.T) {
	defer restore()
	SetPlural(RegularSlice{{}, {}})
	if len(pluralInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(pluralInflections))
	}
}

func TestSetSingular(t *testing.T) {
	defer restore()
	SetSingular(RegularSlice{{}, {}})
	if len(singularInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(singularInflections))
	}
}

func TestSetIrregular(t *testing.T) {
	defer restore()
	SetIrregular(IrregularSlice{{}, {}})
	if len(irregularInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(irregularInflections))
	}
}

func TestSetUncountable(t *testing.T) {
	defer restore()
	SetUncountable([]string{"", ""})
	if len(uncountableInflections) != 2 {
		t.Errorf("Expected len 2, got %d", len(uncountableInflections))
	}
}

func TestVesSingularPlural(t *testing.T) {
	cases := []struct {
		singular string
		plural   string
	}{
		{"leaf", "leaves"},
		{"wolf", "wolves"},
		{"half", "halves"},
	}

	for _, c := range cases {
		// Test lowercase
		if got := Plural(c.singular); got != c.plural {
			t.Errorf("Plural(%q) = %q, want %q", c.singular, got, c.plural)
		}
		if got := Singular(c.plural); got != c.singular {
			t.Errorf("Singular(%q) = %q, want %q", c.plural, got, c.singular)
		}

		// Test Title case
		titleSingular := strings.Title(c.singular)
		titlePlural := strings.Title(c.plural)
		if got := Plural(titleSingular); got != titlePlural {
			t.Errorf("Plural(%q) = %q, want %q", titleSingular, got, titlePlural)
		}
		if got := Singular(titlePlural); got != titleSingular {
			t.Errorf("Singular(%q) = %q, want %q", titlePlural, got, titleSingular)
		}

		// Test UPPERCASE
		upperSingular := strings.ToUpper(c.singular)
		upperPlural := strings.ToUpper(c.plural)
		if got := Plural(upperSingular); got != upperPlural {
			t.Errorf("Plural(%q) = %q, want %q", upperSingular, got, upperPlural)
		}
		if got := Singular(upperPlural); got != upperSingular {
			t.Errorf("Singular(%q) = %q, want %q", upperPlural, got, upperSingular)
		}
	}

	// Verify regular regex rules operate correctly in isolation without irregular overrides
	oldIrregulars := GetIrregular()
	defer SetIrregular(oldIrregulars)
	SetIrregular(IrregularSlice{})

	for _, c := range cases {
		if got := Plural(c.singular); got != c.plural {
			t.Errorf("Plural(%q) without irregulars = %q, want %q", c.singular, got, c.plural)
		}
		if got := Singular(c.plural); got != c.singular {
			t.Errorf("Singular(%q) without irregulars = %q, want %q", c.plural, got, c.singular)
		}

		titleSingular := strings.Title(c.singular)
		titlePlural := strings.Title(c.plural)
		if got := Plural(titleSingular); got != titlePlural {
			t.Errorf("Plural(%q) without irregulars = %q, want %q", titleSingular, got, titlePlural)
		}
		if got := Singular(titlePlural); got != titleSingular {
			t.Errorf("Singular(%q) without irregulars = %q, want %q", titlePlural, got, titleSingular)
		}

		upperSingular := strings.ToUpper(c.singular)
		upperPlural := strings.ToUpper(c.plural)
		if got := Plural(upperSingular); got != upperPlural {
			t.Errorf("Plural(%q) without irregulars = %q, want %q", upperSingular, got, upperPlural)
		}
		if got := Singular(upperPlural); got != upperSingular {
			t.Errorf("Singular(%q) without irregulars = %q, want %q", upperPlural, got, upperSingular)
		}
	}
}

