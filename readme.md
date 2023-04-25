# Consistent hashing
Traditional hashing uses a modulo function to map keys to nodes, this works well for a fixed number of nodes, but adding or removing nodes requires data migration. Consistent hashing maps keys to a range of hash values and uses a hash ring to map keys to nodes, which allows nodes to be added or removed with minimal data migration. This makes consistent hashing more efficient and scalable in distributed systems.


## Reference
- [System Design — Consistent Hashing](https://medium.com/must-know-computer-science/system-design-consistent-hashing-f66fa9b75f3f)
- [Consistent Hashing | Algorithms You Should Know](https://www.youtube.com/watch?v=UF9Iqmg94tk)
