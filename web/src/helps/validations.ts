export function Size(value: String, expected: Number): Boolean {
  return value.length == expected
}

export function Valid(value: string|number, regex: RegExp): Boolean {
  return regex.test(String(value))
}
