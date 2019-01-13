# Algorand-Sortition

This is a toy implementation of Algorand's Cryptographic Sortition algorithm. Cryptographic Sortition secretly and randomly selects a subset of all participants based on Verifiable Random Functions (VRF). The selected subset becomes the committee of this round, which runs an optimized BFT consensus protocol called BA* to decide a single block.