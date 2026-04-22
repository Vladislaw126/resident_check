# What are the roles of the included Go services?

## `apps/approach_a/verifier_terminal/`

This directory contains the shop-facing verifier application used by the coffeeshop clerk. Its responsibility is to accept a user-entered identifier or query input, send that identity query to the central lookup API, and display the returned result to the clerk. It should stay intentionally thin: it is not responsible for deciding residency, interpreting resident records deeply, or storing authoritative data. Its role is to act as the entry point at the point of sale and to make visible how much information Approach A exposes back to the shop.

## `apps/approach_a/central_lookup_api/`

This directory contains the central backend service that implements the core logic of Approach A. Its responsibility is to receive identity queries from the verifier terminal, validate and normalize them, forward them to the municipality lookup mock, receive the resident record response, and return rich identity/residency information back to the verifier terminal. This is the architectural center of the baseline design, and also the place where the main privacy weaknesses of Approach A become visible, because it concentrates lookup power, sees the full query flow, and can return more data than is strictly necessary.

## `apps/approach_a/municipality_lookup_mock/`

This directory contains the mock authority-side service that simulates a municipality or authoritative resident-record lookup system. Its responsibility is to receive a lookup request from the central lookup API, search a mock resident dataset, and return a resident record or failure result. It should behave like an external dependency rather than like part of the same app, because that boundary is important to preserve in your experiments. Its purpose is not to be realistic in every administrative detail, but to give the baseline architecture a concrete external source of resident data.