# policies

These are the [Open Policy Agent](https://www.openpolicyagent.org/) policies used by the services.

Each policy should have a test suite. Run the tests with `opa test -v .` and you should see the following:

```bash
data.mod_5.test_allow_5: PASS (708.779µs)
data.mod_5.test_allow_15: PASS (506.891µs)
data.mod_5.test_disallow_with_3: PASS (539.41µs)
--------------------------------------------------------------------------------
PASS: 3/3
```