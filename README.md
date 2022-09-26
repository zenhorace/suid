# SUID - Short UUIDs

have you ever wanted a UUID but had a constraint on how long the resulting string could be?\
UUIDs (version 4) use 36 characters for the string representation (32 if you remove the hyphens). For majority of situations, that's perfectly fine. This library is to address the cases where you can't afford that string size but still want a solution generates IDs just as unique.*

## Features & Usage:
// TODO


*I make no promises here. I've followed the methods of the [UUID](https://github.com/google/uuid) package, but I'm far from a security expert. If you see some fatal flaw, please let me know 🙂
