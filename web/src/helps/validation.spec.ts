import { expect, test } from 'vitest'
import { Size, Valid } from "./validations"

test('validations Size', () => {
  expect(Size('validations', 11)).toEqual(true)
})

test('validations Regex', () => {
  expect(Valid('Hello', /^[A-z]*$/g)).toEqual(true)
  expect(Valid(123, /^[A-z]*$/g)).toEqual(false)
})
