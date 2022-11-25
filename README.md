# Disable meta key
## What is this?
This a go script to disable the meta key on a key press.
## How does it work?
It works via the kde config file located in your `.config/` directory.
It appends:
```
[ModifierOnlyShortcuts]
Meta=
```
So the meta key is set to nothing and if you press it again it will remove it via reading the file.
## Planned Features
- remove the lines if already present to deactivate it. □
- Get the current user so to go to the correct users home. ☑
- Write the lines to the file. □
- Fix bugs □