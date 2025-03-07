package regolib

const (
	targetLibSrc = `
package hooks["{{.Target}}"]

# Finds all violations for a given target
violation[response] {
	data.hooks["{{.Target}}"].library.matching_constraints[constraint]
	review := get_default(input, "review", {})
	inp := {
		"review": review,
		"constraint": constraint
	}
	inventory[inv]
	data.templates["{{.Target}}"][constraint.kind].violation[r] with input as inp with data.inventory as inv
	response = {
		"msg": r.msg,
		"metadata": {"details": get_default(r, "details", {})},
		"constraint": constraint,
		"review": review,
	}
}


# Finds all violations in the cached state of a given target
audit[response] {
	data.hooks["{{.Target}}"].library.matching_reviews_and_constraints[[review, constraint]]
	inp := {
		"review": review,
		"constraint": constraint,
	}
	inventory[inv]
	data.templates["{{.Target}}"][constraint.kind].violation[r] with input as inp with data.inventory as inv
	response = {
		"msg": r.msg,
		"metadata": {"details": get_default(r, "details", {})},
		"constraint": constraint,
		"review": review,
	}
}

# get_default(data, "external", {}) seems to cause this error:
# "rego_type_error: undefined function data.hooks.<target>.get_default"
inventory[inv] {
	inv = data.external["{{.Target}}"]
}

inventory[{}] {
	not data.external["{{.Target}}"]
}

# get_default returns the value of an object's field or the provided default value.
# It avoids creating an undefined state when trying to access an object attribute that does
# not exist
get_default(object, field, _default) = object[field]

get_default(object, field, _default) = _default {
  not has_field(object, field)
}

has_field(object, field) {
  _ = object[field]
}
`
)
