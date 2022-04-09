// jest-dom adds custom jest matchers for asserting on DOM nodes.
// allows you to do things like:
// expect(element).toHaveTextContent(/react/i)
// learn more: https://github.com/testing-library/jest-dom
import '@testing-library/jest-dom';

if (!process.env.DEBUG) {
  console.debug = () => undefined;
}

// During tests, `scrollIntoView` is not a function that's implemented (or stubbed) by jsdom.
// See: https://github.com/jsdom/jsdom/issues/1695
Element.prototype.scrollIntoView = () => undefined;
