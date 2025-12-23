package twofer

// ShareWith returns a string following "One for you, one for me." template.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
