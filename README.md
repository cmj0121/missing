# Missing #
The simple [pre-commit][0] hook to find the missing but necessary file.

## Example ##
```yaml
- repo: https://github.com/cmj0121/missing
  rev: v0.2.3
  hooks:
    - id: missing-init-py
```

## arguments ##

| argument         | type           | description             |
|------------------|----------------|-------------------------|
| `-e` `--exclude` | List of string | exclude the folder path |

[0]: https://pre-commit.com/
