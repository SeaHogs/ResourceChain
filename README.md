#ResourceChian
# Consensus-Breaking Change

## What Does It Mean By Breaking Consensus?
**Consensus** in a blockchain context refers to all nodes on the network
agreeing on the same state. A consensus-breaking change is a modification
that alters the state machine in such a way that older nodes and newer
nodes will produce different states when processing the same blocks.

## Why Does This Change Break Consensus?
We modified the `Resource` data structure to add a new mandatory field
(`creation_time`) and reorder existing fields. Some of the nodes might run the old
version of the code, they do not recognize this new field and will handle
transactions differently. As a result, the network will have two versions
of the chain state and produce different block hashes, therefore breaking
consensus.

