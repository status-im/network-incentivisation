module.exports = {
  env: {
    browser: true,
    es6: true,
  },
  extends: 'airbnb-base',
  globals: {
    Atomics: 'readonly',
    SharedArrayBuffer: 'readonly',
    artifacts: 'readonly',
    contract: 'readonly',
    describe: 'readonly',
    beforeEach: 'readonly',
    before: 'readonly',
    afterEach: 'readonly',
    after: 'readonly',
    it: 'readonly',
    assert: 'readonly',
  },
  parserOptions: {
    ecmaVersion: 2018,
    sourceType: 'module',
  },
};
